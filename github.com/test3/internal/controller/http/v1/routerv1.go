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

func NewReceivingTransaction(handler *gin.Engine, l logger.Interface, t usecase.ReceivingTransaction) {
	// Options
	handler.Use(gin.Logger())   // logging data about the request
	handler.Use(gin.Recovery()) //send 500 if there is a panic

	h := handler.Group("/v1") // add prefix to Group
	{
		newAuthRoutes(h, t, l)
	}
}
