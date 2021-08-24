package main

import (
	"context"
	"github.com/AKuzyashin/checkbox/internal/api"
	"github.com/AKuzyashin/checkbox/internal/config"
	"github.com/AKuzyashin/checkbox/internal/handlers"
	"github.com/AKuzyashin/checkbox/internal/logging"
	"github.com/AKuzyashin/checkbox/internal/models"
	"github.com/AKuzyashin/checkbox/internal/services"
	"github.com/AKuzyashin/checkbox/internal/storage"
	"github.com/AKuzyashin/checkbox/internal/worker"
	"os"
	"os/signal"
	"sync"
	"syscall"
)


// @title Swagger Example API
// @version 1.0
// @description Service for calculation distance between 2 GEO points

// @contact.name Alexey Kuzyashin
// @contact.email terr.kuzyashin@gmail.com
func main()  {
	logger := logging.NewLogger()
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	wg := sync.WaitGroup{}
	db, err := storage.NewPostgres(cfg)

	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	srv := services.NewServices(cfg, db)
	msgChan := make(chan models.WorkerMsg)
	ctx, cancel := context.WithCancel(context.Background())
	apiApp := api.NewApi(srv, cfg)
	hdl := handlers.NewHandlers(srv, msgChan)
	hdl.Register(apiApp.Gin())
	work := worker.NewWorker(ctx, &wg, srv, msgChan)
	go work.Start()
	go apiApp.Start()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit
	cancel()
	wg.Wait()
}