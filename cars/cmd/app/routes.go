package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()

	mux.Get("/car/:id", http.HandlerFunc(app.ShowCar))

	return LogRequest(SecureHeaders(mux))
}