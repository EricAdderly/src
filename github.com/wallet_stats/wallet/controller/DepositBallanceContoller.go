package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/models"
	"github.com/wallet_stats/wallet/repository"
)

func DepositBallanceConroller(r *http.Request) ([]byte, error) {
	var reqwest models.TransactionsRequest
	err := json.NewDecoder(r.Body).Decode(&reqwest)
	if err != nil {
		logger.ErrorLogWithError("DepositBallanceConroller", err)
		return nil, err
	}

	amount, err := strconv.ParseFloat(reqwest.Amount, 64)
	if err != nil {
		logger.ErrorLogWithError("DepositBallanceConroller", err)
		return nil, err
	}

	var responce models.TransactionsResponce
	responce.Amount = amount

	err = repository.DepositBallance(reqwest.Id, amount)
	if err != nil {
		if err != nil {
			responce.Err_code = err
			responce.Success = false
			jsonResponce, err := json.Marshal(responce)
			if err != nil {
				logger.ErrorLogWithError("DepositBallanceConroller", err)
				return nil, err
			}
			logger.ErrorLogWithError("DepositBallanceConroller", err)
			return jsonResponce, err
		}

	}

	responce.Err_code = nil
	responce.Success = true
	jsonResponce, err := json.Marshal(responce)
	if err != nil {
		logger.ErrorLogWithError("DepositBallanceConroller", err)
		return nil, err
	}

	return jsonResponce, nil
}
