package worker

import (
	"context"
	"database/sql"
	"github.com/AKuzyashin/checkbox/internal/logging"
	"github.com/AKuzyashin/checkbox/internal/models"
	srv "github.com/AKuzyashin/checkbox/internal/services"
	"github.com/AKuzyashin/checkbox/pkg/tomtom"
	"github.com/rs/zerolog"
	"sync"
	"time"
)

type Worker struct {
	ctx      context.Context
	wg       *sync.WaitGroup
	logger   zerolog.Logger
	services *srv.Services
	msgChan  chan models.WorkerMsg
}

func NewWorker(ctx context.Context, wg *sync.WaitGroup, services *srv.Services, msgChan chan models.WorkerMsg) *Worker {
	logger := logging.NewLogger()
	return &Worker{ctx: ctx, wg: wg, services: services, msgChan: msgChan, logger: logger}
}

func (w *Worker) Start() {
	go w.run()
	w.searchOld()
}

func (w *Worker) run() {
	w.wg.Add(1)
	defer w.wg.Done()
	for {
		select {
		case <-w.ctx.Done():
			w.logger.Info().Msg("Stopping worker")
			return
		default:

		}
		select {
		case <-w.ctx.Done():
			w.logger.Info().Msg("Stopping worker")
			return
		case msg := <-w.msgChan:
			result, err := w.services.TomTom.GetRoute(
				tomtom.RoutePoint{
					Latitude:  msg.FromLat,
					Longitude: msg.FromLng,
				}, tomtom.RoutePoint{
					Latitude:  msg.ToLat,
					Longitude: msg.ToLng,
				},
			)
			if err != nil {
				w.logger.Error().Err(err).Send()
			}
			err = w.services.Routes.SaveResult(&models.Route{
				ID:                  msg.RequestID,
				ProcessedAt:         sql.NullTime{Time: time.Now(), Valid: true},
				LengthInMeters:      result.Routes[0].Summary.LengthInMeters,
				TravelTimeInSeconds: result.Routes[0].Summary.TravelTimeInSeconds,
			})
			if err != nil {
				w.logger.Error().Err(err).Send()
			}
		}
	}
}

func (w *Worker) searchOld() {
	unprocessedRoutes := w.services.Routes.GetUnprocessed()
	for _, route := range unprocessedRoutes {
		w.msgChan <- route.ToWorkerMsg()
	}
}
