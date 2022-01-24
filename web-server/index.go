package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	Id      int
	Name    string
	Age     int
	Address string
}

var db *sql.DB
var err error

/*
	Url: /api/users/{id}
	Protected: false
	Description: Fetch Users or single user by Id
*/
func FindUserByIdRequestHandler(res http.ResponseWriter, req *http.Request) {
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
		res.Write([]byte(`{"error": "No record found for given id"}`))
		return
	}

	// Returns null or user{}
	jsonData, _ := json.Marshal(u)
	res.Write([]byte(`{"data": {"user": ` + string(jsonData) + `}}`))

}

/*
	Url: /api/users
	Protected: false
	Description: Fetch Users or single user by Id
*/
func UserRequestHandler(res http.ResponseWriter, req *http.Request) {

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
		res.Write([]byte(`{"error": "No record found"}`))
		return
	}

	res.Write([]byte(`{"data": {"users": ` + string(jsonData) + `}}`))
}

func main() {
	// // Opening Database Connection
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/test")

	if err != nil {
		fmt.Println("Database Connection Error: ", err)
	}
	defer db.Close()

	// Setting up mux (It's like express in node.js)
	r := mux.NewRouter()
	r.HandleFunc("/api/users", UserRequestHandler).Methods("POST", "GET")
	r.HandleFunc("/api/users/{id}", FindUserByIdRequestHandler).Methods("POST", "GET")

	// Handling Get request to fetch User(s)
	http.Handle("/", r)

	fmt.Println("Listening to Port 3000")
	// Start Server
	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
