package main

import (
	"os"

	"github.com/alewkinr/eth-trx-manager/internal"
	"github.com/alewkinr/eth-trx-manager/pkg/graceful"
	"github.com/alewkinr/eth-trx-manager/pkg/logger"
)

const (
	exitCodeOK = iota
	exitCodeNotOK
)

// run — wraps main function to run application and returning exit code.
func run() int {
	slog := logger.New("debug") // todo: debug off

	app := internal.NewApplication(slog)
	if app == nil {
		return exitCodeNotOK
	}

	go graceful.ShutdownMonitor(app.Stop)

	if runErr := app.Run(); runErr != nil {
		slog.Error("run app", "error", runErr)
		return exitCodeNotOK
	}

	return exitCodeOK
}

func main() {
	os.Exit(run())
}
