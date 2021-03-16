package models

type Car struct {
	ID        int
	Name      string
	Price     int
	Available int
	ExpReq    int
}

type Cars []*Car
