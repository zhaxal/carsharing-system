package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type apis struct {
	cars  string
	rents string
	users string
}

type App struct {
	apis apis
}

var token string

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	carsAPI := flag.String("carsAPI", "http://localhost:4001/car/", "Cars API")
	usersAPI := flag.String("usersAPI", "http://localhost:4003/user/", "User API")
	rentsAPI := flag.String("rentsAPI", "http://localhost:4002/", "Rent API")
	flag.Parse()

	app := &App{
		apis: apis{
			cars:  *carsAPI,
			users: *usersAPI,
			rents: *rentsAPI,
		},
	}

	fmt.Println("Input your token")
	fmt.Scanln(&token)

	log.Printf("Server listening on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}
