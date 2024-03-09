package logger

import (
	"log/slog"

	"go.mrchanchal.com/zaphandler"
	"go.uber.org/zap"
)

func New() *slog.Logger {
	zapL, _ := zap.NewProduction()
	defer func() { _ = zapL.Sync() }()

	logger := slog.New(zaphandler.New(zapL))
	slog.SetDefault(logger)

	return logger
}
