package repository

import (
	"database/sql"
	"mytokulist/structs"
	"time"
)

func GetAllCategory(db *sql.DB) (results []structs.Category, err error) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = structs.Category{}

		err = rows.Scan(&category.ID, &category.Name, &category.Created_at, &category.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, category)
	}

	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO category (name) VALUES ($1)"

	errs := db.QueryRow(sql, category.Name)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE category SET name = $1, updated_at = $2 WHERE id = $3"

	category.Updated_at = time.Now().Format(time.RFC3339)

	errs := db.QueryRow(sql, category.Name, category.Updated_at, category.ID)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"

	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
}
