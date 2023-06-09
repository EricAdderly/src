package models

// структура для ответного json-а при изменении и удалении кошелька
type WalletChangeResponce struct {
	Success  bool   `json:"success"`
	Err_code error  `json:"err_code"`
	Id       string `json:"id"`
}

// структура для ответного json-а при перевода с одного кошелька на другой
type TransactionsResponce struct {
	Success  bool    `json:"success"`
	Err_code error   `json:"err_code"`
	Amount   float64 `json:"amount"`
}
