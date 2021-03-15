package models

type Car struct {
	ID        int
	Name      string
	Price     int
	Available bool
	ExpReq    int
}

type Cars []*Car
