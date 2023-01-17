package repository

import (
	"database/sql"
	"mytokulist/structs"
	"time"
)

func GetAllStatus(db *sql.DB) (results []structs.Status, err error) {
	sql := "SELECT * FROM status"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var status = structs.Status{}

		err = rows.Scan(&status.ID, &status.Name, &status.Created_at, &status.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, status)
	}

	return
}

func InsertStatus(db *sql.DB, status structs.Status) (err error) {
	sql := "INSERT INTO status (name) VALUES ($1)"

	errs := db.QueryRow(sql, status.Name)

	return errs.Err()
}

func UpdateStatus(db *sql.DB, status structs.Status) (err error) {
	sql := "UPDATE Status SET name = $1, updated_at = $2 WHERE id = $3"

	status.Updated_at = time.Now().Format(time.RFC3339)

	errs := db.QueryRow(sql, status.Name, status.Updated_at, status.ID)

	return errs.Err()
}

func DeleteStatus(db *sql.DB, status structs.Status) (err error) {
	sql := "DELETE FROM Status WHERE id = $1"

	errs := db.QueryRow(sql, status.ID)

	return errs.Err()
}
