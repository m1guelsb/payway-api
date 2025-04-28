package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         string
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number         string
	CVV            string
	ExpiryMonth    int
	ExpiryYear     int
	CardHolderName string
}

func NewInvoice(accountID string, amount float64, description string, paymentType string, card CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	lastFourDigits := card.Number[len(card.Number)-4:]

	return &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Amount:         amount,
		Status:         string(StatusPending),
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastFourDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (i *Invoice) Process() error {
	if i.Amount > 10000 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().Unix()))
	var newStatus Status

	if randomSource.Float64() <= 0.7 {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}

	i.Status = string(newStatus)
	return nil
}

func (i *Invoice) UpdateStatus(newStatus Status) error {
	if i.Status != string(StatusPending) {
		return ErrInvalidStatus
	}

	i.Status = string(newStatus)
	i.UpdatedAt = time.Now()
	return nil
}
