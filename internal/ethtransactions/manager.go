package ethtransactions

import (
	"context"
	"log/slog"

	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
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
	m := &Manager{log: l, ethClient: ec}

	pk, parseKeyErr := crypto.HexToECDSA(privateKeyHex)
	if parseKeyErr != nil {
		m.log.Error("parse private key", "error", parseKeyErr)
		return nil, ErrCreateTransactionManager
	}

	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		m.log.Error("type converting public key not ok", "publicKey", publicKey)
		return nil, ErrCreateTransactionManager
	}

	m.privateKeyECDSA = pk
	m.signKeyECDSA = publicKeyECDSA

	return m, nil
}

// GetTransaction — get data of the transaction
func (m *Manager) GetTransaction(ctx context.Context, trx *Transaction) (*Transaction, error) {
	log := m.log.With("hash", trx.Hash().String())

	ethTrx, isPending, getTrxErr := m.ethClient.TransactionByHash(ctx, trx.Hash())
	if getTrxErr != nil {
		if errors.Is(getTrxErr, ethereum.NotFound) {
			log.Debug("transaction not found", "error", getTrxErr)
			return nil, nil
		}

		log.Error("get transaction", "error", getTrxErr)
		return nil, ErrGetTransaction
	}

	from, getFromErr := m.getFromAddress(ethTrx)
	if getFromErr != nil {
		log.Error("get from address", "error", getFromErr)
		return nil, ErrGetTransaction
	}

	trx.SetTo(ethTrx.To())
	trx.SetFrom(from)
	trx.SetValue(ethTrx.Value())
	trx.SetTimestamp(ethTrx.Time())
	trx.SetStatus(isPending)

	log.Debug("GetTransaction", "to", trx.To(), "from", trx.From(), "value", trx.Value(), "status", trx.Status(), "timestamp", trx.Timestamp())

	return trx, nil
}

// GetFromAddress — getter for from address of transaction
func (m *Manager) getFromAddress(tx *types.Transaction) (common.Address, error) {
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	return from, err
}

// stdGasLimit — standard gas limit for transaction
const stdGasLimit = uint64(21000)

// CreateTransaction – create a new transaction (sending ETH)
func (m *Manager) CreateTransaction(ctx context.Context, trx *Transaction) (*Transaction, error) {
	log := m.log.With("to", trx.To().String(), "value", trx.Value().String())
	trx.SetFrom(crypto.PubkeyToAddress(*m.signKeyECDSA))

	gasPrice, err := m.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		log.Error("suggest gas price", "error", err)
		return nil, ErrReceiveGasPrice
	}

	nonce, err := m.ethClient.PendingNonceAt(ctx, trx.From())
	if err != nil {
		log.Error("receive nonce", "error", err)
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

	chainID, err := m.ethClient.NetworkID(ctx)
	if err != nil {
		log.Error("receive chain id", "error", err)
		return nil, ErrReceiveChainID
	}

	signedTx, err := types.SignTx(ethTrx, types.NewEIP155Signer(chainID), m.privateKeyECDSA)
	if err != nil {
		log.Error("sign transaction", "error", err)
		return nil, ErrSignTransaction
	}

	if err := m.ethClient.SendTransaction(ctx, signedTx); err != nil {
		log.Error("send transaction", "error", err, "from", trx.From().String())
		return nil, ErrSendTransaction
	}

	trx.SetHash(signedTx.Hash().String())
	trx.SetTimestamp(signedTx.Time())
	trx.SetStatus(true)

	m.log.Debug("CreateTransaction", "hash", trx.Hash().String(), "from", trx.From().String(), "to", trx.To().String(), "value", trx.Value().String(), "status", trx.Status().String(), "timestamp", trx.Timestamp())

	return trx, nil
}
