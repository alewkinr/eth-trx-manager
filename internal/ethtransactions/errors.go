package ethtransactions

import "github.com/pkg/errors"

var (
	ErrReceiveGasPrice          = errors.New("failed to receive gas price")
	ErrReceiveNonce             = errors.New("failed to receive nonce")
	ErrReceiveChainID           = errors.New("failed to receive chain id")
	ErrSignTransaction          = errors.New("failed to sign transaction")
	ErrSendTransaction          = errors.New("failed to send transaction")
	ErrCreateTransactionManager = errors.New("failed to create transactions manager")
)
