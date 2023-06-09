package repository

import (
	"errors"

	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/postgresql"
)

// var emptyWallet models.Wallet

// фу-ия для изменения имени у кошелька
func ChangeWallet(id string, newName string) error {
	if id == "" {
		logger.ErrorLog("ChangeWallet", "id is empty")
		return errors.New("id is empty")
	}

	if newName == "" {
		logger.ErrorLog("ChangeWallet", "name is empty")
		return errors.New("name is empty")
	}

	wallet, err := postgresql.DbGetWallet(id)
	if err != nil {
		logger.ErrorLog("ChangeWallet", "wallet is not found")
		return errors.New("wallet is not found")
	}

	if wallet.Status != "active" {
		logger.ErrorLog("ChangeWallet", "wallet status is not active")
		return errors.New("wallet status is not active")
	}

	err = postgresql.DbChangeWallet(id, newName)
	if err != nil {
		logger.ErrorLog("ChangeWallet", "problems with DB")
		return errors.New("problems with DB")
	}

	return nil
}
