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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// TransactionsAPIController binds http requests to an api service and writes the service results to the http response
type TransactionsAPIController struct {
	service      TransactionsAPIServicer
	errorHandler ErrorHandler
}

// TransactionsAPIOption for how the controller is set up.
type TransactionsAPIOption func(*TransactionsAPIController)

// WithTransactionsAPIErrorHandler inject ErrorHandler into controller
func WithTransactionsAPIErrorHandler(h ErrorHandler) TransactionsAPIOption {
	return func(c *TransactionsAPIController) {
		c.errorHandler = h
	}
}

// NewTransactionsAPIController creates a default api controller
func NewTransactionsAPIController(s TransactionsAPIServicer, opts ...TransactionsAPIOption) Router {
	controller := &TransactionsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the TransactionsAPIController
func (c *TransactionsAPIController) Routes() Routes {
	return Routes{
		"AddTrx": Route{
			strings.ToUpper("Post"),
			"/transactions",
			c.AddTrx,
		},
		"GetByTrxId": Route{
			strings.ToUpper("Get"),
			"/transactions/{hash}",
			c.GetByTrxId,
		},
	}
}

// AddTrx - CreateTransaction
func (c *TransactionsAPIController) AddTrx(w http.ResponseWriter, r *http.Request) {
	createTransactionRequestParam := CreateTransactionRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&createTransactionRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCreateTransactionRequestRequired(createTransactionRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCreateTransactionRequestConstraints(createTransactionRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddTrx(r.Context(), createTransactionRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetByTrxId - GetTransaction
func (c *TransactionsAPIController) GetByTrxId(w http.ResponseWriter, r *http.Request) {
	hashParam := chi.URLParam(r, "hash")
	if hashParam == "" {
		c.errorHandler(w, r, &RequiredError{"hash"}, nil)
		return
	}
	result, err := c.service.GetByTrxId(r.Context(), hashParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
