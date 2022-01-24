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
	Method: GET
	Route: Unprotected
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

func main() {
	// // Opening Database Connection
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/test")

	if err != nil {
		fmt.Println("Database Connection Error: ", err)
	}
	defer db.Close()

	// Setting up mux (It's like express in node.js)
	r := mux.NewRouter()
	r.HandleFunc("/api/users", UserRequestHandler).Methods("GET")
	r.HandleFunc("/api/users/{id}", FindUserByIdRequestHandler).Methods("GET")
	r.HandleFunc("/api/users", CreateUserRequestHandler).Methods("POST")

	// Handling Get request to fetch User(s)
	http.Handle("/", r)

	fmt.Println("Listening to Port 3000")
	// Start Server
	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
