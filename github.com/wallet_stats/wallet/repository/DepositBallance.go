package repository

import (
	"errors"

	statsRepository "github.com/wallet_stats/stats/repository"
	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/postgresql"
)

// фу-я для пополнения кошелька
func DepositBallance(id string, amount float64) error {
	if id == "" {
		logger.ErrorLog("DepositBallance", "id is empty")
		return errors.New("id is empty")
	}

	wallet, err := postgresql.DbGetWallet(id)
	if err != nil {
		logger.ErrorLog("DepositBallance", "wallet is not found")
		return errors.New("wallet is not found")
	}

	if wallet.Status == "inactive" {
		logger.ErrorLog("DepositBallance", "wallet is inactive")
		return errors.New("wallet is inactive")
	}

	if amount <= 0 {
		logger.ErrorLog("DepositBallance", "amount less than zero")
		return errors.New("amount less than zero")
	}

	err = postgresql.DbChangeWalletsAmount(id, amount)
	if err != nil {
		logger.ErrorLog("DepositBallance", "problems with DB")
		return errors.New("problems with DB")
	}

	statsRepository.AddDepositStats(amount)

	return nil
}
