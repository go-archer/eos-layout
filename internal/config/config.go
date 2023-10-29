package config

import (
	"encoding/json"
	"eos-layout/pkg/log"
	"errors"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

type Option func(c *Config)

func WithFile(src string) Option {
	return func(c *Config) {
		c.src = src
	}
}

type Config struct {
	src   string      // 配置文件
	Host  string      `toml:"host" json:"host" yaml:"host"`
	Debug bool        `toml:"debug" json:"debug" yaml:"debug"`
	Log   *log.Config `toml:"log" json:"log" yaml:"log"`
	MySQL *MySQL      `toml:"mysql" json:"mysql" yaml:"mysql"`
	Redis *Redis      `toml:"redis" json:"redis" yaml:"redis"`
}

type MySQL struct {
	Host        string `toml:"host" json:"host" yaml:"host"`
	Port        int    `toml:"port" json:"port" yaml:"port"`
	DBName      string `toml:"db_name" json:"db_name" yaml:"db_name"`
	Username    string `toml:"username" json:"username" yaml:"user"`
	Password    string `toml:"password" json:"password" yaml:"password"`
	MaxIDLE     int    `toml:"max_idle" json:"max_idle" yaml:"max_idle"` // 最大空闲数
	MaxOpen     int    `toml:"max_open" json:"max_open" yaml:"max_open"` // 最大连接数
	Timeout     int    `toml:"timeout" json:"timeout" yaml:"timeout"`    // 连接超时
	AutoMigrate bool   `toml:"auto_migrate" json:"auto_migrate" yaml:"auto_migrate"`
}

type Redis struct {
	Addr         string `toml:"addr" json:"addr" yaml:"addr"`
	Username     string `toml:"username" json:"username" yaml:"username"`
	Password     string `toml:"password" json:"password" yaml:"password"`
	DB           int    `toml:"db" json:"db" yaml:"db"`
	DialTimeout  int64  `toml:"dial_timeout" json:"dial_timeout" yaml:"dial_timeout"`
	ReadTimeout  int64  `toml:"read_timeout" json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout int64  `toml:"write_timeout" json:"write_timeout" yaml:"write_timeout"`
}

var cfg *Config = &Config{src: "./config.toml"}

func New(options ...Option) *Config {
	for _, opt := range options {
		opt(cfg)
	}
	return cfg
}

func (c *Config) Load() error {
	suffix, err := c.suffix()
	if err != nil {
		return errors.Join(err, errors.New("failed load file"))
	}
	f, err := os.Open(c.src)
	if err != nil {
		return err
	}
	defer f.Close()
	switch suffix {
	case "toml":
		return c.toml(f)
	case "json":
		return c.json(f)
	case "yaml":
		return c.yaml(f)
	}
	return nil
}

func (c *Config) toml(f *os.File) error {
	decoder := toml.NewDecoder(f)
	return decoder.Decode(c)
}

func (c *Config) yaml(f *os.File) error {
	decoder := yaml.NewDecoder(f)
	return decoder.Decode(c)
}

func (c *Config) json(f *os.File) error {
	decoder := json.NewDecoder(f)
	return decoder.Decode(c)
}

func (c *Config) suffix() (string, error) {
	if strings.HasSuffix(c.src, "toml") {
		return "toml", nil
	}
	if strings.HasSuffix(c.src, "json") {
		return "json", nil
	}
	if strings.HasSuffix(c.src, "yaml") || strings.HasSuffix(c.src, "yml") {
		return "yaml", nil
	}
	return "", errors.New("unsupported file")
}
