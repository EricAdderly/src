package repository

import (
	"errors"

	"github.com/google/uuid"
	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/postgresql"

	statsRepository "github.com/wallet_stats/stats/repository"
)

// фу-я для создания кошелька
func CreateWallet(name string) (string, error) {
	// проверка строки
	if name == "" {
		logger.ErrorLog("CreateWallet", "name cannot be empty")
		return "", errors.New("name cannot be empty")
	}
	id := uuid.New()
	stringId := id.String()

	_, err := postgresql.DbCreateWallet(stringId, name)
	if err != nil {
		logger.ErrorLog("CreateWallet", "problems with DB")
		return "", errors.New("problems with DB")
	}

	logger.InfoLog("CreateWallet", "wallet Created")
	statsRepository.AddStatsActive()
	statsRepository.AddStatsTotal()
	return stringId, nil
}
