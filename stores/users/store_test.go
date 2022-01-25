package users

import (
	"errors"
	"fmt"
	"layer/user/models"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).AddRow(
		1, "Naruto", "naruto@japan.com", "9999999999", 21,
	)

	tests := []struct {
		desc      string
		id        int
		expected  *models.User
		mockQuery *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", id: 1, expected: &models.User{Id: 1, Name: "Naruto", Email: "naruto@gmail.com", Phone: "9999999999", Age: 21}, mockQuery: mock.ExpectQuery("SELECT * FROM user WHERE id = ?").WithArgs(1).WillReturnRows(rows)},
		{desc: "Case2", id: 2, expected: nil, mockQuery: mock.ExpectQuery("SELECT * FROM user WHERE id = ?").WithArgs(2).WillReturnError(errors.New("Invalid Id"))},
	}

	userStore := New(db)

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			u, err := userStore.GetUserById(test.id)

			if err != nil && !reflect.DeepEqual(u, test.expected) {
				t.Errorf("Expected: %v, Got: %v", rows, u)
			}
		})
	}
}
