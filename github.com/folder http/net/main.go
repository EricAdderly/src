package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Weather struct {
	Main MainJson `json:"main"`
	Wind WindJson `json:"wind"`
}

type MainJson struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
}

type WindJson struct {
	WindSpeed float32 `json:"speed"`
}

func main() {
	town := "London"
	GiveTemp(town)
	GiveTempFeelsLike(town)
	GiveWindSpeed(town)
}

func makeRequest(town string) Weather {
	tempURL := newLink(town)
	resp, err := http.Get(tempURL)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var currentWeather Weather
	err = json.Unmarshal(body, &currentWeather)
	if err != nil {
		log.Fatalln(err)
	}
	return currentWeather
}

func newLink(town string) string {
	return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=697381875f993add6f43195450798d80&units=metric", town)
}

func GiveTemp(town string) {
	currentWeather := makeRequest(town)
	value := currentWeather.Main.Temp
	fmt.Println("Текущая температура:", value)
}

func GiveTempFeelsLike(town string) {
	currentWeather := makeRequest(town)
	value := currentWeather.Main.FeelsLike
	fmt.Println("Ощущается как:", value)
}

func GiveWindSpeed(town string) {
	currentWeather := makeRequest(town)
	value := currentWeather.Wind.WindSpeed
	fmt.Println("Скорость ветра:", value)
}
