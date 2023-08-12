//go:build wireinject
// +build wireinject

package main

import (
	"eos-layout/internal/config"
	"eos-layout/internal/handler"
	"eos-layout/internal/repository"
	"eos-layout/internal/server"
	"eos-layout/internal/service"
	"eos-layout/pkg/log"
	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewAreaHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewAreaService,
)

var RepositorySet = wire.NewSet(
	repository.NewRepository,
	repository.NewAreaRepository,
)

func newApp(cfg *config.Config, logger *log.Logger) (server.Server, func(), error) {
	panic(wire.Build(RepositorySet, ServiceSet, HandlerSet, server.NewHTTPServer))
}
