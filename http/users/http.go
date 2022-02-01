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
	hndlr services.User
}

func New(s services.User) Handler {
	return Handler{s}
}

func (srv Handler) GetUserByIdHandler(res http.ResponseWriter, req *http.Request) {
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

	if id <= 0 {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Id should be greater than 0"}
		err, _ := json.Marshal(newError)
		_, _ = res.Write(err)
		return
	}

	user, err := srv.hndlr.GetUserById(id)
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

func (srv Handler) GetUsersHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	users, err := srv.hndlr.GetUsers()
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Could not fetch users"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	responseData := models.Response{
		Data:       users,
		StatusCode: 200,
		Message:    "Successful operation",
	}

	jsonData, _ := json.Marshal(responseData)
	_, _ = res.Write([]byte(string(jsonData)))

}

func (srv Handler) UpdateUserHandler(res http.ResponseWriter, req *http.Request) {
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

	if convId <= 0 {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Id should be greater than 0"}
		err, _ := json.Marshal(newError)
		_, _ = res.Write(err)
		return
	}

	updatedUser, err := srv.hndlr.UpdateUser(convId, user)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: err.Error()}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	responseData := models.Response{
		Data:       updatedUser,
		StatusCode: 201,
		Message:    "User updated Successfully",
	}

	jsonData, _ := json.Marshal(responseData)

	_, _ = res.Write(jsonData)
}

func (srv Handler) DeleteUserHandler(res http.ResponseWriter, req *http.Request) {
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

	if convId <= 0 {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Id should be greater than 0"}
		err, _ := json.Marshal(newError)
		_, _ = res.Write(err)
		return
	}

	err = srv.hndlr.DeleteUser(convId)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Something went wrong"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	_, _ = res.Write([]byte(`{"data": "user deleted successfully"}`))

}

func (srv Handler) CreateUserHandler(res http.ResponseWriter, req *http.Request) {
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

	createdUser, err := srv.hndlr.CreateUser(user)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: err.Error()}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	responseData := models.Response{
		Data:       createdUser,
		Message:    "User created successfully",
		StatusCode: 201,
	}

	jsonData, _ := json.Marshal(responseData)

	_, _ = res.Write([]byte(string(jsonData)))

}
