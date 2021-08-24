package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/AKuzyashin/checkbox/internal/config"
	"github.com/AKuzyashin/checkbox/internal/logging"
	"github.com/AKuzyashin/checkbox/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type api struct {
	logger zerolog.Logger
	services *services.Services
	app *gin.Engine
	httpServer *http.Server
}

func (a *api) Gin() *gin.Engine {
	return a.app
}

func (a *api) Start() {
	if err := a.httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		a.logger.Info().Err(err).Send()
	}

}

func (a *api) Shutdown() error {
	a.logger.Info().Msg("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return a.httpServer.Shutdown(ctx)
}

func NewApi(srv *services.Services,cfg *config.AppConfig) *api {
	app := gin.Default()
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler: app,
	}
	return &api{
		logger: logging.NewLogger(),
		services: srv,
		app:      app,
		httpServer: httpServer,
	}
}

