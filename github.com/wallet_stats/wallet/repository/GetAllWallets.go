package repository

import (
	"errors"

	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/models"
	"github.com/wallet_stats/wallet/postgresql"
)

// фу-я для получения всех кошельков
func GetAllWallets() ([]models.Wallet, error) {
	responce, err := postgresql.DbGetAllWallets()
	if err != nil {
		logger.ErrorLog("GetAllWallets", "problem with getting wallets from db")
		return nil, errors.New("problem with getting wallets from db")
	}
	return responce, nil
}
