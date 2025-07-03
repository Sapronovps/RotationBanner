package logger

import (
	"go.uber.org/zap"
	"strings"
)

func New(level, file string) *zap.Logger {
	loggerCfg := zap.NewProductionConfig()

	l := strings.ToLower(level)
	switch l {
	case "debug":
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "fatal":
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	case "panic":
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	default:
		panic("Unknown log level")
	}

	if file != "" {
		loggerCfg.OutputPaths = []string{"stderr"}
	} else {
		loggerCfg.OutputPaths = []string{"stderr", file}
	}
	logger, _ := loggerCfg.Build()
	return logger
}
