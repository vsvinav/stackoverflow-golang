package main

import (
	"database/sql"
	"fmt"
)

func upvote(postID int) Post {
	db := getConnection()
	defer db.Close()

	var post Post
	sqlStatement := `UPDATE posts SET score = score + 1 where id = $1`

	// row := db.QueryRow(sqlStatement, postID)
	_, err := db.Exec(sqlStatement, postID)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return Post{}
	case nil:
		// fmt.Println(post)
	default:
		panic(err)
	}
	return post
}

/*func upvote(postID int, userID int) Post {
	db := getConnection()
	defer db.Close()

	var post Post
	sqlStatement := `UPDATE posts SET score = score + 1 where id = $1`
	sqlStatement1 := `UPDATE users SET upvotes = upvotes + 1 where id = $1`

	// row := db.QueryRow(sqlStatement, postID)
	_, err := db.Exec(sqlStatement, postID)
	q1, err1 := db.Exec(sqlStatement1, userID)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return Post{}
	case nil:
		// fmt.Println(post)
	default:
		panic(err)
	}
	switch err1 {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return Post{}
	case nil:
		// fmt.Println(post)
	default:
		panic(err1)
	}
	return post
}
*/
func downvote(postID int) Post {
	db := getConnection()
	defer db.Close()

	var post Post
	sqlStatement := `UPDATE posts SET score = score - 1 where id = $1`
	// row := db.QueryRow(sqlStatement, postID)
	_, err := db.Exec(sqlStatement, postID)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return Post{}
	case nil:
		// fmt.Println(post)
	default:
		panic(err)
	}
	return post
}
