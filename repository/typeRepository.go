package repository

import (
	"database/sql"
	"mytokulist/structs"
	"time"
)

func GetAllType(db *sql.DB) (results []structs.Type, err error) {
	sql := "SELECT * FROM type"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var tipe = structs.Type{}

		err = rows.Scan(&tipe.ID, &tipe.Name, &tipe.Created_at, &tipe.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, tipe)
	}

	return
}

func InsertType(db *sql.DB, tipe structs.Type) (err error) {
	sql := "INSERT INTO type (name) VALUES ($1)"

	errs := db.QueryRow(sql, tipe.Name)

	return errs.Err()
}

func UpdateType(db *sql.DB, tipe structs.Type) (err error) {
	sql := "UPDATE type SET name = $1, updated_at = $2 WHERE id = $3"

	tipe.Updated_at = time.Now().Format(time.RFC3339)

	errs := db.QueryRow(sql, tipe.Name, tipe.Updated_at, tipe.ID)

	return errs.Err()
}

func DeleteType(db *sql.DB, tipe structs.Type) (err error) {
	sql := "DELETE FROM type WHERE id = $1"

	errs := db.QueryRow(sql, tipe.ID)

	return errs.Err()
}
