package v3

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"e-project/internal/usecase"
	"e-project/pkg/logger"
)

type getWeatherRouted struct {
	getWeather usecase.GetWeather
	logger     logger.Interface
	client     *redis.Client
}

func newRegistrationRoutes(handler *gin.RouterGroup, c *redis.Client, t usecase.GetWeather, l logger.Interface) {
	r := &getWeatherRouted{t, l, c}
	h := handler
	{
		h.GET("/weather", r.doGetWeather) //defines the route /translation/do-translate with the HTTP POST method,
	}
}

func (g *getWeatherRouted) doGetWeather(c *gin.Context) {

	pong, err := g.client.Ping().Result()
	if err != nil {
		panic(fmt.Errorf("Ошибка подключения к Redis: %v", err))
	}
	fmt.Println(pong) // Проверка: должно вывести "PONG"

	town := c.Query("town")
	if len(town) == 0 {
		g.logger.Error("http - v3 - doGetWeather")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	// var weather *models.Weather
	// weather, err := g.client.Get(town).Result()
	// if value != nil {

	// }

	// err = g.client.Set(town, "myValue", 100000).Err()
	// if err != nil {
	// 	panic(fmt.Errorf("Ошибка сохранения данных: %v", err))
	// }

	// fmt.Println(value)

	weather, err := g.getWeather.GetMeWeather(c, town)
	if err != nil {
		g.logger.Error("http - v3 - doGetWeather")
		errorResponse(c, http.StatusBadRequest, "problem with getting data")
	}

	fmt.Println(weather)

	c.JSON(http.StatusOK, weather)
}
