package main

import (
	"net/http"

	"github.com/testFolder/app/internal"
)

func main() {
	internal.Tpl, _ = internal.Tpl.ParseGlob("app/templates/*.html")
	http.HandleFunc("/town", internal.AnswerHandleFunc)
	http.HandleFunc("/choosetown", internal.ChoseTownHandleFunc)
	http.HandleFunc("/weather", internal.GetWetherHandleFunc)
	http.ListenAndServe(":8080", nil)
}
