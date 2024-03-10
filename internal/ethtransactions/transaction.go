//go:generate go run github.com/abice/go-enum -f transaction.go --nocase --marshal --sql --sqlnullstr --forceupper
package ethtransactions

import (
	"math/big"
	"time"

	"github.com/alewkinr/eth-trx-manager/pkg/ethereum"
)

// Status — status of the transaction
// ENUM(is_pending, done).
type Status uint

// Transaction — Ethereum transaction entity
type Transaction struct {
	Hash string

	From string
	To   string

	// Value — amount of the transaction (in Wei). To get the value in Ether use ValueInEther() method
	Value  *big.Int
	Status Status

	Timestamp time.Time
}

// ValueInEther — getter for formatted balance value in Ether (i.e. 0.0000000000000000001)
func (t *Transaction) ValueInEther() string {
	return ethereum.ToDecimal(t.Value, 18).Text('f', 18)
}
