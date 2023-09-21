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

	// Use case
	RegistrationUseCase := usecase.New(repo.New(pg))
	AuthUseCase := usecase.NewAuth(repo.NewAuthRepo(pg))

	// HTTP Server
	handlerReg := gin.New()
	v1.NewRegistrationRouter(handlerReg, l, RegistrationUseCase)
	httpServerReg := httpserver.New(handlerReg, httpserver.Port(cfg.HTTP.Port))

	handlerAuth := gin.New()
	v1.NewAuthRouter(handlerAuth, l, AuthUseCase)
	httpServerAuth := httpserver.New(handlerAuth, httpserver.Port("3000"))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Run HTTP servers in goroutines
	go func() {
		err := httpServerReg.Run()
		if err != nil {
			l.Error(fmt.Errorf("app - Run - httpServerReg.Run: %w", err))
		}
	}()

	go func() {
		err := httpServerAuth.Run()
		if err != nil {
			l.Error(fmt.Errorf("app - Run - httpServerAuth.Run: %w", err))
		}
	}()

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServerReg.Notify():
		l.Error(fmt.Errorf("app - Run - httpServerReg.Notify: %w", err))
	}

	// Shutdown
	err = httpServerReg.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServerReg.Shutdown: %w", err))
	}

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServerAuth.Notify():
		l.Error(fmt.Errorf("app - Run - httpServerAuth.Notify: %w", err))
	}

	// Shutdown
	err = httpServerAuth.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServerAuth.Shutdown: %w", err))
	}
}
