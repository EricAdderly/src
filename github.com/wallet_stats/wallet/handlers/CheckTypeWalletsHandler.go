package handlers

import (
	"net/http"

	logger "github.com/wallet_stats/wallet/helpers"
)

// обработчик для проверки каким методом обратились к /wallets
func CheckTypeWalletsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		if id == "" {
			GetAllWalletsHandler(w, r)
			return
		}
		GetWalletHandler(w, r)
		return
	}

	if r.Method == "PUT" {
		ChangeWalletHandleFunc(w, r)
		return
	}

	if r.Method == "DELETE" {
		DeleteWalletHandler(w, r)
		return
	}

	logger.ErrorLog("CheckTypeWalletsHandler", "Method not allowed")
	w.WriteHeader(405)
	w.Write([]byte("Method not allowed"))

}
