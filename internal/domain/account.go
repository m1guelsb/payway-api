package domain

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	ApiKey    string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateApiKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func NewAccount(name, email string) *Account {
	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Balance:   0,
		ApiKey:    generateApiKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}
