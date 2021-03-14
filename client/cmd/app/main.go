package main

import (
	"flag"
	"log"
	"net/http"
)

type apis struct {
	cars string
}

type App struct {
	apis apis
}

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	carsAPI := flag.String("usersAPI", "http://localhost:4000/car/", "Cars API")
	flag.Parse()

	app := &App{
		apis: apis{
			cars: *carsAPI,
		},
	}

	log.Printf("Server listening on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}
