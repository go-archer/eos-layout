package service

import (
	"context"
	"eos-layout/internal/config"
	"eos-layout/internal/status"
	"eos-layout/pkg/log"

	"github.com/gin-gonic/gin"
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

func (s *Service) TID(ctx context.Context) (int64, error) {
	c := ctx.(*gin.Context)
	id, ok := c.Get("TID")
	if !ok {
		return 0, status.ErrorAuthorize
	}
	tid, ok := id.(int64)
	if !ok || tid == 0 {
		return 0, status.ErrorAuthorize
	}
	return tid, nil
}
