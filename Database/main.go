package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id        int64
	Name      string
	Age       int
	Address   string
	IsDeleted bool
}

var (
	id      int
	name    string
	age     string
	address string
)

var users []User

func main() {
	fmt.Println("Database Tutorial!!")

	// var users []User

	// Connect with database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	checkNilError(err)

	_, err = db.Exec("create table  if not exists users(id int primary key AUTO_INCREMENT, name varchar(255) not null, age int not null, address varchar(255) not null)")
	checkNilError(err)

	setUsers(db)

	var choice int

	for {
		fmt.Println()
		fmt.Println("1. Insert new user")
		fmt.Println("2. Delete existing user")
		fmt.Println("3. Update existing user")
		fmt.Println("4. Display all users")
		fmt.Println("Press any key to exit")

		fmt.Printf("--------------------------------------\nMake your choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			InsertRecord(db, "Ridhdhish", 21, "Surat")
			fmt.Println()
			fmt.Println("--------Inserted Successfully--------")
		case 2:
			deleted := DeleteRecord(db, 100)
			fmt.Println()

			if deleted {
				fmt.Println("---------Deleted Successfully--------")
			} else {
				fmt.Printf("Cannot delete record with id %d, because record with id %d does not exist.", id, id)
			}
			fmt.Println()
		case 3:
			UpdateRecord(db, 1)
			fmt.Println()
			fmt.Println("-------Age Updated Successfully------")
		case 4:
			fmt.Println()
			fmt.Println("--------------User Data--------------")
			fmt.Println(users)
			fmt.Println("------------End User Data------------")
			fmt.Println()
		default:
			return
		}
	}
}

// Create and insert a new user
func InsertRecord(db *sql.DB, name string, age int, address string) {
	convAge := strconv.Itoa(age)
	result, err := db.Exec("insert into users(name, age, address) values('" + name + "', '" + convAge + "', '" + address + "')")
	checkNilError(err)
	id, err := result.LastInsertId()
	checkNilError(err)

	u := User{Id: id, Name: name, Age: age, Address: address}
	users = append(users, u)
}

// Updating only Age
func UpdateRecord(db *sql.DB, id int) {
	queryString := "update users set age=? where id=?"

	var index int

	for k, v := range users {
		if v.Id == int64(id) {
			index = k
			break
		}
	}

	u := users[index]
	_, err := db.Exec(queryString, u.Age+1, id)
	checkNilError(err)

	u.Age += 1
	users[index] = u
}

// Delete a perticular user and update the users list
func DeleteRecord(db *sql.DB, id int) bool {
	queryString := "delete from users where id = ?"

	result, err := db.Exec(queryString, id)
	checkNilError(err)
	l, _ := result.RowsAffected()

	if l == 0 {
		return false
	}

	var index int

	// Fetching Index of the user who is deleted
	for k, v := range users {
		if v.Id == int64(id) {
			index = k
			break
		}
	}

	// Updating delete status then reassigning it to the original users list
	u := users[index]
	u.IsDeleted = true
	users[index] = u

	return true
}

// Fetch all records of users
func FetchAllRecords(db *sql.DB) {
	rows, err := db.Query("select * from users")

	checkNilError(err)

	for rows.Next() {
		err := rows.Scan(&id, &name, &age, &address)
		checkNilError(err)
		// fmt.Println(id, name)
	}
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

// Set the users
func setUsers(db *sql.DB) {
	rows, err := db.Query("select * from users")

	checkNilError(err)

	for rows.Next() {
		err := rows.Scan(&id, &name, &age, &address)
		checkNilError(err)

		convAge, _ := strconv.Atoi(age)
		u := User{Id: int64(id), Name: name, Age: convAge, Address: address}
		users = append(users, u)
	}
}
