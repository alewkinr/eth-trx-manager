package ethtransactions

import (
	"context"
	"log/slog"

	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Manager — manager for transactions
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

// GetTransaction — get data of the transaction
func (m *Manager) GetTransaction(ctx context.Context, trx *Transaction) (*Transaction, error) {
	ethTrx, isPending, getTrxErr := m.ethClient.TransactionByHash(ctx, trx.Hash())
	if getTrxErr != nil {
		return nil, fmt.Errorf("get transaction: %w", getTrxErr)
	}

	from, getFromErr := m.getFromAddress(ethTrx)
	if getFromErr != nil {
		return nil, fmt.Errorf("get from address: %w", getFromErr)
	}

	trx.SetTo(ethTrx.To())
	trx.SetFrom(from)
	trx.SetValue(ethTrx.Value())
	trx.SetTimestamp(ethTrx.Time())
	trx.SetStatus(isPending)

	m.log.Debug("GetTransaction", "hash", trx.Hash(), "to", trx.To(), "from", trx.From(), "value", trx.Value(), "status", trx.Status(), "timestamp", trx.Timestamp())

	return trx, nil
}

// GetFromAddress — getter for from address of transaction
func (m *Manager) getFromAddress(tx *types.Transaction) (common.Address, error) {
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	return from, err
}
