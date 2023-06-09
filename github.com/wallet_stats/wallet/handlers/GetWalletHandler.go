package handlers

import (
	"net/http"

	"github.com/wallet_stats/wallet/controller"
	logger "github.com/wallet_stats/wallet/helpers"
)

// обработчик для получения конкретного кошелька
func GetWalletHandler(w http.ResponseWriter, r *http.Request) {
	responce, err := controller.GetWalletHandlerConroller(r)
	if err != nil {
		logger.ErrorLogWithError("GetWalletHandler", err)
		w.WriteHeader(500)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(responce)
}
