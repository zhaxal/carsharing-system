package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (app *App) ShowRents(w http.ResponseWriter, r *http.Request) {
	rents, err := app.Database.GetRentList()
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if rents == nil {
		app.NotFound(w)
		return
	}

	json.NewEncoder(w).Encode(rents)
}

func (app *App) SetRent(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	carId, err := strconv.Atoi(r.URL.Query().Get("carId"))

	rents, err := app.Database.GetRentList()
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if rents == nil {
		app.NotFound(w)
		return
	}

	for _, r := range rents {
		if r.IdCar == carId && r.IdUser == userId {
			fmt.Fprint(w, "already taken")
			return
		}
	}
	app.Database.SetRent(userId, carId)
}
