package handlers

import (
	"net/http"

	"github.com/wallet_stats/wallet/controller"
	logger "github.com/wallet_stats/wallet/helpers"
)

// обработчик пополнения кошелька
func DepositBallanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		logger.ErrorLog("DepositBallanceHandler", "Method not allowed")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}

	responce, err := controller.DepositBallanceConroller(r)
	if err != nil {
		logger.ErrorLogWithError("DepositBallanceHandler", err)
		w.Write(responce)
		return
	}

	w.Write(responce)
}
