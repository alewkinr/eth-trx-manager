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
	"github.com/alewkinr/eth-trx-manager/pkg/ethereum"
	"github.com/ethereum/go-ethereum/common"
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
	trx := &ethtransactions.Transaction{}
	trx.SetValue(big.NewInt(request.Value))

	to := common.HexToAddress(request.To) // todo: fix somehow!!!
	trx.SetTo(&to)

	updatedTrx, createTrxErr := s.trxm.CreateTransaction(ctx, trx)
	if createTrxErr != nil {
		return Response(http.StatusInternalServerError, nil), createTrxErr
	}

	// TODO - update AddTrx with the required logic for this service method.
	// Add api_transactions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Transaction{}) or use other options such as http.Ok ...
	// return Response(200, Transaction{}), nil

	// TODO: Uncomment the next line to return response Response(4XX, ErrBadRequest{}) or use other options such as http.Ok ...
	// return Response(4XX, ErrBadRequest{}), nil

	// TODO: Uncomment the next line to return response Response(5XX, ErrInternalError{}) or use other options such as http.Ok ...
	// return Response(5XX, ErrInternalError{}), nil

	return Response(http.StatusOK, Transaction{
		Hash:      updatedTrx.Hash().String(),
		From:      updatedTrx.From().String(),
		To:        updatedTrx.To().String(),
		Value:     ethereum.ToDecimal(updatedTrx.Value(), 18).Text('f', 18),
		Status:    updatedTrx.Status().String(),
		Timestamp: updatedTrx.Timestamp(),
	}), nil
}

// GetByTrxId - GetTransaction
func (s *TransactionsAPIService) GetByTrxId(ctx context.Context, hash string) (ImplResponse, error) {
	trx := &ethtransactions.Transaction{}
	trx.SetHash(hash)

	updatedTrx, getTrxErr := s.trxm.GetTransaction(ctx, trx)
	if getTrxErr != nil {
		return Response(http.StatusInternalServerError, ErrInternalError{Message: "internal error"}), getTrxErr
	}

	if updatedTrx == nil {
		return Response(http.StatusNoContent, nil), nil
	}
	// TODO - update GetByTrxId with the required logic for this service method.
	// Add api_transactions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Transaction{}) or use other options such as http.Ok ...
	// return Response(200, Transaction{}), nil

	// TODO: Uncomment the next line to return response Response(4XX, ErrBadRequest{}) or use other options such as http.Ok ...
	// return Response(4XX, ErrBadRequest{}), nil

	// TODO: Uncomment the next line to return response Response(5XX, ErrInternalError{}) or use other options such as http.Ok ...
	// return Response(5XX, ErrInternalError{}), nil

	return Response(http.StatusOK, Transaction{
		Hash:      updatedTrx.Hash().String(),
		From:      updatedTrx.From().String(),
		To:        updatedTrx.To().String(),
		Value:     ethereum.ToDecimal(updatedTrx.Value(), 18).String(),
		Status:    updatedTrx.Status().String(),
		Timestamp: updatedTrx.Timestamp(),
	}), nil
}
