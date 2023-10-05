package service

import (
	"eos-layout/internal/config"
	"eos-layout/pkg/log"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewService,
	NewAreaService,
)

// Service 基础服务
type Service struct {
	cfg *config.Config
	log *log.Logger
}

func NewService(cfg *config.Config, log *log.Logger) *Service {
	return &Service{cfg: cfg, log: log}
}
