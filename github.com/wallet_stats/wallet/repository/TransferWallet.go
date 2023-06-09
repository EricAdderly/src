package repository

import (
	"errors"

	statsRepository "github.com/wallet_stats/stats/repository"
	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/postgresql"
)

// фу-я для перевода  дененг с одного кошелька на другой
func TransferWallet(tranferFrom string, amount float64, transferTo string) error {
	if tranferFrom == "" || transferTo == "" {
		logger.ErrorLog("TransferWallet", "id cannot be empty")
		return errors.New("id cannot be empty")
	}

	if amount <= 0 {
		logger.ErrorLog("TransferWallet", "amount less than zero")
		return errors.New("amount less than zero")
	}

	walletFrom, err := postgresql.DbGetWallet(tranferFrom)
	if err != nil {
		logger.ErrorLog("TransferWallet", "problems with DB")
		return errors.New("problems with DB")
	}

	if walletFrom.Balance < amount {
		logger.ErrorLog("TransferWallet", "not enough money")
		return errors.New("not enough money")
	}

	if walletFrom.Status == "inactive" {
		logger.ErrorLog("TransferWallet", "transfer's From wallet is inactive")
		return errors.New("transfer's From wallet is inactive")
	}

	walletTo, err := postgresql.DbGetWallet(transferTo)
	if err != nil {
		logger.ErrorLog("TransferWallet", "problems with DB")
		return errors.New("problems with DB")
	}

	if walletTo.Status == "inactive" {
		logger.ErrorLog("TransferWallet", "transfer's to wallet is inactive")
		return errors.New("transfer's From wallet is inactive")
	}

	err = postgresql.DbChangeWalletsAmount(tranferFrom, -amount)
	if err != nil {
		logger.ErrorLog("TransferWallet", "problems with DB")
		return errors.New("problems with DB")
	}

	err = postgresql.DbChangeWalletsAmount(tranferFrom, amount)
	if err != nil {
		logger.ErrorLog("TransferWallet", "problems with DB")
		return errors.New("problems with DB")
	}

	statsRepository.AddTransferStats(amount)

	return nil
}
