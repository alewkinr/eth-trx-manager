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
)

// TransactionsAPIRouter defines the required methods for binding the api requests to a responses for the TransactionsAPI
// The TransactionsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a TransactionsAPIServicer to perform the required actions, then write the service results to the http response.
type TransactionsAPIRouter interface {
	AddTrx(http.ResponseWriter, *http.Request)
	GetByTrxId(http.ResponseWriter, *http.Request)
}

// WalletsAPIRouter defines the required methods for binding the api requests to a responses for the WalletsAPI
// The WalletsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a WalletsAPIServicer to perform the required actions, then write the service results to the http response.
type WalletsAPIRouter interface {
	GetEthBalanceById(http.ResponseWriter, *http.Request)
}

// TransactionsAPIServicer defines the api actions for the TransactionsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type TransactionsAPIServicer interface {
	AddTrx(context.Context, CreateTransactionRequest) (ImplResponse, error)
	GetByTrxId(context.Context, string) (ImplResponse, error)
}

// WalletsAPIServicer defines the api actions for the WalletsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type WalletsAPIServicer interface {
	GetEthBalanceById(context.Context, string) (ImplResponse, error)
}
