package models

import (
	"database/sql"
)

type Database struct {
	*sql.DB
}

func (db *Database) GetRentList() (Rents, error) {
	stmt := `SELECT user_id, car_id FROM rent`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	rents := Rents{}

	for rows.Next() {
		r := &Rent{}

		err := rows.Scan(&r.IdUser, &r.IdCar)
		if err != nil {
			return nil, err
		}

		rents = append(rents, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rents, nil
}

func (db *Database) SetRent(userId, carId int) (int, error) {
	stmt := `INSERT INTO rent (user_id, car_id) VALUES (?,?)`
	stmt2 := `UPDATE car SET available=false WHERE id=?`

	result, err := db.Exec(stmt, userId, carId)
	db.Exec(stmt2, carId)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}
