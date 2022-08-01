package logger

import (
	"go-starter/config"

	"go.uber.org/zap"
)

var Log *logger

type logger struct {
	*zap.Logger
}

func (l *logger) init() error {
	var err error
	if config.IsProdEnv() {
		l.Logger, err = zap.NewProduction()
	} else {
		l.Logger, err = zap.NewDevelopment()
	}
	l.Logger = l.Logger.With(
		zap.String("service", config.Cfg.ServiceName),
		zap.String("version", config.Cfg.ServiceVersion),
	)
	return err
}

func init() {
	var err error
	Log = &logger{}
	err = Log.init()
	if err != nil {
		panic(err)
	}
}
