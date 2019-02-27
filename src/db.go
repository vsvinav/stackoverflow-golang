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
	dbname         = "overflow2"
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

	fmt.Println("Successfully connected to postgres!")
}

// func deleteAllUsers() {
// 	db := getConnection()
// 	defer db.Close()
// 	sqlStatement := `DELETE FROM users`
// 	_, err := db.Exec(sqlStatement)
// 	if err != nil {
// 		panic(err)
// 	}
// }
