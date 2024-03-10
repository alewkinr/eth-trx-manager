package ethtransactions

import "github.com/pkg/errors"

var (
	ErrGetTransaction  = errors.New("failed to get transaction")
	ErrSendTransaction = errors.New("failed to send transaction")
)
