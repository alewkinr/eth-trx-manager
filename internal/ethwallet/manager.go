package ethwallet

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Manager — manager for wallets
type Manager struct {
	ethClient *ethclient.Client
}

// NewManager — constructor for Manager
func NewManager(ec *ethclient.Client) *Manager {
	return &Manager{
		ethClient: ec,
	}
}

// RefreshBalance — get balance of the wallet from network, updates in wallet
func (m *Manager) RefreshBalance(ctx context.Context, wallet *Wallet) error {
	balance, getBalanceErr := m.ethClient.BalanceAt(ctx, wallet.Address(), nil)
	if getBalanceErr != nil {
		return fmt.Errorf("get balance: %w", getBalanceErr)
	}

	wallet.SetBalance(balance)

	return nil
}
