package ethtransactions

import "context"

type TransactionsRepository interface {
	CreateTransaction(ctx context.Context, trx *Transaction) error
	GetTransaction(ctx context.Context, hash string) (*Transaction, error)
}
