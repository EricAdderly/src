package repository

import (
	"errors"

	statsRepository "github.com/wallet_stats/stats/repository"
	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/postgresql"
)

// фу-я для снятия дненг с кошелька
func WithdrawWallet(id string, amount float64) error {
	if id == "" {
		logger.ErrorLog("WithdrawWallet", "id cannot be empty")
		return errors.New("id cannot be empty")
	}

	if amount <= 0 {
		logger.ErrorLog("WithdrawWallet", "amount less than zero")
		return errors.New("amount less than zero")
	}

	wallet, err := postgresql.DbGetWallet(id)
	if err != nil {
		logger.ErrorLog("WithdrawWallet", "problems with DB")
		return errors.New("problems with DB")
	}

	if wallet.Balance < amount {
		logger.ErrorLog("WithdrawWallet", "not enough money")
		return errors.New("not enough money")
	}

	if wallet.Status == "inactive" {
		logger.ErrorLog("WithdrawWallet", "wallet is inactive")
		return errors.New("wallet is inactive")
	}

	err = postgresql.DbChangeWalletsAmount(id, -amount)
	if err != nil {
		logger.ErrorLog("WithdrawWallet", "problems with DB")
		return errors.New("problems with DB")
	}
	statsRepository.AddWithdrawalStats(amount)

	return nil
}
