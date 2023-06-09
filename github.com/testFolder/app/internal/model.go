package internal

// структуры для обращение к внешнему сервесу
type Weather struct {
	Main MainJson `json:"main"`
	Wind WindJson `json:"wind"`
}

type MainJson struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
}

// DataTransferObject
type WindJson struct {
	WindSpeed float32 `json:"speed"`
}

// структуры для поиска городов по json

type MainJsonCity struct {
	City CitiesMap `json:"city"`
}

type CitiesMap map[string][]string

// структура для отправки результатов клиенту

type WeatherForClient struct {
	ClientTemp          float32
	ClientTempFeelsLike float32
	ClientWindSpeed     float32
}
