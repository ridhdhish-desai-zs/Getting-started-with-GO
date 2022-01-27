package stores

import (
	"layer/user/models"
)

type User interface {
	GetUserById(id int) (*models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(id int, user models.User) (int, error)
	DeleteUser(id int) (int, error)
	GetUserByEmail(email string) bool
	CreateUser(user models.User) (int, error)
}
