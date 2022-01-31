package users

import (
	"errors"
	"layer/user/models"
	"layer/user/services"
	"layer/user/stores"
	"reflect"
)

type User struct {
	u stores.User
}

func New(u stores.User) services.User {
	return &User{u}
}

func (st *User) GetUserById(id int) (models.User, error) {
	user, err := st.u.GetUserById(id)

	if err != nil {
		return models.User{}, errors.New("Cannot fetch user for given id")
	}

	return *user, nil
}

func (st *User) GetUsers() ([]models.User, error) {
	users, err := st.u.GetUsers()

	if err != nil {
		return []models.User{}, errors.New("Cannot fetch users")
	}

	return users, nil
}

// TODO: Call getUserById() and return that user
func (st *User) UpdateUser(id int, user models.User) (int, error) {
	lastInsertedId, err := st.u.UpdateUser(id, user)

	if err != nil {
		return 0, errors.New("Could not able to update user for given id")
	}

	return lastInsertedId, nil
}

func (st *User) DeleteUser(id int) error {
	err := st.u.DeleteUser(id)

	if err != nil {
		return errors.New("Could not able to delete user for given id")
	}

	return nil
}

// TODO: Call getUserById and return the user object
func (st *User) CreateUser(user models.User) (int, error) {

	if reflect.DeepEqual(user, models.User{}) {
		return 0, errors.New("Need user data to create new user")
	}

	validEmail := validateEmail(user.Email)

	if !validEmail {
		return 0, errors.New("Invalid email address")
	}

	validPhone := validatePhone(user.Phone)

	if !validPhone {
		return 0, errors.New("Invalid phone number")
	}

	isValid := st.u.GetUserByEmail(user.Email)

	if !isValid {
		return 0, errors.New("Email id is already in use")
	}

	lastInsertedId, err := st.u.CreateUser(user)

	if err != nil {
		return 0, errors.New("Could not able to create new user")
	}

	return lastInsertedId, nil
}
