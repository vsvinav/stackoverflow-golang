package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host           = "localhost"
	port           = 5432
	dbUser         = "user12"
	dbUserPassword = "postgres"
	dbname         = "overflow"
)

func getConnection() *sql.DB {
	// Generate connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, dbUser, dbUserPassword, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db

}
func connectToDatabase() {
	// Generate connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, dbUser, dbUserPassword, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Check if you are able to communicate to database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

// func addUser(user User) {
// 	db := getConnection()
// 	defer db.Close()

// 	sqlStatement := `
// INSERT INTO users (age, email, first_name, last_name)
// VALUES ($1, $2, $3, $4)`
// 	_, err := db.Exec(sqlStatement, user.Age, user.Email, user.FirstName, user.LastName)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Added user to database", user)
// }

// func getUsers() []User {
// 	var users []User
// 	db := getConnection()
// 	defer db.Close()

// 	rows, err := db.Query("SELECT ID, age, email, first_name, last_name FROM users")
// 	if err != nil {
// 		// handle this error better than this
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		user := User{}
// 		err = rows.Scan(&user.ID, &user.Age, &user.Email, &user.FirstName, &user.LastName)
// 		if err != nil {
// 			// handle this error
// 			panic(err)
// 		}
// 		users = append(users, user)
// 	}
// 	// get any error encountered during iteration
// 	err = rows.Err()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Users in the database:", users)
// 	return users
// }

// func deleteAllUsers() {
// 	db := getConnection()
// 	defer db.Close()
// 	sqlStatement := `DELETE FROM users`
// 	_, err := db.Exec(sqlStatement)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func getUser(userID int32) User {
// 	db := getConnection()
// 	defer db.Close()

// 	var user User
// 	sqlStatement := `SELECT ID, age, email, first_name, last_name FROM users WHERE ID=$1`
// 	row := db.QueryRow(sqlStatement, userID)
// 	err := row.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
// 	switch err {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 		return User{}
// 	case nil:
// 		fmt.Println(user)
// 	default:
// 		panic(err)
// 	}
// 	return user
// }

func main() {

	getConnection()
	connectToDatabase()
	parseBadges()

}
