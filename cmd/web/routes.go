package main

import (
	gorilla "github.com/gorilla/mux"
	"net/http"
)

func (app *Application) routes() http.Handler {
	mux := gorilla.NewRouter()

	mux.HandleFunc("/virtual-terminal", app.VirtualTerminal).Methods("GET")

	return mux
}
