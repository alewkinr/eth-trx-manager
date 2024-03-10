package ethereum

import (
	"math"
	"math/big"
	"reflect"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// addressCheckRegexp — regexp for address validation check
var addressCheckRegexp = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	switch v := iaddress.(type) {
	case string:
		return addressCheckRegexp.MatchString(v)
	case common.Address:
		return addressCheckRegexp.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) *big.Float {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	floatb := new(big.Float)
	floatb.SetString(value.String())
	floatVal := new(big.Float).Quo(floatb, big.NewFloat(math.Pow10(decimals)))

	return floatVal
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}
