package controller

import (
	"encoding/json"
	"net/http"

	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/models"
	"github.com/wallet_stats/wallet/repository"
)

func CreateWalletCotroller(r *http.Request) ([]byte, error) {
	var request models.WalletNameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.ErrorLogWithError("CreateWalletCotroller", err)
		return nil, err
	}

	id, err := repository.CreateWallet(request.Name)
	if err != nil {
		logger.ErrorLogWithError("CreateWalletCotroller", err)
		return nil, err
	}

	responce, err := repository.GetWalletById(id)
	if err != nil {
		logger.ErrorLogWithError("CreateWalletCotroller", err)
		return nil, err
	}

	jsonResponce, err := json.Marshal(responce)
	if err != nil {
		logger.ErrorLogWithError("CreateWalletCotroller", err)
		return nil, err
	}

	return jsonResponce, nil
}
