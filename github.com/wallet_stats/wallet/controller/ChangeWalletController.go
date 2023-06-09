package controller

import (
	"encoding/json"
	"net/http"

	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/models"
	"github.com/wallet_stats/wallet/repository"
)

func ChangeWalletController(r *http.Request) ([]byte, error) {

	var request models.WalletNameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.ErrorLogWithError("ChangeWalletController", err)
		return nil, err
	}

	var responce models.WalletChangeResponce
	responce.Id = request.Id

	err = repository.ChangeWallet(request.Id, request.Name)
	if err != nil {
		logger.ErrorLogWithError("ChangeWalletController", err)
		responce.Err_code = err
		responce.Success = false
		jsonResponce, err := json.Marshal(responce)
		if err != nil {
			logger.ErrorLogWithError("ChangeWalletController", err)
			return nil, err
		}
		logger.ErrorLogWithError("ChangeWalletController", err)
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
