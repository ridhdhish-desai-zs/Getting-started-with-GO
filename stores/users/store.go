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
