package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

// var db *sql.DB = GetDbConnection()
// var err error

/*
	Url: /api/users/{id}
	Method: GET
	Route: Unprotected
	Description: Fetch Users or single user by Id
*/
func FindUserByIdRequestHandler(res http.ResponseWriter, req *http.Request) {
	var db *sql.DB = GetDbConnection()
	defer db.Close()

	// Fetching Record with given User ID
	res.Header().Add("content-type", "application/json")

	vars := mux.Vars(req)
	id := vars["id"]

	row := db.QueryRow("select * from users where id = ?", id)

	if row.Err() != nil {
		// TODO: Error Interface and return error if record for given id not found
		return
	}

	var u User
	err := row.Scan(&u.Id, &u.Name, &u.Age, &u.Address)

	// Returning Error message if not record found
	if err != nil {
		_, err = res.Write([]byte(`{"error": "No record found for given id"}`))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Returns null or user{}
	jsonData, _ := json.Marshal(u)
	_, err = res.Write([]byte(`{"data": {"user": ` + string(jsonData) + `}}`))
	if err != nil {
		fmt.Println(err)
	}

}

/*
	Url: /api/users
	Method: GET
	Route: Unprotected
	Description: Fetch Users or single user by Id
*/
func UserRequestHandler(res http.ResponseWriter, req *http.Request) {
	var db *sql.DB = GetDbConnection()
	defer db.Close()

	// Setting up response content-type
	res.Header().Add("content-type", "application/json")

	var users []User

	// Fetching all records
	result, err := db.Query("select * from users")

	if err != nil {
		fmt.Println(err)
		return
	}

	for result.Next() {
		var u User

		err := result.Scan(&u.Id, &u.Name, &u.Age, &u.Address)
		if err != nil {
			fmt.Println("Could not fetch error")
		}

		users = append(users, u)
	}

	// Returns null or users[{}]
	jsonData, _ := json.Marshal(users)

	if string(jsonData) == "null" {
		_, err = res.Write([]byte(`{"error": "No record found"}`))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	_, err = res.Write([]byte(`{"data": {"users": ` + string(jsonData) + `}}`))
	if err != nil {
		fmt.Println(err)
	}
}

/*
	Url: /api/users
	Method: POST
	Route: Unprotected
	Description: Create new user
*/
func CreateUserRequestHandler(res http.ResponseWriter, req *http.Request) {
	var db *sql.DB = GetDbConnection()
	defer db.Close()

	var u User

	err := json.NewDecoder(req.Body).Decode(&u)

	if err != nil {
		fmt.Println(err)
	}

	result, err := db.Exec("insert into users(name, age, address) values(?, ?, ?)", u.Name, u.Age, u.Address)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.LastInsertId())
}
