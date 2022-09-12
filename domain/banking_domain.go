package domain

import "time"

type Account struct {
	Id        string    `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int       `json:"balance"`
	Phone     string    `json:"phone"`
	Pin       string    `json:"pin"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	Id            string    `json:"id"`
	FromAccountId string    `json:"from_account_id"`
	ToAccountId   string    `json:"to_account_id"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
type Entries struct {
	Id         string    `json:"id"`
	Account_id string    `json:"account_id"`
	Status     string    `json:"status"`
	Amount     int       `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}

type Records struct {
	Account_id string
	Status     string
	Amount     int
	CreatedAt  time.Time
}

// default amount 
var DefaultAmount int = 50

// banking usecase
type BankingUseCase interface {
	CreateAccount(user Account) error
	Transfer(Trasaction Transfer) error
	GetBalance(AccountId string) (*int, error)
	GetHistory(AccountId string) (*[]Entries, error)
}

// banking repository interface
type BankingRepository interface {
	CreateAccount(user Account) error
	Transfer(Trasaction Transfer) error
	GetBalance(AccountId string) (*int, error)
	GetHistory(AccountId string) (*[]Entries, error)
}
