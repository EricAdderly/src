package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetAllCities() MainJsonCity {

	file, err := os.Open("app/town.json")
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "getCitiesFromJson",
			"function": "GetAllCities",
			"error":    err},
		).Error("File town.json hasn't been founded")
	} else if err == nil {
		log.WithFields(log.Fields{
			"file":     "getCitiesFromJson",
			"function": "GetAllCities",
			"error":    nil},
		).Info("File town.json has been founded")
	}

	body, err := ioutil.ReadAll(file)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "getCitiesFromJson",
			"function": "GetAllCities",
			"error":    err},
		).Error("File town.json can't be reed")
	} else if err == nil {
		log.WithFields(log.Fields{
			"file":     "getCitiesFromJson",
			"function": "GetAllCities",
			"error":    nil},
		).Info("File town.json has been read")
	}

	var cities MainJsonCity
	err = json.Unmarshal(body, &cities)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "getCitiesFromJson",
			"function": "GetAllCities",
			"error":    err},
		).Error("File town.json can't be unmarshal")
	} else if err == nil {
		log.WithFields(log.Fields{
			"file":     "getCitiesFromJson",
			"function": "GetAllCities",
			"error":    nil},
		).Info("File town.json has been unmarshaled")
	}

	return cities
}

func FindCity(jsonCities MainJsonCity, letters string) ([]string, error) {
	if letters == "" {
		var result []string
		result = append(result, letters)
		log.WithFields(log.Fields{
			"file":     "getCitiesFromJson",
			"function": "FindCity",
			"error":    "warn"},
		).Warn("The client entered an empty value")
		return result, errors.New("The client entered an empty value")
	}
	firstLetter := letters[0:1]
	value := jsonCities.City[firstLetter]
	var result []string
	for _, v := range value {
		if strings.HasPrefix(v, letters) == true {
			result = append(result, v)
		}
	}
	log.WithFields(log.Fields{
		"file":     "getCitiesFromJson",
		"function": "FindCity",
		"error":    nil},
	).Info("City has been founded")
	return result, nil

}
