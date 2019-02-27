package main

import (
	"database/sql"
	"fmt"
)

func GetPosts() []Post {
	var posts []Post
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, view_count, answer_count, comment_count, view_count, favourite_count, closed_date, title FROM posts")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.ID, &post.ViewCount, &post.AnswerCount, &post.CommentCount, &post.ViewCount, &post.FavoriteCount, &post.ClosedDate, &post.Title)
		if err != nil {
			// handle this error
			panic(err)
		}
		posts = append(posts, post)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	// fmt.Println("Posts in the database:", posts)
	return posts
}

func GetPost(postID int) Post {
	db := getConnection()
	defer db.Close()

	var post Post
	sqlStatement := `SELECT  id,score, view_count, answer_count, comment_count, view_count, favourite_count, closed_date, title FROM posts where id = $1`
	row := db.QueryRow(sqlStatement, postID)
	err := row.Scan(&post.ID, &post.Score, &post.ViewCount, &post.AnswerCount, &post.CommentCount, &post.ViewCount, &post.FavoriteCount, &post.ClosedDate, &post.Title)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return Post{}
	case nil:
		// fmt.Println(post)
	default:
		panic(err)
	}
	// fmt.Printf("Score:%d", post.Score)
	return post
}

func GetPostForTest(postID int) Post {
	db := getConnection()
	defer db.Close()

	var post Post
	sqlStatement := `SELECT  id, view_count, answer_count, comment_count, view_count, favourite_count, closed_date, title FROM posts where id = $1`
	row := db.QueryRow(sqlStatement, postID)
	err := row.Scan(&post.ID, &post.ViewCount, &post.AnswerCount, &post.CommentCount, &post.ViewCount, &post.FavoriteCount, &post.ClosedDate, &post.Title)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return Post{}
	case nil:
		// fmt.Println(post)
	default:
		panic(err)
	}
	// fmt.Printf("Score:%d", post.Score)
	return post
}

// ************* Get Comments for the post

func GetComments() []Comment {
	var comments []Comment
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM comments")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		comment := Comment{}
		err = rows.Scan(&comment.ID, &comment.PostID, &comment.Score, &comment.Text, &comment.CreationDate, &comment.UserID)
		if err != nil {
			// handle this error
			panic(err)
		}
		comments = append(comments, comment)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return comments
}

func GetComment(postID int) Comment {
	db := getConnection()
	defer db.Close()

	var comment Comment
	sqlStatement := `SELECT * FROM comments where post_id = $1`
	row := db.QueryRow(sqlStatement, postID)
	err := row.Scan(&comment.ID, &comment.PostID, &comment.Score, &comment.Text, &comment.CreationDate, &comment.UserID)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return Comment{}
	case nil:
		fmt.Println(comment)
	default:
		panic(err)
	}
	return comment
}

// ************* Get Answers for the post, sorting according to user's choice

func GetAnswers() []Post {
	var posts []Post
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, view_count, answer_count, comment_count, view_count, favourite_count, closed_date, title, body FROM posts")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.ID, &post.ViewCount, &post.AnswerCount, &post.CommentCount, &post.ViewCount, &post.FavoriteCount, &post.ClosedDate, &post.Title, &post.Body)
		if err != nil {
			// handle this error
			panic(err)
		}
		posts = append(posts, post)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	// fmt.Println("Posts in the database:", posts)
	return posts
}

func GetAnswer(postID int, choice string) Post {
	db := getConnection()
	defer db.Close()

	var post Post
	if choice == "score" {
		sqlStatement := `SELECT  id, view_count, answer_count, comment_count, view_count, favourite_count, closed_date, title FROM posts where id = $1 order by score`
		row := db.QueryRow(sqlStatement, postID)
		err := row.Scan(&post.ID, &post.ViewCount, &post.AnswerCount, &post.CommentCount, &post.ViewCount, &post.FavoriteCount, &post.ClosedDate, &post.Title)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			return Post{}
		case nil:
			fmt.Println(post)
		default:
			panic(err)
		}
		return post
	} else if choice == "creation date" {
		sqlStatement := `SELECT  id, view_count, answer_count, comment_count, view_count, favourite_count, closed_date, title FROM posts where id = $1 order by creation_date`
		row := db.QueryRow(sqlStatement, postID)
		err := row.Scan(&post.ID, &post.ViewCount, &post.AnswerCount, &post.CommentCount, &post.ViewCount, &post.FavoriteCount, &post.ClosedDate, &post.Title)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			return Post{}
		case nil:
			fmt.Println(post)
		default:
			panic(err)
		}
		return post
	} else if choice == "last activity date" {
		sqlStatement := `SELECT  id, view_count, answer_count, comment_count, view_count, favourite_count, closed_date, title FROM posts where id = $1 order by last_activity_date`
		row := db.QueryRow(sqlStatement, postID)
		err := row.Scan(&post.ID, &post.ViewCount, &post.AnswerCount, &post.CommentCount, &post.ViewCount, &post.FavoriteCount, &post.ClosedDate, &post.Title)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			return Post{}
		case nil:
			fmt.Println(post)
		default:
			panic(err)
		}
		return post
	} else {
		sqlStatement := `SELECT  id, view_count, answer_count, comment_count, view_count, favourite_count, closed_date, title FROM posts where id = $1`
		row := db.QueryRow(sqlStatement, postID)
		err := row.Scan(&post.ID, &post.ViewCount, &post.AnswerCount, &post.CommentCount, &post.ViewCount, &post.FavoriteCount, &post.ClosedDate, &post.Title)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			return Post{}
		case nil:
			fmt.Println(post)
		default:
			panic(err)
		}
		return post
	}
}

// ************* Get Users
func GetUsers() []User {
	var users []User
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT display_name, id, reputation, location  FROM users")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.DisplayName, &user.ID, &user.Reputation, &user.Location)
		if err != nil {
			// handle this error
			panic(err)
		}
		users = append(users, user)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return users
}

func GetUser(userID int) User {
	db := getConnection()
	defer db.Close()

	var user User
	sqlStatement := `SELECT display_name, id, reputation, location  FROM users where id = $1`
	row := db.QueryRow(sqlStatement, userID)
	err := row.Scan(&user.DisplayName, &user.ID, &user.Reputation, &user.Location)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return User{}
	case nil:
		fmt.Println(user)
	default:
		panic(err)
	}
	return user
}

// ********** Get votes
func GetVotes(postID int) int {
	db := getConnection()
	defer db.Close()

	var post Post
	sqlStatement := `SELECT score from posts where id = $1`
	row := db.QueryRow(sqlStatement, postID)
	err := row.Scan(&post.Score)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		// fmt.Println(post)
	default:
		panic(err)
	}
	return post.Score
}

// ******** Get customer
func GetCustomers() []Customer {
	var customers []Customer
	db := getConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		customer := Customer{}
		err = rows.Scan(&customer.ID, &customer.Username, &customer.Password, &customer.Email)
		if err != nil {
			// handle this error
			panic(err)
		}
		customers = append(customers, customer)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return customers
}

func GetCustomer(customerID int) Customer {
	db := getConnection()
	defer db.Close()

	var customer Customer
	sqlStatement := `SELECT *  FROM customer where id = $1`
	row := db.QueryRow(sqlStatement, customerID)
	err := row.Scan(&customer.ID, &customer.Username, &customer.Password, &customer.Email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return Customer{}
	case nil:
		fmt.Println(customer)
	default:
		panic(err)
	}
	return customer
}
