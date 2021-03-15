package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (app *App) ShowCar(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	car, err := app.Database.GetCar(id)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if car == nil {
		app.NotFound(w)
		return
	}

	json.NewEncoder(w).Encode(car)
}
