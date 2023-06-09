package handlers

import (
	"net/http"

	"github.com/wallet_stats/wallet/controller"
	logger "github.com/wallet_stats/wallet/helpers"
)

// обработчик для изменения клиентов
func ChangeWalletHandleFunc(w http.ResponseWriter, r *http.Request) {
	responce, err := controller.ChangeWalletController(r)
	if err != nil {
		logger.ErrorLogWithError("ChangeWalletHandleFunc", err)
		w.Write(responce)
		return
	}

	w.Write(responce)
}
