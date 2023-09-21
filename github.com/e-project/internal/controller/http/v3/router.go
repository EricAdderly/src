package v3

import (
	"e-project/internal/usecase"
	"e-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewGetWeatherRouter(handler *gin.Engine, l logger.Interface, t usecase.GetWeather) {
	// Options
	handler.Use(gin.Logger())   // logging data about the request
	handler.Use(gin.Recovery()) //send 500 if there is a panic
	c := dial()
	// Routers
	h := handler.Group("/v1") // add prefix to Group
	{
		newRegistrationRoutes(h, c, t, l)
	}
}
