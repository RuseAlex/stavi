package main

import (
	"fmt"
	"net/http"
	"stavi/internal/config"
	"stavi/internal/logger"
	"stavi/internal/templates"
	"time"
)

const VERSION = "1.0"
const DEBUG = true

type Application struct {
	cfg     config.Config
	logger  *logger.Logger
	tc      templates.TemplateCache
	version string
}

func (app *Application) serveHTTP() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", app.cfg.Port),
		Handler:           app.routes(),
		IdleTimeout:       60 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	app.logger.Println(
		logger.INFO,
		fmt.Sprintf("starting HTTP server on port %s", app.cfg.Port),
	)

	return srv.ListenAndServe()
}

func (app *Application) serveHTTPS() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", app.cfg.Port),
		Handler:           app.routes(),
		IdleTimeout:       60 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	app.logger.Println(
		logger.INFO,
		fmt.Sprintf("starting HTTPS server on port %s", app.cfg.Port),
	)

	return srv.ListenAndServeTLS("", "")
}

func main() {
	app := Application{
		cfg:    config.Config{},
		logger: logger.New(DEBUG, "./logs/test.log"),
	}
	// Load the configs and check if there are any problems
	cfgErr := app.cfg.LoadEnv()
	if cfgErr.Err != nil {
		app.logger.Println(cfgErr.Level, cfgErr.Err.Error())
	} else {
		// this will only print in debugging mode
		app.logger.Println(cfgErr.Level, "env loaded successfully")
	}

	// start http server
	err := app.serveHTTP()
	if err != nil {
		app.logger.Println(logger.FATAL, err.Error())
	}
}
