package main

import (
	"encoding/json"
	"eos-layout/internal/config"
	"eos-layout/pkg/log"
	"flag"
	"fmt"
	"go.uber.org/zap"
)

var (
	configFile = "./config.toml"
)

func init() {
	flag.StringVar(&configFile, "c", "config.toml", "config file, eg: -c config.toml")
}

func main() {
	flag.Parse()
	cfg := config.New(config.WithFile(configFile))
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
