package repository

import (
	"database/sql"
	"mytokulist/models"
)

func GetAdmins(db *sql.DB) (results []models.Users, err error) {
	sql := "SELECT * FROM users WHERE role = 'Admin'"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var admin = models.Users{}

		err = rows.Scan(&admin.ID, &admin.Username, &admin.Email, &admin.Password, &admin.Role, &admin.Created_at, &admin.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, admin)
	}

	return
}

func GetUsers(db *sql.DB) (results []models.Users, err error) {
	sql := "SELECT * FROM users WHERE role = 'User'"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = models.Users{}

		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.Created_at, &user.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, user)
	}

	return
}
