package main

import (
	"fmt"
	"os"
)

func ParseBadges() {

	// Open our xmlFile
	xmlFile, err := os.Open("Badges.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Badges.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

}

func ParseComments() {

	// Open our xmlFile
	xmlFile, err := os.Open("Comments.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Comments.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

}
func ParsePostHistory() {

	// Open our xmlFile
	xmlFile, err := os.Open("PostHistory.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened PostHistory.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

}
func ParsePostLinks() {

	// Open our xmlFile
	xmlFile, err := os.Open("PostLinks.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened PostLinks.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

}
func ParsePosts() {

	// Open our xmlFile
	xmlFile, err := os.Open("Posts.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Posts.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

}

func ParseTags() {

	// Open our xmlFile
	xmlFile, err := os.Open("Tags.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Tags.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

}

func ParseUsers() {

	// Open our xmlFile
	xmlFile, err := os.Open("Users.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

}

func ParseVotes() {

	// Open our xmlFile
	xmlFile, err := os.Open("Votes.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Votes.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

}
