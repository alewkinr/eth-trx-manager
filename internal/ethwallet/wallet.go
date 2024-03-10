package ethwallet

import (
	"math/big"

	"github.com/alewkinr/eth-trx-manager/pkg/ethereum"
)

// Wallet — Ethereum wallet entity
type Wallet struct {
	Address string

	Balance *big.Float
}

// Validate — validate info method
func (w *Wallet) Validate() error {
	if !ethereum.IsValidAddress(w.Address) || ethereum.IsZeroAddress(w.Address) {
		return ErrInvalidAddress
	}
	return nil
}

// FormattedBalance — getter for formatted balance value (i.e. 0.0000000000000000001)
func (w *Wallet) FormattedBalance() string {
	return w.Balance.Text('f', 18)
}
