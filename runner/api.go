package runner

import (
	"context"
	"sync"

	"go-starter/config"
	"go-starter/logger"
	"go-starter/router"
	"go-starter/service"

	"go.uber.org/zap"
)

// API is the interface for rest api runner
type API interface {
	Go(ctx context.Context, wg *sync.WaitGroup)
}

type api struct {
}

func (runner *api) Go(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	loggerInstance := logger.Log.With(
		zap.Int("port", config.Cfg.APIPort),
	)
	loggerInstance.Info("starting the api server")

	dependencies := service.Init()
	router.NewAPI().Start(dependencies)
}

// NewAPI returns an instance of the REST API runner
func NewAPI() API {
	return &api{}
}
