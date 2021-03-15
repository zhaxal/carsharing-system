package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func (app *App) rent(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	carId := r.URL.Query().Get("carId")

	url := fmt.Sprintf("%s%s", app.apis.rents, fmt.Sprintf("rent?userId=%s&carId=%s", userId, carId))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Fprint(w, string(contents))
}
