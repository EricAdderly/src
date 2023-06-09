package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/models"
	"github.com/wallet_stats/wallet/repository"
)

func WithdrawWalletCortroller(r *http.Request) ([]byte, error) {
	var request models.TransactionsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.ErrorLogWithError("WithdrawWalletCortroller", err)
		return nil, err
	}

	amount, err := strconv.ParseFloat(request.Amount, 64)
	if err != nil {
		logger.ErrorLogWithError("WithdrawWalletCortroller", err)
		return nil, err
	}

	var responce models.TransactionsResponce
	responce.Amount = amount

	err = repository.WithdrawWallet(request.Id, amount)
	if err != nil {
		responce.Err_code = err
		responce.Success = false
		jsonResponce, err := json.Marshal(responce)
		if err != nil {
			logger.ErrorLogWithError("WithdrawWalletCortroller", err)
			return nil, err
		}
		logger.ErrorLogWithError("WithdrawWalletCortroller", err)
		return jsonResponce, err
	}

	responce.Err_code = nil
	responce.Success = true
	jsonResponce, err := json.Marshal(responce)
	if err != nil {
		logger.ErrorLogWithError("WithdrawWalletCortroller", err)
		return nil, err
	}

	return jsonResponce, nil
}
