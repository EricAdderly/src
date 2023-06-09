package controller

import (
	"encoding/json"
	"net/http"

	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/models"
	"github.com/wallet_stats/wallet/repository"
)

func DeleteWalletController(r *http.Request) ([]byte, error) {
	id := r.URL.Query().Get("id")

	var responce models.WalletChangeResponce
	responce.Id = id

	err := repository.DeleteWallet(id)
	if err != nil {
		logger.ErrorLogWithError("DeleteWalletController", err)
		responce.Err_code = err
		responce.Success = false
		jsonResponce, err := json.Marshal(responce)
		if err != nil {
			logger.ErrorLogWithError("DeleteWalletController", err)
			return nil, err
		}
		logger.ErrorLogWithError("DeleteWalletController", err)
		return jsonResponce, err
	}

	responce.Err_code = nil
	responce.Success = true
	jsonResponce, err := json.Marshal(responce)

	if err != nil {
		logger.ErrorLogWithError("ChangeWalletController", err)
		return nil, err
	}

	return jsonResponce, nil
}
