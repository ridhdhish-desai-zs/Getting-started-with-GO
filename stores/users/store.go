package users

import (
	"context"
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

/*
GET /api/users/{id}
Fetch user for given id
*/
func (u *dbStore) GetUserById(id int) (*models.User, error) {
	db := u.db

	row := db.QueryRow("SELECT * FROM user WHERE id = ?", id)

	var user models.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
	if err != nil {
		return nil, errors.New("Invalid Id")
	}

	return &user, nil
}

/*
GET /api/users
Fetch all users
*/
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

/*
PUT /api/users/{id}
Update user for given id
*/
func (u *dbStore) UpdateUser(id int, user models.User) (int, error) {
	db := u.db
	var err error

	ctx := context.Background()

	tx, _ := db.BeginTx(ctx, nil)

	if user.Name != "" {
		_, err = tx.ExecContext(ctx, "UPDATE user SET name = ? WHERE id = ?", user.Name, id)
		if err != nil {
			_ = tx.Rollback()
			return 0, errors.New("Internal server error")
		}
	}

	if user.Email != "" {
		_, err = tx.ExecContext(ctx, "UPDATE user SET email = ? WHERE id = ?", user.Email, id)
		if err != nil {
			_ = tx.Rollback()
			return 0, errors.New("Internal server error")
		}
	}

	if user.Age != 0 {
		_, err = tx.ExecContext(ctx, "UPDATE user SET age = ? WHERE id = ?", user.Age, id)
		if err != nil {
			_ = tx.Rollback()
			return 0, errors.New("Internal server error")
		}
	}

	if user.Phone != "" {
		_, err = tx.ExecContext(ctx, "UPDATE user SET phone = ? WHERE id = ?", user.Phone, id)
		if err != nil {
			_ = tx.Rollback()
			return 0, errors.New("Internal server error")
		}
	}

	// _, err = db.Exec("UPDATE user SET name = ?, email = ?, phone = ?, age = ? WHERE id = ?", user.Name, user.Email, user.Phone, user.Age, id)

	_ = tx.Commit()

	return id, nil
}

/*
DELETE /api/users/{id}
Delete user for given id
*/
func (u *dbStore) DeleteUser(id int) (int, error) {
	db := u.db

	result, err := db.Exec("DELETE FROM user WHERE id = ?", id)

	if err != nil {
		return 0, errors.New("Could not delete user for given id")
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return 0, errors.New("Could not delete user for given id")
	}

	return int(rowsAffected), nil
}

// Only used for email validation (email exists or not)
func (u *dbStore) GetUserByEmail(email string) bool {
	db := u.db

	row := db.QueryRow("SELECT * FROM user WHERE email = ?", email)

	var user models.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)

	return err != nil
}

/*
POST /api/users
Creating new user
*/
func (u *dbStore) CreateUser(user models.User) (int, error) {
	db := u.db

	result, err := db.Exec("INSERT INTO user(name, email, phone, age) VALUES(?, ?, ?, ?)", user.Name, user.Email, user.Phone, user.Age)

	if err != nil {
		return 0, errors.New("Could not create new user")
	}

	lastInsertedId, _ := result.LastInsertId()

	return int(lastInsertedId), nil
}
