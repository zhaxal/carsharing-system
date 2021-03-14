package models

import "time"

type User struct {
	ID         int
	Name       string
	Surname    string
	Birthdate  time.Time
	Experience int
}
