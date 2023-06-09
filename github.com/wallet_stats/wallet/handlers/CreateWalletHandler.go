package handlers

import (
	"net/http"

	"github.com/wallet_stats/wallet/controller"
	logger "github.com/wallet_stats/wallet/helpers"
)

// обработчик создания кошелька
func CreateWalletHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.ErrorLog("CreateWalletHandler", "Method not allowed")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}

	responce, err := controller.CreateWalletCotroller(r)
	if err != nil {
		logger.ErrorLogWithError("CreateWalletHandler", err)
		w.Write(responce)
		return
	}

	w.Write(responce)

}
