package otelx

import (
	"os"

	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
	provider *sdklog.LoggerProvider
}

func NewLogger() *Logger {
	localCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	logger := zap.New(
		zapcore.NewTee(localCore),
		zap.AddCaller(),
	)

	return &Logger{
		Logger: logger,
	}
}

func NewNopLogger() *Logger {
	return &Logger{
		Logger: zap.NewNop(),
	}
}
