package driver

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySql() (*sql.DB, error) {
	conn := "root:@tcp(localhost:3306)/test"

	db, err := sql.Open("mysql", conn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
