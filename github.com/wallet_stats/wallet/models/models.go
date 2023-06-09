package models

// структура кошелька
type Wallet struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
	Status  string  `json:"status"`
} // TODO: выравнивание памяти

// структура кошелька завернутая в хэш-мапу
type DataWalletList struct {
	DataStore map[string]Wallet
}

// глоабльная переменная, где хранятся данные о кошельках
var WalletList DataWalletList
