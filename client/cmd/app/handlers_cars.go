package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)



func (app *App) carsView(w http.ResponseWriter, r *http.Request)  {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	url := fmt.Sprintf("%s%b", app.apis.cars, id)

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

	fmt.Printf("%s\n", string(contents))

}