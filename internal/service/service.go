package service

import (
	"eos-layout/internal/config"
	"eos-layout/pkg/log"
)

type Service interface {
	Config() *config.Config
	Log() *log.Logger
}

// Service 基础服务
type service struct {
	cfg *config.Config
	log *log.Logger
}

func NewService(cfg *config.Config, log *log.Logger) Service {
	return &service{cfg: cfg, log: log}
}

func (s service) Config() *config.Config {
	return s.cfg
}

func (s service) Log() *log.Logger {
	return s.log
}
