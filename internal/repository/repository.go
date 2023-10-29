package repository

import (
	"context"
	"eos-layout/internal/config"
	"eos-layout/pkg/log"
	"eos-layout/pkg/uuid"
	"errors"
	"fmt"
	"time"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Set = wire.NewSet(
	NewRepository,
	NewAreaRepository,
	NewRepositoryContext,
)

type Repository interface {
	DB(ctx ...context.Context) *gorm.DB
	RDS() *redis.Client
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, data any, exp time.Duration) error
	Del(ctx context.Context, key string) error
	Lock(ctx context.Context, key string, acquire, timeout time.Duration) (string, error)
	UnLock(ctx context.Context, key, code string) bool
	Context() *Context
	Log() *log.Logger
}

type repository struct {
	cfg *config.Config
	db  *gorm.DB
	rds *redis.Client
	log *log.Logger
}

// 仓储上下文对象
var repositoryContext *Context

// Context 仓储上下文，用于共享仓储操作
type Context struct {
	Area AreaRepository
}

func NewRepositoryContext(area AreaRepository) *Context {
	repositoryContext = &Context{
		Area: area,
	}
	return repositoryContext
}

func NewRepository(cfg *config.Config, log *log.Logger) Repository {
	repo := &repository{cfg: cfg, log: log}
	repo.initDB()
	repo.initRedis()
	return repo
}

func (r *repository) initDB() {
	if r.cfg.MySQL == nil {
		return
	}
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		r.cfg.MySQL.Username, r.cfg.MySQL.Password, r.cfg.MySQL.Host, r.cfg.MySQL.Port, r.cfg.MySQL.DBName)
	db, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("mysql error: %s", err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("mysql error: %s", err))
	}
	sqlDB.SetMaxIdleConns(r.cfg.MySQL.MaxIDLE)
	sqlDB.SetMaxOpenConns(r.cfg.MySQL.MaxOpen)
	r.db = db
}

func (r *repository) initRedis() {
	if r.cfg.Redis == nil {
		return
	}
	rds := redis.NewClient(&redis.Options{
		Addr:         r.cfg.Redis.Addr,
		Username:     r.cfg.Redis.Username,
		Password:     r.cfg.Redis.Password,
		DB:           r.cfg.Redis.DB,
		DialTimeout:  time.Millisecond * time.Duration(r.cfg.Redis.DialTimeout),
		ReadTimeout:  time.Millisecond * time.Duration(r.cfg.Redis.ReadTimeout),
		WriteTimeout: time.Millisecond * time.Duration(r.cfg.Redis.WriteTimeout),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rds.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err))
	}
	r.rds = rds
}

func (r *repository) Context() *Context {
	return repositoryContext
}

func (r *repository) Log() *log.Logger {
	return r.log
}

func (r *repository) DB(ctx ...context.Context) *gorm.DB {
	if len(ctx) > 0 {
		return r.db.WithContext(ctx[0])
	}
	return r.db
}

func (r *repository) AutoMigrate(table any) error {
	if r.cfg.MySQL.AutoMigrate {
		err := r.db.AutoMigrate(table)
		if err != nil {
			r.log.Error(err.Error())
			return err
		}
	}
	return nil
}

func (r *repository) Lock(ctx context.Context, key string, acquire, timeout time.Duration) (string, error) {
	code := uuid.UUID()
	endTime := time.Now().Add(acquire).UnixNano()
	for time.Now().UnixNano() <= endTime {
		if success, err := r.rds.SetNX(ctx, key, code, timeout).Result(); err != nil {
			return "", err
		} else if success {
			return code, nil
		} else if r.rds.TTL(ctx, key).Val() == -1 {
			r.rds.Expire(ctx, key, timeout)
		}
		time.Sleep(time.Millisecond)
	}
	return "", errors.New("lock timeout")
}

func (r *repository) UnLock(ctx context.Context, key, code string) bool {
	tx := func(tx *redis.Tx) error {
		if v, err := tx.Get(ctx, key).Result(); err != nil && err != redis.Nil {
			return err
		} else if v == code {
			_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
				pipe.Del(ctx, key)
				return nil
			})
			return err
		}
		return nil
	}
	for {
		if err := r.rds.Watch(ctx, tx, key); err == nil {
			return true
		} else if err == redis.TxFailedErr {
			r.log.Warn("watch key is modified,retry to release lock. ", zap.Error(err), zap.String("key", key), zap.String("code", code))
		} else {
			return false
		}
	}
}

func (r *repository) RDS() *redis.Client {
	return r.rds
}

func (r *repository) Get(ctx context.Context, key string) ([]byte, error) {
	if r.rds == nil {
		return nil, errors.New("redis connection error")
	}
	return r.rds.Get(ctx, key).Bytes()
}

func (r *repository) Set(ctx context.Context, key string, data interface{}, exp time.Duration) error {
	if r.rds == nil {
		return errors.New("redis connection error")
	}
	return r.rds.Set(ctx, key, data, exp).Err()
}

func (r *repository) Del(ctx context.Context, key string) error {
	if r.rds == nil {
		return errors.New("redis connection error")
	}
	return r.rds.Del(ctx, key).Err()
}
