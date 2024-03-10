package logger

import (
	"log/slog"

	"go.mrchanchal.com/zaphandler"
	"go.uber.org/zap"
)

func New(level string) *slog.Logger {
	config := zap.NewProductionConfig()

	lvl := zap.NewAtomicLevel()
	_ = lvl.UnmarshalText([]byte(level))
	config.Level = lvl

	zapL, _ := config.Build()
	defer func() { _ = zapL.Sync() }()

	logger := slog.New(zaphandler.New(zapL))
	slog.SetDefault(logger)

	return logger
}
