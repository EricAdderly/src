package repository

import (
	"errors"

	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/models"
	"github.com/wallet_stats/wallet/postgresql"
)

// фу-я для получения кошелька по ID
func GetWalletById(id string) (models.Wallet, error) {
	var wallet models.Wallet
	if id == "" {
		logger.ErrorLog("GetWalletById", "Name cannot be empty")
		return wallet, errors.New("name cannot be empty")
	}

	wallet, err := postgresql.DbGetWallet(id)
	if err != nil {
		logger.ErrorLog("GetWalletById", "problems with DB")
		return wallet, errors.New("problems with DB")
	}

	return wallet, nil
}
