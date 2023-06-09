package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func makeRequest(town string) Weather {
	tempURL := newLink(town)

	log.WithFields(log.Fields{
		"file":     "integrationToOpenweathermap",
		"function": "makeRequest",
		"error":    nil},
	).Info("The URL has been created")

	resp, err := http.Get(tempURL)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "integrationToOpenweathermap",
			"function": "makeRequest",
			"error":    err},
		).Error("The Program received an error from integration")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "integrationToOpenweathermap",
			"function": "makeRequest",
			"error":    err},
		).Error("The json response can't be read")
	}

	var currentWeather Weather
	err = json.Unmarshal(body, &currentWeather)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "integrationToOpenweathermap",
			"function": "makeRequest",
			"error":    err},
		).Error("The json response can't be unmarshal")
	}

	log.WithFields(log.Fields{
		"file":     "integrationToOpenweathermap",
		"function": "makeRequest",
		"error":    nil},
	).Info("The json response was received, read and saved")

	return currentWeather
}

func newLink(town string) string {
	return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=697381875f993add6f43195450798d80&units=metric", town)
}

func GiveTemp(town string) float32 {
	currentWeather := makeRequest(town)
	value := currentWeather.Main.Temp

	log.WithFields(log.Fields{
		"file":     "integrationToOpenweathermap",
		"function": "GiveTemp",
		"error":    nil},
	).Info("The temperature has been send to client")

	return value
}

func GiveTempFeelsLike(town string) float32 {
	currentWeather := makeRequest(town)
	value := currentWeather.Main.FeelsLike

	log.WithFields(log.Fields{
		"file":     "integrationToOpenweathermap",
		"function": "GiveTempFeelsLike",
		"error":    nil},
	).Info("The temperature feels like has been send to client")

	return value
}

func GiveWindSpeed(town string) float32 {
	currentWeather := makeRequest(town)
	value := currentWeather.Wind.WindSpeed

	log.WithFields(log.Fields{
		"file":     "integrationToOpenweathermap",
		"function": "GiveWindSpeed",
		"error":    nil},
	).Info("The Windspeed feels like has been send to client")

	return value
}
