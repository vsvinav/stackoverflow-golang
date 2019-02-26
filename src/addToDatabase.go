package main

import (
	"fmt"
	"regexp"
)

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

func addCustomer(customer Customer) {
	db := getConnection()
	defer db.Close()

	sqlStatement := `
INSERT INTO customer(id, username, password, email)
VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlStatement, customer.ID, customer.username, customer.password, customer.email)
	if err != nil {
		panic(err)
	}

}

func takeCustomerInput() Customer {
	var customer Customer
	fmt.Println("Enter id:")
	fmt.Scan(&customer.ID)
	fmt.Println("Enter username:")
	var re = regexp.MustCompile(`(?m)^\w{5,}$`)

	fmt.Scanln(&customer.username)
	fmt.Println("Enter Password:")
	fmt.Scanln(&customer.password)
	fmt.Println("Enter email")
	fmt.Scanln(&customer.email)
	return customer
}
