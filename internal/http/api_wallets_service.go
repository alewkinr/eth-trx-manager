/*
 * Ethereum transactions manager
 *
 * Ethereum transactions manager
 *
 * API version: 1.0.0
 * Contact: alewkinr@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package http

import (
	"context"
	"net/http"

	"github.com/alewkinr/eth-trx-manager/internal/ethwallet"
	"github.com/pkg/errors"
)

// WalletsAPIService is a service that implements the logic for the WalletsAPIServicer
// This service should implement the business logic for every endpoint for the WalletsAPI API.
// Include any external packages or services that will be required by this service.
type WalletsAPIService struct {
	wm *ethwallet.Manager
}

// NewWalletsAPIService creates a default api service
func NewWalletsAPIService(wm *ethwallet.Manager) WalletsAPIServicer {
	return &WalletsAPIService{wm}
}

// GetEthBalanceById - GetWalletBalance
func (s *WalletsAPIService) GetEthBalanceById(ctx context.Context, address string) (ImplResponse, error) {
	wallet := &ethwallet.Wallet{}
	wallet.SetAddress(address)

	refreshBalanceErr := s.wm.RefreshBalance(ctx, wallet)

	switch {
	case errors.Is(refreshBalanceErr, nil):
		return Response(http.StatusOK, Wallet{
			Address: wallet.Address().String(),
			Balance: wallet.Balance().String(),
		}), nil

	case errors.Is(refreshBalanceErr, ethwallet.ErrInvalidAddress):
		return Response(http.StatusBadRequest, ErrBadRequest{Message: refreshBalanceErr.Error()}), refreshBalanceErr

	default:
		return Response(http.StatusInternalServerError, ErrInternalError{Message: refreshBalanceErr.Error()}), refreshBalanceErr
	}
	// TODO: Add api_wallets_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
}
