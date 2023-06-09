package controller

import (
	"encoding/json"
	"net/http"

	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/repository"
)

func GetWalletHandlerConroller(r *http.Request) ([]byte, error) {
	id := r.URL.Query().Get("id")

	responce, err := repository.GetWalletById(id)
	if err != nil {
		logger.ErrorLogWithError("GetWalletHandlerConroller", err)
		return nil, err
	}

	jsonResponce, err := json.Marshal(responce)
	if err != nil {
		logger.ErrorLogWithError("CreateWalletCotroller", err)
		return nil, err
	}

	return jsonResponce, nil
}
