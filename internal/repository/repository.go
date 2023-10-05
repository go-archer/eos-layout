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
)

type Repository struct {
	cfg *config.Config
	db  *gorm.DB
	rds *redis.Client
	log *log.Logger
}

func NewRepository(cfg *config.Config, log *log.Logger) *Repository {
	repo := &Repository{cfg: cfg, log: log}
	repo.initDB()
	repo.initRedis()
	return repo
}

func (r Repository) initDB() {
	if r.cfg.MySQL == nil {
		return
	}
	db, err := gorm.Open(mysql.Open(r.cfg.MySQL.DSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("mysql error: %s", err))
	}
	r.db = db
}

func (r Repository) initRedis() {
	if r.cfg.Redis == nil {
		return
	}
	rds := redis.NewClient(&redis.Options{
		Addr:         r.cfg.Redis.Addr,
		Password:     r.cfg.Redis.Password,
		DB:           r.cfg.Redis.DB,
		DialTimeout:  time.Duration(r.cfg.Redis.DialTimeout),
		ReadTimeout:  time.Duration(r.cfg.Redis.ReadTimeout),
		WriteTimeout: time.Duration(r.cfg.Redis.WriteTimeout),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rds.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err))
	}
	r.rds = rds
}

func (r Repository) DB(ctx ...context.Context) *gorm.DB {
	if len(ctx) > 0 {
		return r.db.WithContext(ctx[0])
	}
	return r.db
}

func (r Repository) AutoMigrate(table any) error {
	if r.cfg.MySQL.AutoMigrate {
		err := r.db.AutoMigrate(table)
		if err != nil {
			r.log.Error(err.Error())
			return err
		}
	}
	return nil
}

func (r Repository) Lock(ctx context.Context, key string, acquire, timeout time.Duration) (string, error) {
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

func (r Repository) UnLock(ctx context.Context, key, code string) bool {
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

func (r Repository) Cache() *redis.Client {
	return r.rds
}

func (r Repository) CacheGet(ctx context.Context, key string) ([]byte, error) {
	if r.rds == nil {
		return nil, errors.New("redis connection error")
	}
	return r.rds.Get(ctx, key).Bytes()
}

func (r Repository) CacheSet(ctx context.Context, key string, data interface{}, exp time.Duration) error {
	if r.rds == nil {
		return errors.New("redis connection error")
	}
	return r.rds.Set(ctx, key, data, exp).Err()
}
