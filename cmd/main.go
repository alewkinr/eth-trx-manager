package main

import (
	"os"

	"github.com/alewkinr/eth-trx-manager/internal"
	"github.com/alewkinr/eth-trx-manager/pkg/graceful"
)

const (
	exitCodeOK = iota
	exitCodeNotOK
)

// run — wraps main function to run application and returning exit code.
func run() int {
	app, createAppErr := internal.NewApplication()
	if app == nil || createAppErr != nil {
		return exitCodeNotOK
	}

	go graceful.ShutdownMonitor(app.Stop)

	if runErr := app.Run(); runErr != nil {
		return exitCodeNotOK
	}

	return exitCodeOK
}

func main() {
	os.Exit(run())
}
