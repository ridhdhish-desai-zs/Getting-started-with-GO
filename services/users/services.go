package users

import (
	"errors"
	"layer/user/models"
	"layer/user/services"
	"layer/user/stores"
)

type User struct {
	u stores.User
}

func New(u stores.User) services.User {
	return &User{u}
}

// From services we are calling store method (database interaction) which will be in store, and after getting the result we can
// implement middleware and send the response to http.

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

func (st *User) UpdateUser(id int, user models.User) (int, error) {
	lastInsertedId, err := st.u.UpdateUser(id, user)

	if err != nil {
		return 0, errors.New("Could not able to update user for given id")
	}

	return lastInsertedId, nil
}
