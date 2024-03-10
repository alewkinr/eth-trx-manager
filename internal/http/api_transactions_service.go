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
	"math/big"
	"net/http"

	"github.com/alewkinr/eth-trx-manager/internal/ethtransactions"
)

// TransactionsAPIService is a service that implements the logic for the TransactionsAPIServicer
// This service should implement the business logic for every endpoint for the TransactionsAPI API.
// Include any external packages or services that will be required by this service.
type TransactionsAPIService struct {
	trxm *ethtransactions.Manager
}

// NewTransactionsAPIService creates a default api service
func NewTransactionsAPIService(trxm *ethtransactions.Manager) TransactionsAPIServicer {
	return &TransactionsAPIService{
		trxm: trxm,
	}
}

// AddTrx - CreateTransaction
func (s *TransactionsAPIService) AddTrx(ctx context.Context, request CreateTransactionRequest) (ImplResponse, error) {
	trx := &ethtransactions.Transaction{
		To:    request.To,
		Value: big.NewInt(request.Value),
	}

	createTrxErr := s.trxm.CreateTransaction(ctx, trx)
	if createTrxErr != nil {
		return Response(http.StatusInternalServerError, &ErrInternalError{InternalErrorMessage}), createTrxErr
	}

	return Response(http.StatusOK, &Transaction{
		Hash:      trx.Hash,
		From:      trx.From,
		To:        trx.To,
		Value:     trx.ValueInEther(),
		Status:    trx.Status.String(),
		Timestamp: trx.Timestamp,
	}), nil
}

// GetByTrxId - GetTransaction
func (s *TransactionsAPIService) GetByTrxId(ctx context.Context, hash string) (ImplResponse, error) {
	updatedTrx, getTrxErr := s.trxm.GetTransaction(ctx, hash)
	if getTrxErr != nil {
		return Response(http.StatusInternalServerError, &ErrInternalError{InternalErrorMessage}), getTrxErr
	}

	if updatedTrx == nil {
		return Response(http.StatusNoContent, nil), nil
	}

	return Response(http.StatusOK, &Transaction{
		Hash:      updatedTrx.Hash,
		From:      updatedTrx.From,
		To:        updatedTrx.To,
		Value:     updatedTrx.ValueInEther(),
		Status:    updatedTrx.Status.String(),
		Timestamp: updatedTrx.Timestamp,
	}), nil
}
