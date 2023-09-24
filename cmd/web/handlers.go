package main

import (
	"net/http"
	"stavi/internal/logger"
)

func (app *Application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	app.logger.Println(logger.INFO, "hit the handler")
}
