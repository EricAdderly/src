package statsHandlers

import (
	"net/http"

	statsRepository "github.com/wallet_stats/stats/repository"
	logger "github.com/wallet_stats/wallet/helpers"
)

// обработчик запросов по получения всей статистики
func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		logger.ErrorLog("GetStatsHandler", "Method not allowed")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}

	responce, err := statsRepository.GetStatsController(r)
	if err != nil {
		logger.ErrorLogWithError("GetStatsHandler", err)
		w.Write(responce)
	}

	w.Write(responce)

}
