package users

import (
	"encoding/json"
	"fmt"
	"layer/user/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	S services.User
}

func (h Handler) UserById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)

	userId := params["id"]

	id, err := strconv.Atoi(userId)
	if err != nil {
		// TODO: Handle and return error response
		_, _ = res.Write([]byte(`{"error": "Some error here"}`))
	}

	user, err := h.S.GetUserById(id)
	if err != nil {
		// TODO: Return status code and error message
		fmt.Println(err)
		return
	}

	data, _ := json.Marshal(user)
	_, _ = res.Write(data)

}
