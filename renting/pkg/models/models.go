package models

type Rent struct {
	IdUser int
	IdCar  int
}

type Rents []*Rent
