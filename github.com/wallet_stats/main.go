package main

import (
	"net/http"

	_ "github.com/lib/pq"
	statsHandlers "github.com/wallet_stats/stats/handlers"
	"github.com/wallet_stats/wallet/handlers"
	logger "github.com/wallet_stats/wallet/helpers"
	"github.com/wallet_stats/wallet/postgresql"
)

type StorageConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	SSLMode  string
}

func main() {
	err := postgresql.ConnectionToDB()
	if err != nil {
		logger.ErrorLog("main", "problem with starting DB connection")
		return
	}

	// models.WalletList.DataStore = make(map[string]models.Wallet)
	http.HandleFunc("/wallet", handlers.CreateWalletHandler)
	http.HandleFunc("/wallets", handlers.CheckTypeWalletsHandler)
	http.HandleFunc("/wallets/deposit", handlers.DepositBallanceHandler)
	http.HandleFunc("/wallets/withdraw", handlers.WithdrawWalletHadler)
	http.HandleFunc("/wallets/transfer", handlers.TransferWalletHadler)
	http.HandleFunc("/stats/wallets", statsHandlers.GetStatsHandler)
	http.ListenAndServe(":8080", nil)

}

// type Wallet struct {
// 	Id   string
// 	next *Wallet
// 	previous *Wallet
// }

// func (w *Wallet) insertWallet(id string) {
// 	temp1 := &Wallet{Id, nil}

// 	if

// }
