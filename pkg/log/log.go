package log

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const LOGGER_KEY = "LOGGER"

type Logger struct {
	*zap.Logger
}

func New(cfg *Config) *Logger {
	return initLogger(cfg)
}

func initLogger(cfg *Config) *Logger {
	if cfg == nil {
		cfg = defaultConfig
	}

	cfg = check(cfg)

	var level zapcore.Level

	switch cfg.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	hook := lumberjack.Logger{
		Filename:   cfg.File,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	var encoder zapcore.Encoder
	if cfg.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig())
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig())
	}
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), level)
	if !cfg.Prod {
		return &Logger{zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
	}
	return &Logger{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}

}

func encoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000000000"))
}

func (l *Logger) NewContext(ctx *gin.Context, fields ...zapcore.Field) {
	ctx.Set(LOGGER_KEY, l.WithContext(ctx).With(fields...))
}

func (l *Logger) WithContext(ctx *gin.Context) *Logger {
	if ctx == nil {
		return l
	}
	c, _ := ctx.Get(LOGGER_KEY)
	z, ok := c.(*zap.Logger)
	if ok {
		return &Logger{z}
	}
	return l
}
