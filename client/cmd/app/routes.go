package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()

	mux.Get("/cars/view/all", isAuthorized(app.viewCars))
	mux.Get("/users/view/all", isAuthorized(app.viewUsers))

	mux.Get("/rents/rent", isAuthorized(app.rent))

	return LogRequest(SecureHeaders(mux))
}
