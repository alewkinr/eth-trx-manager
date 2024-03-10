package ethwallet

import "context"

type WalletRepository interface {
	// GetWalletBalance returns the balance of the wallet
	GetWalletBalance(ctx context.Context, address string) (*Wallet, error)
}
