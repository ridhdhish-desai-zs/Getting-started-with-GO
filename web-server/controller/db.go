package controller

import (
	"database/sql"
	"fmt"
)

func GetDbConnection() *sql.DB {
	// // Opening Database Connection
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")

	if err != nil {
		fmt.Println("Database Connection Error: ", err)
	}
	// defer db.Close()

	return db
}
