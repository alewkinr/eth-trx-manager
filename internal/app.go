package internal

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	stdhttp "net/http"

	"github.com/alewkinr/eth-trx-manager/internal/config"
	"github.com/alewkinr/eth-trx-manager/internal/ethtransactions"
	"github.com/alewkinr/eth-trx-manager/internal/ethwallet"
	"github.com/alewkinr/eth-trx-manager/internal/http"
	"github.com/alewkinr/eth-trx-manager/pkg/ethereum"
	"github.com/alewkinr/eth-trx-manager/pkg/logger"
)

// Application — main application container
type Application struct {
	cfg *config.Config
	log *slog.Logger
	srv *stdhttp.Server

	walletAPI       http.Router
	transactionsAPI http.Router

	ethCloseFunc func()
}

// NewApplication — constructor for application
func NewApplication() (*Application, error) {
	app := &Application{}
	app.cfg = config.MustNewConfig()
	app.log = logger.New(app.cfg.Log.Level)

	ethClient, closeFunc, connEthClientErr := ethereum.NewClient(app.cfg.Ethereum.URL)
	if connEthClientErr != nil {
		app.log.Error("ethereum client connection", "error", connEthClientErr)
		return nil, fmt.Errorf("ethereum client connection: %w", connEthClientErr)
	}
	app.ethCloseFunc = closeFunc

	walletMngr := ethwallet.NewManager(ethClient, app.log)
	app.walletAPI = http.NewWalletsAPIController(http.NewWalletsAPIService(walletMngr))

	trxMngr, newTrxMngrErr := ethtransactions.NewManager(ethClient, app.log, app.cfg.Ethereum.PrivateKey)
	if newTrxMngrErr != nil {
		app.log.Error("transactions manager build", "error", newTrxMngrErr)
		return nil, fmt.Errorf("transactions manager build: %w", newTrxMngrErr)
	}
	app.transactionsAPI = http.NewTransactionsAPIController(http.NewTransactionsAPIService(trxMngr))

	app.srv = &stdhttp.Server{
		Addr:    fmt.Sprintf("%s:%s", app.cfg.Host, app.cfg.Port),
		Handler: http.NewRouter(app.walletAPI, app.transactionsAPI),
	}

	return app, nil
}

// Run — run application
func (a *Application) Run() error {
	defer a.ethCloseFunc()

	a.log.Info("✅ Server is running...", "host", a.cfg.Host, "port", a.cfg.Port)
	if runErr := a.srv.ListenAndServe(); !errors.Is(runErr, stdhttp.ErrServerClosed) {
		return fmt.Errorf("http server start: %w", runErr)
	}

	return nil
}

func (a *Application) Stop(ctx context.Context) {
	a.log.Info("🛑 Server is shutting down...")
	if shutdownErr := a.srv.Shutdown(ctx); shutdownErr != nil {
		a.log.Error("http server shutdown", "error", shutdownErr)
	}
}
