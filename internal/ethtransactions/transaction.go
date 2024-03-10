//go:generate go run github.com/abice/go-enum -f transaction.go --nocase --marshal --sql --sqlnullstr --forceupper
package ethtransactions

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Status — status of the transaction
// ENUM(is_pending, done).
type Status uint

// Transaction — Ethereum transaction entity
type Transaction struct {
	hash  string
	to    string
	from  string
	value *big.Int

	status    Status // TODO: add status to oapi spec
	timestamp time.Time
}

// SetHash — setter for hash field
func (t *Transaction) SetHash(hash string) {
	t.hash = hash
}

// Hash — getter for hash field
func (t *Transaction) Hash() common.Hash {
	return common.HexToHash(t.hash)
}

// SetTo — setter for to field
func (t *Transaction) SetTo(to *common.Address) {
	t.to = to.String()
}

// To — getter for to field
func (t *Transaction) To() common.Address {
	return common.HexToAddress(t.to)
}

// SetFrom — setter for from field
func (t *Transaction) SetFrom(sender common.Address) {
	t.from = sender.String()
}

// From — getter for from field
func (t *Transaction) From() common.Address {
	return common.HexToAddress(t.from)
}

// SetValue — setter for value field
func (t *Transaction) SetValue(v *big.Int) {
	t.value = v
}

// Value — getter for value field
func (t *Transaction) Value() *big.Int {
	return t.value
}

// SetTimestamp — setter for timestamp field
func (t *Transaction) SetTimestamp(tmstmp time.Time) {
	t.timestamp = tmstmp
}

// Timestamp — getter for timestamp field
func (t *Transaction) Timestamp() time.Time {
	return t.timestamp
}

// Status — getter for status field
func (t *Transaction) Status() Status {
	return t.status
}

// SetStatus — setting status field
func (t *Transaction) SetStatus(isPending bool) {
	if isPending {
		t.status = StatusIsPending
	}
	if !isPending {
		t.status = StatusDone
	}
}
