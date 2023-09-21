package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/test2/handlers"
)

func main() {
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/refresh", handlers.Refresh)
	http.HandleFunc("/join", handlers.RegistrationHandler)
	expirationTime := time.Now().Add(time.Minute * 60)
	time := time.Now()
	fmt.Println(time, expirationTime)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// что нужно добавить:
// аут
// разбить логику
// добавить проверку на тип запроса
// добавить БД (уточнить как она будет выглядеть(поля))

// регистрация
// разбить логику
// прикрутить БД
// добавить проверку на тип запроса
