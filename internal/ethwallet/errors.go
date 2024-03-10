package ethwallet

import "github.com/pkg/errors"

var (
	// ErrInvalidAddress — error for invalid wallet address
	ErrInvalidAddress       = errors.New("invalid wallet address")
	ErrReceiveWalletBalance = errors.New("failed to receive wallet balance")
)
