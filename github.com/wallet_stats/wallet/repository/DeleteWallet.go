package repository

import (
	"errors"

	statsRepository "github.com/wallet_stats/stats/repository"
	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/postgresql"
)

// фу-я для удаления кошелька
func DeleteWallet(id string) error {
	if id == "" {
		logger.ErrorLog("DeleteWallet", "name cannot be empty")
		return errors.New("name cannot be empty")
	}

	wallet, err := postgresql.DbGetWallet(id)
	if err != nil {
		logger.ErrorLog("DeleteWallet", "wallet is not found")
		return errors.New("wallet is not found")
	}

	if wallet.Status == "inactive" {
		logger.ErrorLog("DeleteWallet", "wallet is inactive")
		return errors.New("wallet is inactive")
	}

	err = postgresql.DbDeleteWallet(id)
	if err != nil {
		logger.ErrorLog("DeleteWallet", "problems with DB")
		return errors.New("problems with DB")
	}

	statsRepository.AddInactiveStats()

	return nil
}
