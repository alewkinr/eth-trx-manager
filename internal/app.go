package internal

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	stdhttp "net/http"

	"github.com/alewkinr/eth-trx-manager/internal/config"
	"github.com/alewkinr/eth-trx-manager/internal/http"
)

type Application struct {
	cfg *config.Config
	log *slog.Logger
	srv *stdhttp.Server

	balanceAPI http.Router
}

// NewApplication — constructor for application
func NewApplication(
	lgr *slog.Logger,
) *Application {
	app := &Application{}
	app.cfg = config.MustNewConfig()

	app.balanceAPI = http.NewWalletsAPIController(http.NewWalletsAPIService())

	app.srv = &stdhttp.Server{
		Addr:    fmt.Sprintf("%s:%s", app.cfg.Host, app.cfg.Port),
		Handler: http.NewRouter(app.balanceAPI),
	}

	return app
}

// Run — run application
func (a *Application) Run() error {
	if runErr := a.srv.ListenAndServe(); !errors.Is(runErr, stdhttp.ErrServerClosed) {
		return fmt.Errorf("http server start: %w", runErr)
	}

	return nil
}

func (a *Application) Stop(ctx context.Context) {
	if shutdownErr := a.srv.Shutdown(ctx); shutdownErr != nil {
		a.log.Error("http server shutdown", "error", shutdownErr)
	}
}
