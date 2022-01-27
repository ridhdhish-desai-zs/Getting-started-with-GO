package main

import (
	"fmt"
	"layer/user/driver"
	userHttp "layer/user/http/users"
	userServices "layer/user/services/users"
	userstore "layer/user/stores/users"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := driver.ConnectToMySql()

	if err != nil {
		// TODO: Handle error
		return
	}

	st := userstore.New(db)
	sr := userServices.New(st)
	handler := userHttp.Handler{S: sr}

	router := mux.NewRouter()
	router.Path("/api/users/{id}").Methods("GET").HandlerFunc(handler.GetUserByIdHandler)
	router.Path("/api/users").Methods("GET").HandlerFunc(handler.GetUsersHandler)
	router.Path("/api/users/{id}").Methods("PUT").HandlerFunc(handler.UpdateUserHandler)
	router.Path("/api/users/{id}").Methods("DELETE").HandlerFunc(handler.DeleteUserHandler)
	router.Path("/api/users").Methods("POST").HandlerFunc(handler.CreateUserHandler)

	http.Handle("/", router)

	fmt.Println("Listening to port 3000")
	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
