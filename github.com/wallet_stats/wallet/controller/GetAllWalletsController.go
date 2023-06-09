package controller

import (
	"encoding/json"

	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/repository"
)

func GetAllWalletsController() ([]byte, error) {

	responce, err := repository.GetAllWallets()
	if err != nil {
		logger.ErrorLogWithError("GetAllWalletsController", err)
		return nil, err
	}

	jsonResponce, err := json.Marshal(responce)
	if err != nil {
		logger.ErrorLogWithError("GetAllWalletsController", err)
		return nil, err
	}
	return jsonResponce, nil
}
