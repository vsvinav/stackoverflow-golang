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

	fmt.Println("Successfully connected!")
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

//************ MAIN Function
func main() {

	getConnection()
	connectToDatabase()
	// // Milestone 2 Q1:
	// posts := getPosts()
	// fmt.Println(getPost(posts[0].ID))
	// Milestone 2 Q2:
	// comments := getComments()
	// fmt.Println(getPost(comments[6].PostID))
	// Milestone 2 Q3:
	// fmt.Println(getAnswer(posts[5].ID, "score"))
	// fmt.Println(getAnswer(posts[5].ID, "creation date"))
	// fmt.Println(getAnswer(posts[5].ID, "last activity date"))
	//Milestone 2 Q4:
	// users := getUsers()
	// fmt.Println(getUser(users[10].ID))
	// fmt.Println(upvote(posts[0].ID))
	// fmt.Println(getPost(posts[1].ID))
	// fmt.Println(downvote(posts[1].ID))
	addCustomer(takeCustomerInput())

}
