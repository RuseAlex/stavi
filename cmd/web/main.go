package main

import (
	"stavi/internal/config"
	"stavi/internal/logger"
)

const VERSION = "1.0"
const CSSVERSION = "1"

type Application struct {
	cfg     config.Config
	logger  logger.Logger
	version string
}

func main() {
	app := Application{
		cfg: config.Config{},
	}
	cfgErr := app.cfg.LoadEnv()
	if cfgErr.Err != nil {
		app.logger.Println(cfgErr.Level, "error loading the configs")
	}
}
