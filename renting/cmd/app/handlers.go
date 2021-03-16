package main

import (
	models2 "carsharing-system/cars/pkg/models"
	"carsharing-system/users/pkg/models"
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

	urlUser := fmt.Sprintf("%s%b", app.apis.users, userId)
	urlCar := fmt.Sprintf("%s%b", app.apis.cars, carId)

	user := &models.User{}
	car := &models2.Car{}
	app.getAPIContent(urlUser, user)
	app.getAPIContent(urlCar, car)

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

	if user.Experience >= car.ExpReq {
		app.Database.SetRent(userId, carId)
		fmt.Fprint(w, "database updated")
		return

	} else {
		fmt.Fprint(w, "failed")
	}

}
