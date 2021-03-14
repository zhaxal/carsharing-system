package main

import (
	"carsharing-system/cars/pkg/models"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type App struct {
	Database  *models.Database
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "root:aserty1234@/car_sharing?parseTime=true", "MySQL DSN")
	flag.Parse()

	db := connect(*dsn)
	defer db.Close()

	app := &App{
		Database:  &models.Database{DB: db},
	}

	log.Printf("Server listening on %s", *addr)
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
