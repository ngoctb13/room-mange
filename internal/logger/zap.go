package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	var (
		level       = zapcore.InfoLevel
		development = false
	)

	if os.Getenv("DEBUG") != "" {
		level = zap.DebugLevel
		development = true
	}

	zapLogger := zap.Must(zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: development,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "name",
			MessageKey:    "message",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
			EncodeName:    zapcore.FullNameEncoder,
		},
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
		DisableStacktrace: true,
	}.Build())

	zap.ReplaceGlobals(zapLogger)
	zap.RedirectStdLog(zapLogger)

	return zapLogger
}
