package ethtransactions

import (
	"context"
	"log/slog"
)

// Manager — manager for transactions
type Manager struct {
	log  *slog.Logger
	trxr TransactionsRepository
}

// NewManager — constructor for Manager
func NewManager(trxr TransactionsRepository, l *slog.Logger) (*Manager, error) {
	m := &Manager{log: l, trxr: trxr}

	return m, nil
}

// GetTransaction — get data of the transaction
func (m *Manager) GetTransaction(ctx context.Context, hash string) (*Transaction, error) {
	log := m.log.With("hash", hash)

	trx, err := m.trxr.GetTransaction(ctx, hash)
	if err != nil {
		log.Error("get transaction", "error", err)
		return nil, ErrGetTransaction
	}

	log.Debug("GetTransaction", "to", trx.To, "from", trx.From, "value", trx.ValueInEther(), "status", trx.Status, "timestamp", trx.Timestamp)

	return trx, nil
}

// CreateTransaction – create a new transaction (sending ETH)
func (m *Manager) CreateTransaction(ctx context.Context, trx *Transaction) error {
	log := m.log.With("to", trx.To, "value", trx.ValueInEther())

	err := m.trxr.CreateTransaction(ctx, trx)
	if err != nil {
		log.Error("create transaction", "error", err)
		return ErrSendTransaction
	}

	m.log.Debug("CreateTransaction", "hash", trx.Hash, "from", trx.From, "to", trx.To, "value", trx.ValueInEther(), "status", trx.Status.String(), "timestamp", trx.Timestamp)

	return nil
}
