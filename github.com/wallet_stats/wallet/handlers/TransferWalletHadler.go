package handlers

import (
	"net/http"

	"github.com/wallet_stats/wallet/controller"
	logger "github.com/wallet_stats/wallet/helpers"
)

// обработчик перевода денег с одного кошелька на другой
func TransferWalletHadler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		logger.ErrorLog("TransferWalletHadler", "Method not allowed")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}

	responce, err := controller.TransferWalletControler(r)
	if err != nil {
		logger.ErrorLogWithError("TransferWalletHadler", err)
		w.Write(responce)
		return
	}

	w.Write(responce)
}
