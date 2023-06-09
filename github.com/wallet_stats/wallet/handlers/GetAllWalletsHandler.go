package handlers

import (
	"net/http"

	"github.com/wallet_stats/wallet/controller"
	logger "github.com/wallet_stats/wallet/helpers"
)

// обработчик для получения всех кошельков
func GetAllWalletsHandler(w http.ResponseWriter, r *http.Request) {
	responce, err := controller.GetAllWalletsController()
	if err != nil {
		logger.ErrorLogWithError("GetAllWalletsHandler", err)
		w.WriteHeader(500)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(responce)
}
