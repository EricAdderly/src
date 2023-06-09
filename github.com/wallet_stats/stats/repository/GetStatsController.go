package statsRepository

import (
	"encoding/json"
	"net/http"

	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/postgresql"
)

// обрабатывает запрос и отдаёт статистику в формате для ответа
func GetStatsController(r *http.Request) ([]byte, error) {
	stats, err := postgresql.DbGetStats()
	if err != nil {
		logger.ErrorLog("GetStatsController", "problem with getting stats from db")
		return nil, err
	}

	jsonResponce, err := json.Marshal(stats)
	if err != nil {
		logger.ErrorLogWithError("GetStatsController", err)
		return []byte{}, err
	}

	return jsonResponce, nil
}
