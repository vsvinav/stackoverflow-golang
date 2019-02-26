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

func addBadge(badge Badge) {
	db := getConnection()
	defer db.Close()

	sqlStatement := `
INSERT INTO badges (id, user_id, name, date, class, tagbased)
VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement, badge.ID, badge.UserID, badge.Name, badge.Date, badge.Class, badge.TagBased)
	if err != nil {
		panic(err)
	}

	fmt.Println("Added badge to database", badge)
}

func addComment(comment Comment) {
	db := getConnection()
	defer db.Close()

	sqlStatement := `
INSERT INTO comments (id, post_id, score, text, creation_date, user_id)
VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement, comment.ID, comment.PostID, comment.Score, comment.Text, comment.CreationDate, comment.UserID)
	if err != nil {
		panic(err)
	}

}

func addPostHistory(posthist Posthist) {
	db := getConnection()
	defer db.Close()

	sqlStatement := `
INSERT INTO post_history (id, post_history_type_id, post_id, revision_guid, creation_date, user_id, text, comment)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Exec(sqlStatement, posthist.ID, posthist.PostHistoryTypeID, posthist.PostID, posthist.RevisionGUID, posthist.CreationDate, posthist.UserID, posthist.Text, posthist.Comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(posthist.ID)

}

func addPostLinks(postlink Postlink) {
	db := getConnection()
	defer db.Close()

	sqlStatement := `
INSERT INTO post_link (id, creation_date, post_id, related_post_id, link_type_id)
VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, postlink.ID, postlink.CreationDate, postlink.PostID, postlink.RelatedPostID, postlink.LinkTypeID)
	if err != nil {
		panic(err)
	}
	fmt.Println(postlink.ID)

}
func addPosts(post Post) {
	db := getConnection()
	defer db.Close()

	sqlStatement := `
INSERT INTO posts (id, post_type_id, parent_id, accepted_answer_id, creation_date, score, view_count, body, owner_user_id, owner_display_name,last_editor_user_id, last_edit_date, last_activity_date, last_editor_display_name, title, tags, answer_count, comment_count, favourite_count, closed_date, community_owned_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)`
	_, err := db.Exec(sqlStatement, post.ID, post.PostTypeID, post.ParentID, post.AcceptedAnswerID, post.CreationDate, post.Score, post.ViewCount, post.Body, post.OwnerUserID, post.OwnerDisplayName, post.LastEditorUserID, post.LastEditDate, post.LastActivityDate, post.LastEditorDisplayName, post.Title, post.Tags, post.AnswerCount, post.CommentCount, post.FavoriteCount, post.ClosedDate, post.CommunityOwnedDate)
	if err != nil {
		panic(err)
	}
	fmt.Println(post.ID)

}
func addTags(tag Tag) {
	db := getConnection()
	defer db.Close()

	sqlStatement := `
INSERT INTO tags (id, tag_name, count, excerpt_post_id, wiki_post_id)
VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, tag.ID, tag.TagName, tag.Count, tag.ExcerptPostID, tag.WikiPostID)
	if err != nil {
		panic(err)
	}
	fmt.Println(tag.ID)

}
func addUsers(user User) {
	db := getConnection()
	defer db.Close()

	sqlStatement := `
INSERT INTO users (id , reputation , creation_date, display_name , last_access_date, website_url , location, about_me, views , upvotes ,downvotes , account_id , profile_image_url )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	_, err := db.Exec(sqlStatement, user.ID, user.Reputation, user.CreationDate, user.DisplayName, user.LastAccessDate, user.WebsiteURL, user.Location, user.AboutMe, user.Views, user.UpVotes, user.DownVotes, user.AccountID, user.ProfileImageURL)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.ID)

}

func addVotes(vote Vote) {
	db := getConnection()
	defer db.Close()

	sqlStatement := `
INSERT INTO votes(id, post_id , creation_date , vote_type_id , user_id , bounty_amount )
VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement, vote.ID, vote.PostID, vote.CreationDate, vote.VoteTypeID, vote.UserID, vote.BountyAmount)
	if err != nil {
		panic(err)
	}
	fmt.Println(vote.ID)

}

func getPosts() []Post {
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

func getPost(postID int) Post {
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
	return post
}

// ************* Get Comments for the post

func getComments() []Comment {
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

func getComment(postID int) Comment {
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

func getAnswers() []Post {
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

func getAnswer(postID int, choice string) Post {
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

func main() {

	getConnection()
	connectToDatabase()
	// // Milestone 2 Q1:
	posts := getPosts()
	// fmt.Println(getPost(posts[0].ID))
	// Milestone 2 Q2:
	// comments := getComments()
	// fmt.Println(getPost(comments[6].PostID))
	// Milestone 2 Q3:
	fmt.Println(getAnswer(posts[5].ID, "score"))
	fmt.Println(getAnswer(posts[5].ID, "creation date"))
	fmt.Println(getAnswer(posts[5].ID, "last activity date"))
}
