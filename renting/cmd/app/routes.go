package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()

	mux.Get("/rent", http.HandlerFunc(app.SetRent))

	mux.Get("/cars/view/:id", http.HandlerFunc(app.carsView))
	mux.Get("/users/view/:id", http.HandlerFunc(app.usersView))

	return LogRequest(SecureHeaders(mux))
}
