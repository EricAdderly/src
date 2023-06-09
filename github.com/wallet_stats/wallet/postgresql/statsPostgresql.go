package postgresql

import (
	"errors"

	statsModel "github.com/wallet_stats/stats/model"
	logger "github.com/wallet_stats/wallet/helpers"
)

// Получение статистики из БД
func DbGetStats() (statsModel.Stats, error) {
	var stats statsModel.Stats
	if err := db.QueryRow(
		`SELECT total, active, inactive, deposited, withdrawn, transferred
		From public.stats
		Where stats.id = 1`).Scan(&stats.Total, &stats.Active, &stats.Inactive, &stats.Deposited, &stats.Withdrawn, &stats.Transferred); err != nil {
		logger.ErrorLog("DbGetStats", "problem with getting stats from db")
		return stats, errors.New("problem with getting stats from db")
	}

	return stats, nil
}

// Добавление в статистику общее количество кошельков в БД
func DbChangeTotal(newTotal int) error {
	_, err := db.Exec(
		`UPDATE public.stats
		SET total = total + $1
		Where stats.id = 1;`,
		newTotal)

	if err != nil {
		logger.ErrorLog("DbChangeTotal", "problem with change stats in db")
		return errors.New("problem with change stats in db")
	}

	return nil
}

// Добавление в статистику активных кошельков в БД
func DbChangeActive(newActive int) error {
	_, err := db.Exec(
		`UPDATE public.stats
		SET active = active + $1
		Where stats.id = 1;`,
		newActive)

	if err != nil {
		logger.ErrorLog("DbChangeActive", "problem with change stats in db")
		return errors.New("problem with change stats in db")
	}

	return nil
}

// Добавление в статистику количества неактивных кошельков в БД
func DbChangeInactive(newInActive int) error {
	_, err := db.Exec(
		`UPDATE public.stats
		SET Inactive = Inactive + $1
		Where stats.id = 1;`,
		newInActive)

	if err != nil {
		logger.ErrorLog("DbChangeInActive", "problem with change stats in db")
		return errors.New("problem with change stats in db")
	}

	return nil
}

// Добавление в статистику количества вложенных денег в БД
func DbChangeDeposited(newDeposited float64) error {
	_, err := db.Exec(
		`UPDATE public.stats
		SET Deposited = Deposited + $1
		Where stats.id = 1;`,
		newDeposited)

	if err != nil {
		logger.ErrorLog("DbChangeInActive", "problem with change stats in db")
		return errors.New("problem with change stats in db")
	}

	return nil
}

// Добавление в статистику количества снятых дененг в БД
func DbChangeWithdrawn(newWithdrawn float64) error {
	_, err := db.Exec(
		`UPDATE public.stats
		SET Withdrawn = Withdrawn + $1
		Where stats.id = 1;`,
		newWithdrawn)

	if err != nil {
		logger.ErrorLog("DbChangeWithdrawn", "problem with change stats in db")
		return errors.New("problem with change stats in db")
	}

	return nil
}

// Добавление в статистику количества переданных денег в БД
func DbChangeTransferred(newTransfered float64) error {
	_, err := db.Exec(
		`UPDATE public.stats
		SET transferred = transferred + $1
		Where stats.id = 1;`,
		newTransfered)

	if err != nil {
		logger.ErrorLog("DbChangeTransferred", "problem with change stats in db")
		return errors.New("problem with change stats in db")
	}

	return nil
}
