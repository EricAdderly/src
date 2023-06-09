// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// type Weather struct {
// 	Main MainJson `json:"main"`
// 	Wind WindJson `json:"wind"`
// }

// type MainJson struct {
// 	Temp      float32 `json:"temp"`
// 	FeelsLike float32 `json:"feels_like"`
// }

// //DataTransferObject
// type WindJsonModel  struct {
// 	WindSpeed float32 `json:"speed"`
// }

// func main() {
// 	town := "London"
// 	GiveTemp(town)
// 	GiveTempFeelsLike(town)
// 	GiveWindSpeed(town)
// }

// func makeRequest(town string) Weather {
// 	tempURL := newLink(town)
// 	resp, err := http.Get(tempURL)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var currentWeather Weather
// 	err = json.Unmarshal(body, &currentWeather)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	return currentWeather
// }

// func newLink(town string) string {
// 	return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=697381875f993add6f43195450798d80&units=metric", town)
// }

// func GiveTemp(town string) {
// 	currentWeather := makeRequest(town)
// 	value := currentWeather.Main.Temp
// 	fmt.Println("Текущая температура:", value)
// }

// func GiveTempFeelsLike(town string) {
// 	currentWeather := makeRequest(town)
// 	value := currentWeather.Main.FeelsLike
// 	fmt.Println("Ощущается как:", value)
// }

// func GiveWindSpeed(town string) {
// 	currentWeather := makeRequest(town)
// 	value := currentWeather.Wind.WindSpeed
// 	fmt.Println("Скорость ветра:", value)
// }

package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// type MainJsonCity struct {
// 	City CitiesMap `json:"city"`
// }

// type CitiesMap map[string][]string

var tpl *template.Template

func main() {
	// jsonCities := GetAllCities()
	// possibleCities := FindCity(jsonCities, "Vi")
	// fmt.Println(possibleCities)
	tpl, _ = tpl.ParseGlob("github.com/folderHttp/net/templates/*.html")
	http.HandleFunc("/town", AnswerHandleFunc)
	http.HandleFunc("/chosetown", choseTownHandleFunc)
	http.ListenAndServe(":8080", nil)

}

func AnswerHandleFunc(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "qwestion.html", nil)
	// tpl, _ := template.ParseFiles("github.com/folder/http/net/templates/")
	// tpl.ExecuteTemplate(w, "qwestion.html", nil)
}

func choseTownHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Choose town")
	town := r.FormValue("Town")
	fmt.Println(town)

}

// func GetAllCities() MainJsonCity {

// 	file, err := os.Open("github.com/folder http/net/town.json")
// 	if err != nil {
// 		panic(err)
// 	}

// 	body, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var cities MainJsonCity
// 	err = json.Unmarshal(body, &cities)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return cities
// }

// func FindCity(jsonCities MainJsonCity, letters string) []string {
// 	firstLetter := letters[0:1]
// 	value := jsonCities.City[firstLetter]
// 	var result []string
// 	for _, v := range value {
// 		if strings.HasPrefix(v, letters) == true {
// 			result = append(result, v)
// 		}
// 	}

// 	return result

// }
