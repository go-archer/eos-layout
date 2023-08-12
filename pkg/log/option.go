package log

type Config struct {
	File       string `toml:"file" json:"file" yaml:"file"`                      // 日志文件路径 out.log
	Level      string `toml:"level" json:"level" yaml:"level"`                   // 日志级别
	MaxSize    int    `toml:"max_size" json:"max_size" yaml:"max_size"`          // 每个日志文件最大大小 单位：M
	MaxAge     int    `toml:"max_age" json:"max_age" yaml:"max_age"`             // 日志文件最大保存天数
	MaxBackups int    `toml:"max_backups" json:"max_backups" yaml:"max_backups"` // 可以为日志文件保存的最大备份数
	Compress   bool   `toml:"compress" json:"compress" yaml:"compress"`          // 是否压缩
	Encoding   string `toml:"encoding" json:"encoding" yaml:"encoding"`          // 日志编码 console | json
	Prod       bool   `toml:"prod" json:"prod" yaml:"prod"`                      // 是否生产环境
}

var (
	defaultConfig = &Config{
		File:       "./logs/info.log",
		MaxSize:    50,
		MaxAge:     30,
		Level:      "info",
		MaxBackups: 0,
		Compress:   false,
		Encoding:   "console",
		Prod:       false,
	}
)

func check(cfg *Config) *Config {
	if len(cfg.File) == 0 {
		cfg.File = defaultConfig.File
	}
	if cfg.MaxSize <= 0 {
		cfg.MaxSize = defaultConfig.MaxSize
	}
	if cfg.MaxAge <= 0 {
		cfg.MaxAge = defaultConfig.MaxAge
	}
	if len(cfg.Level) == 0 {
		cfg.Level = defaultConfig.Level
	}
	if cfg.MaxBackups < 0 {
		cfg.MaxBackups = defaultConfig.MaxBackups
	}
	if len(cfg.Encoding) == 0 {
		cfg.Encoding = defaultConfig.Encoding
	}
	return cfg
}

func NewConfig(options ...Option) *Config {
	for _, opt := range options {
		opt(defaultConfig)
	}
	return defaultConfig
}

type Option func(c *Config)

func Filename(filename string) Option {
	return func(c *Config) {
		c.File = filename
	}
}

func Level(level string) Option {
	return func(c *Config) {
		c.Level = level
	}
}

func MaxSize(maxSize int) Option {
	return func(c *Config) {
		c.MaxSize = maxSize
	}
}

func MaxAge(maxAge int) Option {
	return func(c *Config) {
		c.MaxAge = maxAge
	}
}

func MaxBackups(maxBackups int) Option {
	return func(c *Config) {
		c.MaxBackups = maxBackups
	}
}

func Compress(compress bool) Option {
	return func(c *Config) {
		c.Compress = compress
	}
}

func Encoding(encoding string) Option {
	return func(c *Config) {
		c.Encoding = encoding
	}
}

func Prod(prod bool) Option {
	return func(c *Config) {
		c.Prod = prod
	}
}
