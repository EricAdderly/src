package handlers

import (
	"net/http"

	"github.com/wallet_stats/wallet/controller"
	logger "github.com/wallet_stats/wallet/helpers"
)

// обработчик удаления кошелька
func DeleteWalletHandler(w http.ResponseWriter, r *http.Request) {
	responce, err := controller.DeleteWalletController(r)
	if err != nil {
		logger.ErrorLogWithError("DeleteWalletHandler", err)
		w.Write(responce)
		return
	}

	w.Write(responce)
}
