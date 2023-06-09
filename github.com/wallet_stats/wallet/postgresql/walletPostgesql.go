package postgresql

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/models"
)

type StorageConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	SSLMode  string
}

var cfg = StorageConfig{
	Host:     "localhost",
	Port:     "5432",
	Username: "postgres",
	Password: "cdxjWRHO",
	Database: "Wallet",
	SSLMode:  "disable",
	// Host:     viper.GetString("db.host"),
	// Port:     viper.GetString("db.port"),
	// Username: viper.GetString("db.username"),
	// Password: viper.GetString("db.password"),
	// Database: viper.GetString("db.dbname"),
	// SSLMode:  viper.GetString("db.sslmode"),
}

type DataBaseWalletModel struct {
	ID      string  `json:"wallet_id"`
	Name    string  `json:"wallet_name"`
	Balance float64 `json:"wallet_balance"`
	Status  string  `json:"wallet_status"`
}

var db *sqlx.DB

// подключение к БД
func ConnectionToDB() error {
	var err error
	db, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Database, cfg.Password, cfg.SSLMode))
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Изменение имени кошелька в БД
func DbChangeWallet(id, newName string) error {
	result, err := db.Exec(
		`UPDATE public.wallets
		SET wallet_name = $1
		WHERE wallets.wallet_id = $2;`,
		newName,
		id)

	if err != nil {
		logger.ErrorLog("DbChangeWallet", "problem with change wallet in db")
		return errors.New("problem with change wallet in db")
	}

	modStrings, _ := result.RowsAffected()
	if modStrings == 0 {
		logger.ErrorLog("DbChangeWallet", "wallet is not founded in db")
		return errors.New("wallet is not founded in db")
	}

	return nil
}

// Создание кошелька в БД
func DbCreateWallet(id, name string) (models.Wallet, error) {
	var wallet models.Wallet
	_, err := db.Exec(
		`Insert into public.wallets
			(wallet_id, wallet_name, wallet_balance, wallet_status)
			VALUES ($1, $2, $3, $4) returning id`,
		id,
		name,
		0,
		"active")

	if err != nil {
		logger.ErrorLog("DbCreateWallet", "problem with change wallet in db")
		return wallet, errors.New("problem with creation a wallet in db")
	}

	wallet = models.Wallet{
		Id:      id,
		Name:    name,
		Balance: 0,
		Status:  "active",
	}

	return wallet, nil
}

// Удаление кошелька в БД (перевод в статус inactive)
func DbDeleteWallet(id string) error {
	_, err := db.Exec(
		`UPDATE public.wallets
		SET wallet_status = $1
		WHERE wallets.wallet_id = $2;`,
		"inactive",
		id)

	if err != nil {
		logger.ErrorLog("DbDeleteWallet", "problem with change wallet's status in db")
		return errors.New("problem with change wallet's status in db")
	}

	return nil
}

// Изменение баланска кошелька в БД
func DbChangeWalletsAmount(id string, amount float64) error {
	_, err := db.Exec(
		`UPDATE public.wallets
		SET wallet_balance = wallet_balance + $1
		Where wallets.wallet_id = $2`,
		amount,
		id)

	if err != nil {
		logger.ErrorLog("DbChangeWalletsAmount", "problem with deposit wallet in db")
		return errors.New("problem with deposit wallet in db")
	}

	return nil
}

// Получение кошелька из БД
func DbGetWallet(id string) (models.Wallet, error) {
	var dbWallet DataBaseWalletModel
	var wallet models.Wallet
	if err := db.QueryRow(
		`SELECT wallet_id, wallet_name, wallet_balance, wallet_status
		From public.wallets
		Where wallets.wallet_id = $1`,
		id).Scan(&dbWallet.ID, &dbWallet.Name, &dbWallet.Balance, &dbWallet.Status); err != nil {
		logger.ErrorLog("DbGetWallet", "problem with getting wallet from db")
		return wallet, errors.New("problem with getting wallet from db")
	}

	wallet = models.Wallet{
		Id:      dbWallet.ID,
		Name:    dbWallet.Name,
		Balance: dbWallet.Balance,
		Status:  dbWallet.Status,
	}

	return wallet, nil
}

// Получение всех кошельков
func DbGetAllWallets() ([]models.Wallet, error) {
	var wallets []models.Wallet
	row, err := db.Query(`
	SELECT wallet_id, wallet_name, wallet_balance, wallet_status
	From public.wallets`)
	if err != nil {
		logger.ErrorLog("DbGetAllWallet", "problem with getting wallets from db")
		return wallets, errors.New("problem with getting wallets from db")
	}

	defer row.Close()
	for row.Next() {
		item := models.Wallet{}
		err = row.Scan(&item.Id, &item.Name, &item.Balance, &item.Status)
		if err != nil {
			logger.ErrorLog("DbGetAllWallet", "problem with getting wallets from db")
			return wallets, errors.New("problem with getting wallets from db")
		}
		wallets = append(wallets, item)
	}

	return wallets, nil
}
