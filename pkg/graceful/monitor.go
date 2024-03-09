package graceful

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// defaultTimeout — default shutdown timeout
const defaultTimeout = 30 * time.Second

// ShutdownMonitor — monitors the shutdown signal and gracefully stops the application
func ShutdownMonitor(stopFunc func(ctx context.Context)) {
	stopping := make(chan os.Signal, 1)
	signal.Notify(stopping, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-stopping

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	stopFunc(ctx)
}
