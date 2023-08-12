package repository

import (
	"context"
	"eos-layout/internal/config"
	"eos-layout/pkg/log"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func (r *Repository) initDB() {
	if r.cfg.MySQL == nil {
		return
	}
	db, err := gorm.Open(mysql.Open(r.cfg.MySQL.DSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("mysql error: %s", err))
	}
	r.db = db
}

func (r *Repository) initRedis() {
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
