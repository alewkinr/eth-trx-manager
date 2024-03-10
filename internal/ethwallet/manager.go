package ethwallet

import (
	"context"
	"log/slog"
)

// Manager — manager for wallets
type Manager struct {
	log *slog.Logger
	wr  WalletRepository
}

// NewManager — constructor for Manager
func NewManager(wr WalletRepository, l *slog.Logger) *Manager {
	return &Manager{
		log: l,
		wr:  wr,
	}
}

// RefreshBalance — get balance of the wallet from network, updates in wallet
func (m *Manager) RefreshBalance(ctx context.Context, address string) (*Wallet, error) {
	log := m.log.With("address", address)
	wallet := &Wallet{Address: address}

	if validateErr := wallet.Validate(); validateErr != nil {
		return nil, ErrInvalidAddress
	}

	walletWithBalance, getBalanceErr := m.wr.GetWalletBalance(ctx, wallet.Address)
	if getBalanceErr != nil {
		log.Error("get balance at", "error", getBalanceErr)
		return nil, ErrReceiveWalletBalance
	}

	wallet.Balance = walletWithBalance.Balance

	log.Debug("RefreshBalance", "balance", wallet.Balance)

	return wallet, nil
}
