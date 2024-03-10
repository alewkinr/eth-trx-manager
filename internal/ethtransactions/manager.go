package ethtransactions

import (
	"context"
	"fmt"
	"log/slog"

	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Manager — manager for transactions
type Manager struct {
	privateKeyECDSA *ecdsa.PrivateKey
	signKeyECDSA    *ecdsa.PublicKey
	log             *slog.Logger
	ethClient       *ethclient.Client
}

// NewManager — constructor for Manager
func NewManager(ec *ethclient.Client, l *slog.Logger, privateKeyHex string) (*Manager, error) {
	pk, parseKeyErr := crypto.HexToECDSA(privateKeyHex)
	if parseKeyErr != nil {
		// todo: add logging
		return nil, ErrCreateTransactionManager
	}
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		// todo: add logging
		return nil, ErrCreateTransactionManager
	}

	return &Manager{
		privateKeyECDSA: pk,
		signKeyECDSA:    publicKeyECDSA,
		log:             l,
		ethClient:       ec,
	}, nil
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

// CreateTransaction – create a new transaction (sending ETH)
func (m *Manager) CreateTransaction(ctx context.Context, trx *Transaction) (*Transaction, error) {
	trx.SetFrom(crypto.PubkeyToAddress(*m.signKeyECDSA))

	stdGasLimit := uint64(3600)
	gasPrice, err := m.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		// todo: add logging
		return nil, ErrReceiveGasPrice
	}

	nonce, err := m.ethClient.PendingNonceAt(context.Background(), trx.From())
	if err != nil {
		// todo: add logging
		return nil, ErrReceiveNonce
	}

	ethTrx := types.NewTransaction(
		nonce,
		trx.To(),
		trx.Value(),
		stdGasLimit,
		gasPrice,
		nil,
	)

	chainID, err := m.ethClient.NetworkID(context.Background())
	if err != nil {
		return nil, ErrReceiveChainID
	}

	signedTx, err := types.SignTx(ethTrx, types.NewEIP155Signer(chainID), m.privateKeyECDSA)
	if err != nil {
		return nil, ErrSignTransaction
	}

	if err := m.ethClient.SendTransaction(ctx, signedTx); err != nil {
		return nil, ErrSendTransaction
	}

	trx.SetHash(signedTx.Hash().String())
	trx.SetTimestamp(signedTx.Time())
	trx.SetStatus(true)

	return trx, nil
}
