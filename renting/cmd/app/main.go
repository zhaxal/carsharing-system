package main

import (
	"carsharing-system/renting/pkg/models"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type apis struct {
	cars  string
	users string
}

type App struct {
	Database *models.Database
	apis     apis
}

func main() {
	addr := flag.String("addr", ":4002", "HTTP network address")
	carsAPI := flag.String("carsAPI", "http://localhost:4001/car/", "Cars API")
	usersAPI := flag.String("usersAPI", "http://localhost:4003/user/", "User API")
	dsn := flag.String("dsn", "root:aserty1234@/car_sharing?parseTime=true", "MySQL DSN")
	flag.Parse()

	db := connect(*dsn)
	defer db.Close()

	app := &App{
		Database: &models.Database{DB: db},
		apis: apis{
			cars:  *carsAPI,
			users: *usersAPI,
		},
	}

	log.Printf("Renting service listening on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)

}

func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
