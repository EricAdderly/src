package statsModel

// Структура для сбора статистики
type Stats struct {
	Total       int     `json:"total"`
	Active      int     `json:"active"`
	Inactive    int     `json:"inactive"`
	Deposited   float64 `json:"deposited"`
	Withdrawn   float64 `json:"withdrawn"`
	Transferred float64 `json:"transferred"`
}

// Глобальная переменная для сбора статиситки
var Allstats Stats = Stats{
	Total:       0,
	Active:      0,
	Inactive:    0,
	Deposited:   0,
	Withdrawn:   0,
	Transferred: 0,
}
