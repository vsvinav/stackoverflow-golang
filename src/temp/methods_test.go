package main

import (
	"fmt"
	"testing"
)

func TestForUpvotes(t *testing.T) {
	// fmt.Println("I am going in ")
	posts := GetPosts()
	fetchUpvote := GetVotes(posts[0].ID)
	upvote(posts[0].ID)
	got := GetVotes(posts[0].ID)
	if got == fetchUpvote+1 {
		fmt.Println("Test passed")
	} else {
		t.Errorf("upvotes = %d; want %d", fetchUpvote, got)
	}
}

func TestForDownvotes(t *testing.T) {
	// fmt.Println("I am going in ")
	posts := GetPosts()
	fetchDownVote := GetVotes(posts[0].ID)
	downvote(posts[0].ID)
	got := GetVotes(posts[0].ID)
	if got == fetchDownVote-1 {
		fmt.Println("Test passed")
	} else {
		t.Errorf("downvotes = %d; want %d", fetchDownVote, got)
	}

}

func TestForGetUser(t *testing.T) {
	users := GetUsers()
	got := users[1]
	expected := GetUser(1)
	fmt.Println(got, expected)
	if got == expected {
		fmt.Println("Test Passed")

	} else {
		t.Errorf("expected = %v; got = %v", expected, got)
	}

}

func TestForGetPost(t *testing.T) {
	posts := GetPosts()
	got := posts[0]
	expected := GetPostForTest(posts[0].ID)
	fmt.Println(got, expected)
	if got == expected {
		fmt.Println("Test Passed")

	} else {
		t.Errorf("expected = %v; got = %v", expected, got)
	}

}

func TestCustomerUserName(t *testing.T) {
	var cust Customer = 
}