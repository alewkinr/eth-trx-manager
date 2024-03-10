package ethwallet

import (
	"math/big"

	"github.com/alewkinr/eth-trx-manager/pkg/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Wallet — Ethereum wallet entity
type Wallet struct {
	// *sync.RWMutex // todo: uncomment?

	// address — address of the wallet in string type
	addressStr string
	// balance — current balance of the wallet
	balance *big.Float
}

// Validate — validate info method
func (w *Wallet) Validate() error {
	if !ethereum.IsValidAddress(w.addressStr) || ethereum.IsZeroAddress(w.addressStr) {
		return ErrInvalidAddress
	}
	return nil
}

// Balance — balance getter
func (w *Wallet) Balance() *big.Float {
	return w.balance
}

// SetBalance — balance setter
func (w *Wallet) SetBalance(balance *big.Int) {
	w.balance = ethereum.ToDecimal(balance, 18)
}

// Address — getter for address field
func (w *Wallet) Address() common.Address {
	return common.HexToAddress(w.addressStr)
}

// SetAddress — setter for address field
func (w *Wallet) SetAddress(address string) {
	w.addressStr = address
}
