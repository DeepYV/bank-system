package usecase

import (
	domain "github.com/banking/domain"
)

type bankingucase struct {
	bankingRepo domain.BankingRepository
}

// inject the dependency for stickerusecase
func NewBankingUsecase(repo domain.BankingRepository) domain.BankingUseCase {
	return &bankingucase{
		bankingRepo: repo,
	}
}

// CreateAccount service layer
func (bu *bankingucase) CreateAccount(user domain.Account) error {

	err := bu.bankingRepo.CreateAccount(user)
	if err != nil {

		return err
	}

	return nil
}

// Transfer transaction layer
func (bu *bankingucase) Transfer(Trasaction domain.Transfer) error {

	err := bu.bankingRepo.Transfer(Trasaction)
	if err != nil {

		return err

	}

	return nil
}

// GetHistory Fetch all the transaction based on ID
func (bu *bankingucase) GetHistory(id string) (*[]domain.Entries, error) {

	history, err := bu.bankingRepo.GetHistory(id)
	if err != nil {

		return nil, err
	}

	return history, nil
}

// GetBalance check the balance in the account
func (bu *bankingucase) GetBalance(id string) (*int, error) {

	balance, err := bu.bankingRepo.GetBalance(id)
	if err != nil {

		return balance, err
	}

	return balance, nil
}
