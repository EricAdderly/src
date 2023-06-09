package handlers

import (
	"net/http"

	"github.com/wallet_stats/wallet/controller"
	logger "github.com/wallet_stats/wallet/helpers"
)

// обработчик для снятия денег
func WithdrawWalletHadler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		logger.ErrorLog("WithdrawWalletHadler", "Method not allowed")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}

	responce, err := controller.WithdrawWalletCortroller(r)
	if err != nil {
		logger.ErrorLogWithError("WithdrawWalletHadler", err)
		w.WriteHeader(500)
		w.Write(responce)
		return
	}

	w.Write(responce)

}
