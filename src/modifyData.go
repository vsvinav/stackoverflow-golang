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
