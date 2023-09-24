package main

import (
	gorilla "github.com/gorilla/mux"
	"net/http"
)

func (app *Application) routes() http.Handler {
	mux := gorilla.NewRouter()

	return mux
}
