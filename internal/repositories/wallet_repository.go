package repositories

import (
	"context"
	"fmt"

	"github.com/alewkinr/eth-trx-manager/internal/ethwallet"
	"github.com/alewkinr/eth-trx-manager/pkg/cache"
	"github.com/alewkinr/eth-trx-manager/pkg/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type WalletRepository struct {
	store     cache.Cache
	ethClient *ethclient.Client
}

// NewWalletRepository — constructor for WalletRepository
func NewWalletRepository(ethClient *ethclient.Client, store cache.Cache) *WalletRepository {
	return &WalletRepository{ethClient: ethClient, store: store}
}

func (r *WalletRepository) GetWalletBalance(ctx context.Context, address string) (*ethwallet.Wallet, error) {
	if wallet, isCached := r.store.Get(address); isCached {
		return wallet.(*ethwallet.Wallet), nil
	}

	hexAddr := common.HexToAddress(address)
	balance, getBalanceErr := r.ethClient.BalanceAt(ctx, hexAddr, nil)
	if getBalanceErr != nil {
		return nil, fmt.Errorf("get balance: %w", getBalanceErr)
	}

	wallet := &ethwallet.Wallet{}
	wallet.Address = address
	wallet.Balance = ethereum.ToDecimal(balance, 18)

	r.store.Add(wallet.Address, wallet)

	return wallet, nil
}
