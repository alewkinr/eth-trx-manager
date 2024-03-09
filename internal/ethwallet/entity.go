package ethwallet

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// Wallet — Ethereum wallet entity
type Wallet struct {
	*sync.RWMutex

	// address — address of the wallet in string type
	addressStr string
	// balance — current balance of the wallet
	balance *big.Int
}

// Balance — balance getter
func (w *Wallet) Balance() *big.Int {
	return w.balance
}

// SetBalance — balance setter
func (w *Wallet) SetBalance(balance *big.Int) {
	w.Lock()
	w.balance = balance
	w.Unlock()
}

// Address — getter for address field
func (w *Wallet) Address() common.Address {
	return common.HexToAddress(w.addressStr)
}

// SetAddress — setter for address field
func (w *Wallet) SetAddress(address string) {
	w.Lock()
	w.addressStr = address
	w.Unlock()
}
