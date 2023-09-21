package usecase

import (
	"context"
	models "e-project/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetWeatherUseCase struct {
	repo GetWetherRepo
}

func NewGetWeather(r GetWetherRepo) *GetWeatherUseCase {
	return &GetWeatherUseCase{
		repo: r,
	}
}

func (uc *GetWeatherUseCase) GetMeWeather(ctx context.Context, town string) (*models.Weather, error) {

	tempURL := newLink(town)
	resp, err := http.Get(tempURL)
	if err != nil {
		return nil, fmt.Errorf("GetWeatherUseCase - GetWeather - uc.repo.MakeRequest: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("GetWeatherUseCase - GetWeather - uc.repo.MakeRequest: %w", err)
	}

	var currentWeather *models.Weather
	err = json.Unmarshal(body, &currentWeather)
	if err != nil {
		return nil, fmt.Errorf("GetWeatherUseCase - GetWeather - uc.repo.MakeRequest: %w", err)
	}

	return currentWeather, nil
}

func newLink(town string) string {
	return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=697381875f993add6f43195450798d80&units=metric", town)
}
