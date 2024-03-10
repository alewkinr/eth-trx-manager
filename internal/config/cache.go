package config

import "time"

// Cache — cache settings
type Cache struct {
	WalletsSize      int           `koanf:"wallets_size"`
	WalletsTTL       time.Duration `koanf:"wallets_ttl"`
	TransactionsSize int           `koanf:"trx_size"`
	TransactionsTTL  time.Duration `koanf:"trx_ttl"`
}
