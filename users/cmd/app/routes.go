package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()

	mux.Get("/user/:id", http.HandlerFunc(app.ShowUser))

	return LogRequest(SecureHeaders(mux))
}
