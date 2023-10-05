//go:build wireinject
// +build wireinject

package main

import (
	"imall/internal/config"
	"imall/internal/handler"
	"imall/internal/repository"
	"imall/internal/server"
	"imall/internal/service"
	"imall/pkg/log"

	"github.com/google/wire"
)

func newApp(cfg *config.Config, logger *log.Logger) (server.Server, func(), error) {
	panic(wire.Build(repository.Set, service.Set, handler.Set, server.NewHTTPServer))
}
