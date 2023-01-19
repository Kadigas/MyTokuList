package repository

import (
	"database/sql"
	"mytokulist/models"
)

func GetAllWatchlist(db *sql.DB, users_id int64) (results []models.Watchlist, err error) {
	sql := "SELECT * FROM Watchlist WHERE users_id = $1"

	rows, err := db.Query(sql, users_id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var wl = models.Watchlist{}

		err = rows.Scan(&wl.ID, &wl.Users_id, &wl.Movie_id)
		if err != nil {
			panic(err)
		}

		results = append(results, wl)
	}

	return
}

func InsertWatchlist(db *sql.DB, wl models.Watchlist, users_id int64) (err error) {
	sql := "INSERT INTO Watchlist (users_id, movie_id) VALUES ($1, $2)"

	errs := db.QueryRow(sql, users_id, wl.Movie_id)

	return errs.Err()
}

func DeleteWatchlist(db *sql.DB, wl models.Watchlist, users_id int64) (err error) {
	sql := "DELETE FROM Watchlist WHERE id = $1 AND users_id = $2"

	errs := db.QueryRow(sql, wl.ID, users_id)

	return errs.Err()
}
