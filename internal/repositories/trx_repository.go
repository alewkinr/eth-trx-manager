package repositories

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/alewkinr/eth-trx-manager/internal/ethtransactions"
	"github.com/alewkinr/eth-trx-manager/pkg/cache"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// stdGasLimit — standard gas limit for transaction
const stdGasLimit = uint64(21000)

type TransactionsRepository struct {
	store           cache.Cache
	privateKeyECDSA *ecdsa.PrivateKey
	signKeyECDSA    *ecdsa.PublicKey
	ethClient       *ethclient.Client
}

// NewTransactionsRepository — constructor for TransactionsRepository
func NewTransactionsRepository(ethClient *ethclient.Client, privateKeyHex string, store cache.Cache) (*TransactionsRepository, error) {
	pk, parseKeyErr := crypto.HexToECDSA(privateKeyHex)
	if parseKeyErr != nil {
		return nil, fmt.Errorf("parse private key: %w", parseKeyErr)
	}

	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("convert public key to ECDSA")
	}

	return &TransactionsRepository{
		privateKeyECDSA: pk,
		signKeyECDSA:    publicKeyECDSA,
		ethClient:       ethClient,
		store:           store,
	}, nil
}

func (r *TransactionsRepository) GetTransaction(ctx context.Context, hash string) (*ethtransactions.Transaction, error) {
	if trx, isCached := r.store.Get(hash); isCached {
		return trx.(*ethtransactions.Transaction), nil
	}

	trxHash := common.HexToHash(hash)

	ethTrx, isPending, getTrxErr := r.ethClient.TransactionByHash(ctx, trxHash)
	if getTrxErr != nil {
		if errors.Is(getTrxErr, ethereum.NotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("get transaction by hash: %w", getTrxErr)
	}

	from, getFromErr := types.Sender(types.LatestSignerForChainID(ethTrx.ChainId()), ethTrx)
	if getFromErr != nil {
		return nil, fmt.Errorf("get transaction sender: %w", getFromErr)
	}

	var trxStatus ethtransactions.Status
	if isPending {
		trxStatus = ethtransactions.StatusIsPending
	}
	if !isPending {
		trxStatus = ethtransactions.StatusDone
	}

	domainTrx := &ethtransactions.Transaction{
		Hash:      hash,
		From:      from.String(),
		To:        ethTrx.To().String(),
		Value:     ethTrx.Value(),
		Timestamp: ethTrx.Time(),
		Status:    trxStatus,
	}

	r.store.Add(domainTrx.Hash, domainTrx)

	return domainTrx, nil
}

func (r *TransactionsRepository) CreateTransaction(ctx context.Context, trx *ethtransactions.Transaction) error {
	trx.From = crypto.PubkeyToAddress(*r.signKeyECDSA).String()

	gasPrice, err := r.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("receive gas price: %w", err)
	}

	nonce, err := r.ethClient.PendingNonceAt(ctx, common.HexToAddress(trx.From))
	if err != nil {
		return fmt.Errorf("receive nonce: %w", err)
	}

	ethTrx := types.NewTransaction(
		nonce,
		common.HexToAddress(trx.To),
		trx.Value,
		stdGasLimit,
		gasPrice,
		nil,
	)

	chainID, err := r.ethClient.NetworkID(ctx)
	if err != nil {
		return fmt.Errorf("receive chain id: %w", err)
	}

	signedTx, err := types.SignTx(ethTrx, types.NewEIP155Signer(chainID), r.privateKeyECDSA)
	if err != nil {
		return fmt.Errorf("sign transaction: %w", err)
	}

	if err := r.ethClient.SendTransaction(ctx, signedTx); err != nil {
		return fmt.Errorf("send transaction: %w", err)
	}

	trx.Hash = signedTx.Hash().String()
	trx.Timestamp = signedTx.Time()
	trx.Status = ethtransactions.StatusIsPending

	return nil
}
