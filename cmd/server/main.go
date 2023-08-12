package main

import (
	"encoding/json"
	"eos-layout/internal/config"
	"eos-layout/pkg/log"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New(config.WithFile("./config.toml"))
	err := cfg.Load()
	if err != nil {
		panic(err)
	}

	js, _ := json.Marshal(cfg)
	fmt.Println(string(js))
	logger := log.New(cfg.Log)
	app, cleanup, err := newApp(cfg, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	logger.Info("server start", zap.String("host", "http://"+cfg.Host))
	app.Run()
}
