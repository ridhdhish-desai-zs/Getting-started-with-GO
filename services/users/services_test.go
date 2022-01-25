package users

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"layer/user/models"
	"layer/user/stores"
)

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	tests := []struct {
		desc     string
		id       int
		expected models.User
		mockCall *gomock.Call
	}{
		{
			desc:     "Case1",
			id:       1,
			expected: models.User{Id: 1, Name: "Naruto", Email: "naruto@japan.com", Phone: "9999999999", Age: 18},
			mockCall: mockUserStore.EXPECT().GetUserById(1).Return(&models.User{Id: 1, Name: "Naruto", Email: "naruto@japan.com", Phone: "9999999999", Age: 18}, nil),
		},
		{
			desc:     "Case2",
			id:       2,
			expected: models.User{},
			mockCall: mockUserStore.EXPECT().GetUserById(2).Return(&models.User{}, errors.New("Cannot fetch user for given id")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			user, err := testUserService.GetUserById(test.id)

			if err != nil && !reflect.DeepEqual(test.expected, user) {
				t.Errorf("Expected: %v, Got: %v", test.expected, user)
			}
		})
	}
}
