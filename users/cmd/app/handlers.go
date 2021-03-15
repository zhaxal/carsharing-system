package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (app *App) ShowUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	user, err := app.Database.GetUser(id)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if user == nil {
		app.NotFound(w)
		return
	}

	json.NewEncoder(w).Encode(user)
}
