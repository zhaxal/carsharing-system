package models

import (
	"database/sql"
)

type Database struct {
	*sql.DB
}

func (db *Database) GetCar(id int) (*Car, error) {
	stmt := `SELECT id, name, price, available, experience_required FROM car
	WHERE id = ?`

	row := db.QueryRow(stmt, id)

	c := &Car{}

	err := row.Scan(&c.ID, &c.Name, &c.Price, &c.Available, &c.ExpReq)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return c, nil
}
