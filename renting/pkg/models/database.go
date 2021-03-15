package models

import (
	"database/sql"
)

type Database struct {
	*sql.DB
}

func (db *Database) SetRent(userId, carId int) (int, error) {
	stmt := `INSERT INTO rent (user_id, car_id) VALUES (?,?)`

	result, err := db.Exec(stmt, userId, carId)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}
