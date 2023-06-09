package statsRepository

import (
	log "github.com/sirupsen/logrus"
	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/postgresql"
)

// статистика общего количества кошельков
func AddStatsTotal() {
	err := postgresql.DbChangeTotal(1)
	if err != nil {
		logger.ErrorLog("AddStatsTotal", "problem with getting stats from db")
		return
	}

	log.WithFields(log.Fields{
		"file":  "stats",
		"func":  "AddStatsTotal",
		"error": nil,
	}).Info("Stats total has been added")
}

// статистика активных кошельков
func AddStatsActive() {
	err := postgresql.DbChangeActive(1)
	if err != nil {
		logger.ErrorLog("AddStatsActive", "problem with getting stats from db")
		return
	}

	log.WithFields(log.Fields{
		"file":  "stats",
		"func":  "AddStatsActive",
		"error": nil,
	}).Info("Stats active has been added")
}

// статистика неактивных кошельков
func AddInactiveStats() {
	err := postgresql.DbChangeActive(-1)
	if err != nil {
		logger.ErrorLog("AddInactiveStats", "problem with getting stats from db")
		return
	}

	err = postgresql.DbChangeInactive(1)
	if err != nil {
		logger.ErrorLog("AddInactiveStats", "problem with getting stats from db")
		return
	}

	log.WithFields(log.Fields{
		"file":  "stats",
		"func":  "AddInactiveStats",
		"error": nil,
	}).Info("Stats inactive has been added")
}

// статистика пополнения кошельков
func AddDepositStats(amount float64) {
	err := postgresql.DbChangeDeposited(amount)
	if err != nil {
		logger.ErrorLog("AddInactiveStats", "problem with getting stats from db")
		return
	}

	log.WithFields(log.Fields{
		"file":  "stats",
		"func":  "AddDepositStats",
		"error": nil,
	}).Info("Stats deposit has been added")
}

// статистика снятия дененг с кошелька
func AddWithdrawalStats(amount float64) {
	err := postgresql.DbChangeWithdrawn(amount)
	if err != nil {
		logger.ErrorLog("AddWithdrawalStats", "problem with getting stats from db")
		return
	}

	log.WithFields(log.Fields{
		"file":  "stats",
		"func":  "AddWithdrawalStats",
		"error": nil,
	}).Info("Stats withdrawal has been added")
}

// статистика переводов с кошелька на другой кошелёк
func AddTransferStats(amount float64) {
	err := postgresql.DbChangeTransferred(amount)
	if err != nil {
		logger.ErrorLog("AddTransferStats", "problem with getting stats from db")
		return
	}

	log.WithFields(log.Fields{
		"file":  "stats",
		"func":  "AddWithdrawalStats",
		"error": nil,
	}).Info("Stats Transfer has been added")
}
