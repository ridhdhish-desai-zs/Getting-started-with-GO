package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"webserver/routes"
)

func main() {
	// Setting up mux router
	router := routes.NewRouter()

	// Handling Get request to fetch User(s)
	http.Handle("/", router)

	fmt.Println("Listening to Port 3000")
	// Start Server
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
