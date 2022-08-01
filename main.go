package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go-starter/config"
	"go-starter/logger"
	"go-starter/runner"

	"github.com/urfave/cli"
)

func shutDown(shutDownChannel chan *bool) {
	shutDown := true
	shutDownChannel <- &shutDown
	close(shutDownChannel)
}

// @title Go Starter API
// @version 0.0.1
// @description Go Starter API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func main() {

	var shutDownChannel chan *bool
	defer logger.Log.Sync()

	app := cli.NewApp()
	app.Name = config.Cfg.ServiceName
	app.Version = config.Cfg.ServiceVersion
	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start the api service",
			Action: func(c *cli.Context) error {
				logger.Log.Info("starting the api service")
				ctx := context.Background()

				var wg sync.WaitGroup
				wg.Add(1)

				go runner.NewAPI().Go(ctx, &wg)

				wg.Wait()
				return nil
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(
		signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		<-signalChannel
		shutDown(shutDownChannel)
	}()
}
