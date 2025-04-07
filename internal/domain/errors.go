package domain

import "errors"

var (
	// ErrAccountNotFound is returned when an account is not found in the database.
	ErrAccountNotFound = errors.New("account not found")
	// ErrDuplicatedAPIKey is returned when an account with the same API key already exists.
	ErrDuplicatedAPIKey = errors.New("api key already exists")
	// ErrInvoiceNotFound is returned when an invoice is not found in the database.
	ErrInvoiceNotFound = errors.New("invoice not found")
	// ErrUnauthorized is returned when an API key is not authorized to access a resource.
	ErrUnauthorized = errors.New("unauthorized access")
)
