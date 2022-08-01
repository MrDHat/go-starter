package runner

import (
	"context"
	"sync"

	"go-starter/config"
	"go-starter/logger"

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
	logger.Log.Info("starting the api server",
		zap.Int("port", config.Cfg.APIPort),
	)

	// TODO: start API server
}

// NewAPI returns an instance of the REST API runner
func NewAPI() API {
	return &api{}
}
