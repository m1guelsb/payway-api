package domain

import "errors"

var (
	ErrAccountNotFound  = errors.New("account not found")
	ErrDuplicatedAPIKey = errors.New("api key already exists")
	ErrInvoiceNotFound  = errors.New("invoice not found")
	ErrUnauthorized     = errors.New("unauthorized access")
	ErrInvalidAmount    = errors.New("amount must be greater than zero")
	ErrInvalidStatus    = errors.New("invalid status, must be one of: pending, approved, rejected")
)
