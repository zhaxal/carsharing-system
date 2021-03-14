package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()

	mux.Get("/cars/view/:id", http.HandlerFunc(app.carsView))

	return LogRequest(SecureHeaders(mux))
}