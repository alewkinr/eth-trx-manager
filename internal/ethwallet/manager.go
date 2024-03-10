package ethwallet

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Manager — manager for wallets
type Manager struct {
	log       *slog.Logger
	ethClient *ethclient.Client
}

// NewManager — constructor for Manager
func NewManager(ec *ethclient.Client, l *slog.Logger) *Manager {
	return &Manager{
		log:       l,
		ethClient: ec,
	}
}

// RefreshBalance — get balance of the wallet from network, updates in wallet
func (m *Manager) RefreshBalance(ctx context.Context, wallet *Wallet) error {
	if validateErr := wallet.Validate(); validateErr != nil {
		return fmt.Errorf("validate wallet: %w", validateErr)
	}

	balance, getBalanceErr := m.ethClient.BalanceAt(ctx, wallet.Address(), nil)
	if getBalanceErr != nil {
		return fmt.Errorf("get balance: %w", getBalanceErr)
	}

	wallet.SetBalance(balance)

	m.log.Debug("RefreshBalance", "address", wallet.Address(), "balance", wallet.Balance())

	return nil
}
