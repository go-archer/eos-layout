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

func newApp(cfg *config.Config, logger *log.Logger) (server.Server, func(), error) {
	panic(wire.Build(repository.Set, service.Set, handler.Set, server.NewHTTPServer))
}
