// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"e-project/internal/usecase"
	"e-project/internal/usecase/repo"
	"e-project/pkg/httpserver"
	"e-project/pkg/logger"
	"e-project/pkg/postgres"

	"e-project/config"
	v1 "e-project/internal/controller/http/v1"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax)) //connection to db
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// // Use case
	ReceivingTransactionUseCase := usecase.NewReceiveTransaction(repo.New(pg))
	// AuthUseCase := usecase.NewAuth(repo.NewAuthRepo(pg))

	// HTTP Server
	handler := gin.New()
	v1.NewReceivingTransaction(handler, l, ReceivingTransactionUseCase)
	// v2.NewRegistrationRouter(handler, l, RegistrationUseCase)             // initializes routes and handlers
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port)) // starting the server

	// Waiting signal
	interrupt := make(chan os.Signal, 1)                    // creating a channel
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM) // signals of interruption

	select {
	case s := <-interrupt: // waiting for an interrupt signal to be received
		l.Info("app - Run - signal: " + s.String()) // after receiving the signal, log
	case err = <-httpServer.Notify(): // if we got an error
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err)) // Logign
	}

	// Shutdown
	err = httpServer.Shutdown() // HTTP server stops
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServerAuth.Shutdown: %w", err)) // if err - logging
	}
}
