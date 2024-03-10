package ethwallet

import (
	"context"
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
	log := m.log.With("address", wallet.Address().String())
	if validateErr := wallet.Validate(); validateErr != nil {
		return ErrInvalidAddress
	}

	balance, getBalanceErr := m.ethClient.BalanceAt(ctx, wallet.Address(), nil)
	if getBalanceErr != nil {
		log.Error("get balance at", "address", wallet.Address(), "error", getBalanceErr)
		return ErrReceiveWalletBalance
	}

	wallet.SetBalance(balance)

	log.Debug("RefreshBalance", "balance", wallet.Balance())

	return nil
}
