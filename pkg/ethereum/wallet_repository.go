package ethereum

import (
	"context"
	"fmt"

	"github.com/alewkinr/eth-trx-manager/internal/ethwallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type WalletRepository struct {
	ethClient *ethclient.Client
}

// NewWalletRepository — constructor for WalletRepository
func NewWalletRepository(ethClient *ethclient.Client) *WalletRepository {
	return &WalletRepository{ethClient: ethClient}
}

func (w *WalletRepository) GetWalletBalance(ctx context.Context, address string) (*ethwallet.Wallet, error) {
	hexAddr := common.HexToAddress(address)

	balance, getBalanceErr := w.ethClient.BalanceAt(ctx, hexAddr, nil)
	if getBalanceErr != nil {
		return nil, fmt.Errorf("get balance: %w", getBalanceErr)
	}

	wallet := &ethwallet.Wallet{}
	wallet.Address = address
	wallet.Balance = ToDecimal(balance, 18)

	return wallet, nil
}
