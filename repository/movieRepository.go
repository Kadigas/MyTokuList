package repository

import (
	"database/sql"
	"mytokulist/structs"
	"time"
)

func GetAllMovie(db *sql.DB) (results []structs.Movie, err error) {
	sql := "SELECT * FROM movie"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var movie = structs.Movie{}

		err = rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Image_url, &movie.Air_date, &movie.Total_episode, &movie.Current_status, &movie.Category_id, &movie.Type_id, &movie.Created_at, &movie.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, movie)
	}

	return
}

func InsertMovie(db *sql.DB, movie structs.Movie) (err error) {
	sql := "INSERT INTO Movie (title, description, image_url, air_date, total_episode, current_status, category_id, type_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	errs := db.QueryRow(sql, movie.Title, movie.Description, movie.Image_url, movie.Air_date, movie.Total_episode, movie.Current_status, movie.Category_id, movie.Type_id)

	return errs.Err()
}

func UpdateMovie(db *sql.DB, movie structs.Movie) (err error) {
	sql := "UPDATE Movie SET title = $1, updated_at = $2, description = $3, image_url = $4, air_date = $5, total_episode = $6, current_status = $7 WHERE id = $8"

	movie.Updated_at = time.Now().Format(time.RFC3339)

	errs := db.QueryRow(sql, movie.Title, movie.Updated_at, movie.Description, movie.Image_url, movie.Air_date, movie.Total_episode, movie.Current_status, movie.ID)

	return errs.Err()
}

func AttachMovie(db *sql.DB, movie structs.Movie) (err error) {
	sql := "UPDATE Movie SET category_id = $1, type_id = $2 WHERE id = $3"

	movie.Updated_at = time.Now().Format(time.RFC3339)

	errs := db.QueryRow(sql, movie.Category_id, movie.Type_id, movie.ID)

	return errs.Err()
}

func DeleteMovie(db *sql.DB, movie structs.Movie) (err error) {
	sql := "DELETE FROM Movie WHERE id = $1"

	errs := db.QueryRow(sql, movie.ID)

	return errs.Err()
}
