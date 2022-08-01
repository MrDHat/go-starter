package router

import (
	"fmt"
	"net/http"
	"strings"

	"go-starter/config"
	_ "go-starter/docs"
	"go-starter/logger"
	"go-starter/service"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type API interface {
	Start(dependencies service.Services)
}

type api struct {
}

func (r *api) attachMiddlewares(e *echo.Echo) {
	e.Debug = config.IsDevEnv()
	e.Use(echozap.ZapLogger(logger.Log.Logger))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	// for swagger
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
}

func (r *api) initRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]bool{
			"up": true,
		})
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (r *api) Start(dependencies service.Services) {
	e := echo.New()

	// Middleware
	r.attachMiddlewares(e)

	// Routes
	r.initRoutes(e)

	// Start server
	port := fmt.Sprintf(":%v", config.Cfg.APIPort)
	e.Logger.Fatal(e.Start(port))
}

func NewAPI() API {
	return &api{}
}
