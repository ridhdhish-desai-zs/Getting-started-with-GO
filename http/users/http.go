package users

import (
	"encoding/json"
	"fmt"
	"layer/user/models"
	"layer/user/services"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	S services.User
}

/*
URL: /api/users/{id}
Method: GET
Route: Unprotected
Description: Fetch user by it's id
*/
func (h Handler) GetUserByIdHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)

	userId := params["id"]

	id, err := strconv.Atoi(userId)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Invalid User id"}
		err, _ := json.Marshal(newError)
		_, _ = res.Write(err)
		return
	}

	user, err := h.S.GetUserById(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. User id not found"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)
		return
	}

	jsonData, _ := json.Marshal(user)
	_, _ = res.Write([]byte(fmt.Sprintf(`{"data": {"user": %v}}`, string(jsonData))))

}

/*
URL: /api/users
Method: GET
Route: Unprotected
Description: Fetch all users
*/
func (h Handler) GetUsersHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	users, err := h.S.GetUsers()
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Could not fetch users"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	jsonData, _ := json.Marshal(users)
	_, _ = res.Write([]byte(fmt.Sprintf(`{"data": {"users": %v}}`, string(jsonData))))

}

/*
URL: /api/users/{id}
Method: PUT
Route: Unprotected
Description: Update user for given id
*/
func (h Handler) UpdateUserHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil || reflect.DeepEqual(user, models.User{}) {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Cannot parse request data"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	params := mux.Vars(req)
	id := params["id"]

	convId, err := strconv.Atoi(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Invalid id"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	_, err = h.S.UpdateUser(convId, user)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Something went wrong"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	_, _ = res.Write([]byte(`{"data": "user updated successfully"}`))
}

/*
URL: /api/users/{id}
Method: DELETE
Route: Unprotected
Description: Delete user for given id
*/
func (h Handler) DeleteUserHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)
	id := params["id"]

	convId, err := strconv.Atoi(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Invalid id"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	_, err = h.S.DeleteUser(convId)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Something went wrong"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	_, _ = res.Write([]byte(`{"data": "user deleted successfully"}`))

}

/*
URL: /api/users
Method: POST
Route: Unprotected
Description: Create new user
*/
func (h Handler) CreateUserHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil || reflect.DeepEqual(user, models.User{}) {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Cannot parse request data"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	_, err = h.S.CreateUser(user)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Something went wrong"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	_, _ = res.Write([]byte(`{"data": "user created successfully"}`))

}
