package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByApiKey(apiKey string) (*Account, error)
	FindByID(id string) (*Account, error)
	UpdateBalance(account *Account) error
}
