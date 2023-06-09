package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/models"
	"github.com/wallet_stats/wallet/repository"
)

func TransferWalletControler(r *http.Request) ([]byte, error) {
	var request models.TransferRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.ErrorLogWithError("TransferWalletControler", err)
		return nil, err
	}

	amount, err := strconv.ParseFloat(request.Amount, 64)
	if err != nil {
		logger.ErrorLogWithError("TransferWalletControler", err)
		return nil, err
	}

	var responce models.TransactionsResponce
	responce.Amount = amount

	err = repository.TransferWallet(request.Id, amount, request.TransferTo)
	if err != nil {
		responce.Err_code = err
		responce.Success = false
		jsonResponce, err4 := json.Marshal(responce)
		if err4 != nil {
			logger.ErrorLogWithError("TransferWalletControler", err)
			return nil, err4
		}
		logger.ErrorLogWithError("TransferWalletControler", err)
		return jsonResponce, err
	}

	responce.Err_code = nil
	responce.Success = true
	jsonResponce, err := json.Marshal(responce)
	if err != nil {
		logger.ErrorLogWithError("TransferWalletControler", err)
		return nil, err
	}

	return jsonResponce, nil

}
