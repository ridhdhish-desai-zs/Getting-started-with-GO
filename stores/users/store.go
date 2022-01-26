package users

import (
	"database/sql"
	"errors"

	"layer/user/models"
	"layer/user/stores"
)

type dbStore struct {
	db *sql.DB
}

func New(db *sql.DB) stores.User {
	return &dbStore{db: db}
}

// GET /api/users/{id}
func (u *dbStore) GetUserById(id int) (*models.User, error) {
	db := u.db

	row := db.QueryRow("SELECT * FROM user WHERE id = ?", id)

	if row.Err() != nil {
		return nil, errors.New("Invalid Id")
	}

	var user models.User
	_ = row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)

	return &user, nil
}

// GET /api/users
func (u *dbStore) GetUsers() ([]models.User, error) {
	db := u.db

	rows, err := db.Query("SELECT * FROM user")

	if err != nil {
		return []models.User{}, errors.New("Cannot fetch users")
	}

	var users []models.User

	for rows.Next() {
		var u models.User

		_ = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Phone, &u.Age)

		users = append(users, u)
	}

	return users, nil
}
