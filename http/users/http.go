package users

import (
	"encoding/json"
	"fmt"
	"layer/user/models"
	"layer/user/services"
	"net/http"
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
func (h Handler) UserById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)

	userId := params["id"]

	id, err := strconv.Atoi(userId)
	if err != nil {
		// TODO: Handle and return error response
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Invalid User id"}
		err, _ := json.Marshal(newError)
		_, _ = res.Write(err)
		return
	}

	user, err := h.S.GetUserById(id)
	if err != nil {
		// TODO: Return status code and error message
		fmt.Println(err)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. User id not found"}
		err, _ := json.Marshal(newError)
		_, _ = res.Write(err)
		return
	}

	data, _ := json.Marshal(user)
	_, _ = res.Write(data)

}
