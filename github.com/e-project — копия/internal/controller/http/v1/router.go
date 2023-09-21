// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"github.com/gin-gonic/gin"

	"e-project/internal/usecase"
	"e-project/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRegistrationRouter(handler *gin.Engine, l logger.Interface, t usecase.Registration) {
	// Options
	handler.Use(gin.Logger())   // logging data about the request
	handler.Use(gin.Recovery()) //send 500 if there is a panic

	// K8s probe
	// handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) }) //defines the /healthy route, which returns the HTTP 200 OK status

	// Prometheus metrics
	//handler.GET("/metrics", gin.WrapH(promhttp.Handler())) // //defines the /metrics route

	// Routers
	//Question можно ли поменять структуру хэндлера для дальнейшей перекладки
	h := handler.Group("/v1") // add prefix to Group
	{
		newRegistrationRoutes(h, t, l)
	}
}

func NewAuthRouter(handler *gin.Engine, l logger.Interface, t usecase.Auth) {
	// Options
	handler.Use(gin.Logger())   // logging data about the request
	handler.Use(gin.Recovery()) //send 500 if there is a panic

	h := handler.Group("/v1") // add prefix to Group
	{
		newAuthRoutes(h, t, l)
	}
}
