package repository

import (
	"errors"

	domain "github.com/banking/domain"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type PostgresBankingRepository struct {
	Conn *gorm.DB
}

// NewpostgresStickerRepository inject dependency to postgresStickerRepository
func NewPostgresBankingRepository(Conn *gorm.DB) domain.BankingRepository {
	return &PostgresBankingRepository{Conn}
}

// CreateAccount register user in the bank with balance 50.
func (bu *PostgresBankingRepository) CreateAccount(user domain.Account) error {

	err := bu.Conn.Create(&user).Error
	if err != nil {

		log.Error(err)

		return err
	}

	return nil
}


// Transfer make trasaction from sender to reciever then update the balance.
func (bu *PostgresBankingRepository) Transfer(Trasaction domain.Transfer) error {

	tx := bu.Conn.Begin()

	// extract ID of sender and reciever
	senderID := Trasaction.FromAccountId
	recieverID := Trasaction.ToAccountId
	amount := Trasaction.Amount
	senderUser, err := bu.GetUser(senderID)

	if err != nil {

		log.Error(err)

		return errors.New("Sender Account doesnt exists")
	}

	// Extract the user
	recieverUser, err1 := bu.GetUser(recieverID)

	if err1 != nil {

		log.Error(err1)

		return errors.New("reciever Account doesnt exists")
	}

	// check if enough balance is there
	if senderUser.Balance-amount < 0 {

		return errors.New("insufficient balance")
	}

	// update amount in both the user 
	senderUser.Balance = senderUser.Balance - amount
	recieverUser.Balance = recieverUser.Balance + amount

	err = bu.UpdateAmount(senderUser.Id, senderUser.Balance)
	if err != nil {

		return errors.New("cannot update amount")
	}
	err = bu.UpdateAmount(recieverUser.Id, recieverUser.Balance)
	if err != nil {

		return errors.New("cannot update reciever amount")
	}
	err = tx.Model(&domain.Transfer{}).Create(&Trasaction).Error

	if err != nil {

		return errors.New("transaction cannot be made")

	}

	// add the entries 
	var entries domain.Entries

	entries.Account_id = senderUser.Id
	entries.Status = "deduct"
	entries.Amount = amount
	err = bu.CreateEntries(entries)
	if err != nil {

		return errors.New("cannot create sender entry")
	}

	entries.Account_id = recieverUser.Id
	entries.Status = "credit"
	entries.Amount = amount

	err = bu.CreateEntries(entries)
	if err != nil {

		return errors.New("cannot create reviever entry")
	}

	tx.Commit()

	return nil
}

// GetHistory Fetch the trasactions made by user 
func (bu *PostgresBankingRepository) GetHistory(id string) (*[]domain.Entries, error) {

	var records []domain.Entries

	err := bu.Conn.Model(&domain.Entries{}).Where("account_id=?", id).Find(&records).Error
	if err != nil {

		log.Error(err)

		return nil, err
	}

	return &records, nil
}

// GetBalance fetch the current balance. 
func (bu *PostgresBankingRepository) GetBalance(id string) (*int, error) {

	user, err := bu.GetUser(id)

	if err != nil {

		log.Error(err)

		return nil, errors.New("Account doesnt exists")
	}

	return &user.Balance, nil
}

// GetUser take the user from the database 
func (bu PostgresBankingRepository) GetUser(accountId string) (*domain.Account, error) {

	var user domain.Account

	err := bu.Conn.Model(&domain.Account{}).Where("id = ?", accountId).Find(&user).Error
	if err != nil {

		return nil, err
	}

	return &user, nil
}
// UpdateAmount update the amount in the entries
func (bu PostgresBankingRepository) UpdateAmount(accountId string, amount int) error {

	tx := bu.Conn.Begin()
	defer tx.Rollback()

	err := tx.Model(&domain.Account{}).Where("id = ?", accountId).Update("balance", amount).Error

	if err != nil {

		return err
	}
	tx.Commit()

	return nil
}

// CreateEntries after trasaction its used to make the entries 
func (bu PostgresBankingRepository) CreateEntries(entries domain.Entries) error {

	tx := bu.Conn.Begin()

	err := tx.Model(&domain.Entries{}).Create(&entries).Error
	if err != nil {
		
		return err
	}

	tx.Commit()
	return nil
}
