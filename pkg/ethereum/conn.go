package ethereum

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

// NewClient — creates new Ethereum client
func NewClient(nodeURL string) (*ethclient.Client, func(), error) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return nil, nil, fmt.Errorf("connect to Ethereum node: %w", err)
	}

	return client, client.Close, nil
}
