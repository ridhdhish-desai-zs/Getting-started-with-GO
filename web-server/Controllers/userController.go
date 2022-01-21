package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      int
	Name    string
	Age     int
	Address string
}

/*
	Url: /api/users
	Protected: false
	Description: Fetch Users or single user by Id
*/
func UserRequestHandler1(res http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")

	if err != nil {
		fmt.Println("Database Connection Error: ", err)
	}
	defer db.Close()

	if req.Method == http.MethodGet {
		params := req.URL.Query()

		// Setting up response content-type
		res.Header().Add("content-type", "application/json")

		if len(params) == 0 {
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

		} else {
			// Fetching Record with given User ID
			id := params.Get("id")
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
	}
}
