package main

import (
	"fmt"
	"testing"
)

func TestForUpvotes(t *testing.T) {
	// fmt.Println("I am going in ")
	posts := getPosts()
	fetchUpvote := getVotes(posts[0].ID)
	upvote(posts[0].ID)
	got := getVotes(posts[0].ID)
	if got == fetchUpvote+1 {
		fmt.Println("Test passed")
	} else {
		t.Errorf("upvotes = %d; want %d", fetchUpvote, got)
	}
}

func TestForDownvotes(t *testing.T) {
	// fmt.Println("I am going in ")
	posts := getPosts()
	fetchDownVote := getVotes(posts[0].ID)
	downvote(posts[0].ID)
	got := getVotes(posts[0].ID)
	if got == fetchDownVote-1 {
		fmt.Println("Test passed")
	} else {
		t.Errorf("downvotes = %d; want %d", fetchDownVote, got)
	}
}
