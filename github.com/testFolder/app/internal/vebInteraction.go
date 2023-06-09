package internal

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var Tpl *template.Template

var jsonCities = GetAllCities()

func AnswerHandleFunc(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "question.html", nil)
}

func ChoseTownHandleFunc(w http.ResponseWriter, r *http.Request) {
	letters := r.FormValue("town")
	possibleCities, err := FindCity(jsonCities, letters)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "vebInteraction",
			"function": "ChoseTownHandleFunc",
			"error":    err},
		).Error("The client entered an empty value")
	}
	Tpl.ExecuteTemplate(w, "chooseTown.html", possibleCities)
}

func GetWetherHandleFunc(w http.ResponseWriter, r *http.Request) {
	town := r.FormValue("town")
	temp := GiveTemp(town)
	tempFeelsLike := GiveTempFeelsLike(town)
	windSpeed := GiveWindSpeed(town)
	v := WeatherForClient{
		ClientTemp:          temp,
		ClientTempFeelsLike: tempFeelsLike,
		ClientWindSpeed:     windSpeed,
	}
	Tpl.ExecuteTemplate(w, "weather.html", v)
}
