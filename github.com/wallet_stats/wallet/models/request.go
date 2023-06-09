package models

// структура для обработки запросов при изменении и зодания кошелька
type WalletNameRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// структура для обработки запросов при пополнении/снятия дененг с кошелька
type TransactionsRequest struct {
	Id     string `json:"id"`
	Amount string `json:"amount"`
}

// структура для обработки запросов перевода с одного кошелька на другой
type TransferRequest struct {
	Id         string `json:"id"`
	Amount     string `json:"amount"`
	TransferTo string `json:"transfer_to"`
}
