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
	var cust Customer = Customer{ID: 12, Username: "test", Password: "1234556666666666", Email: "something"}

	got := ValidateUsername(cust.Username)
	expected := true
	if got != expected {
		fmt.Println("Test Passed")

	} else {
		t.Errorf("expected = %v; got = %v", expected, got)
	}
}

func TestCustomerEmail(t *testing.T) {
	var cust Customer = Customer{ID: 12, Username: "test", Password: "1234556666666666", Email: "something"}

	got := ValidateEmail(cust.Email)
	expected := true
	if got != expected {
		fmt.Println("Test Passed")

	} else {
		t.Errorf("expected = %v; got = %v", expected, got)
	}
}

func TestCustomerPassword(t *testing.T) {
	var cust Customer = Customer{ID: 12, Username: "test", Password: "1234556666666666", Email: "something"}

	got := ValidatePassword(cust.Password)
	expected := true
	if got != expected {
		fmt.Println("Test Passed")

	} else {
		t.Errorf("expected = %v; got = %v", expected, got)
	}
}

func TestForGetComment(t *testing.T) {
	comments := GetComments()
	got := comments[1]
	expected := GetComment(comments[1].PostID)
	fmt.Println(got, expected)
	if got == expected {
		fmt.Println("Test Passed")

	} else {
		t.Errorf("expected = %v; got = %v", expected, got)
	}

}
