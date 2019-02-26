package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func parseBadges() {

	// Open our xmlFile
	xmlFile, err := os.Open("Badges.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Badges.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var b Badges
	xml.Unmarshal(byteValue, &b)
	for i := 0; i < len(b.Badges); i++ {
		addBadge(b.Badges[i])
	}

}

func parseComments() {

	// Open our xmlFile
	xmlFile, err := os.Open("Comments.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Comments.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var c Comments
	xml.Unmarshal(byteValue, &c)
	for i := 0; i < len(c.Comments); i++ {
		addComment(c.Comments[i])
	}

}
func parsePostHistory() {

	// Open our xmlFile
	xmlFile, err := os.Open("PostHistory.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened PostHistory.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var ph Posthistory
	xml.Unmarshal(byteValue, &ph)
	for i := 0; i < len(ph.Posthistory); i++ {
		addPostHistory(ph.Posthistory[i])
	}

}
func parsePostLinks() {

	// Open our xmlFile
	xmlFile, err := os.Open("PostLinks.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened PostLinks.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var plink Postlinks
	fmt.Println(plink)
	xml.Unmarshal(byteValue, &plink)
	for i := 0; i < len(plink.Postlinks); i++ {
		addPostLinks(plink.Postlinks[i])
	}
	defer xmlFile.Close()

}
func parsePosts() {

	// Open our xmlFile
	xmlFile, err := os.Open("Posts.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Posts.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var posts Posts
	xml.Unmarshal(byteValue, &posts)
	for i := 0; i < len(posts.Posts); i++ {
		addPosts(posts.Posts[i])
	}
	defer xmlFile.Close()

}

func parseTags() {

	// Open our xmlFile
	xmlFile, err := os.Open("Tags.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Tags.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var tags Tags
	xml.Unmarshal(byteValue, &tags)
	for i := 0; i < len(tags.Tags); i++ {
		addTags(tags.Tags[i])
	}
	defer xmlFile.Close()

}

func parseUsers() {

	// Open our xmlFile
	xmlFile, err := os.Open("Users.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var users Users
	xml.Unmarshal(byteValue, &users)
	for i := 0; i < len(users.Users); i++ {
		addUsers(users.Users[i])
	}
	defer xmlFile.Close()

}

func parseVotes() {

	// Open our xmlFile
	xmlFile, err := os.Open("Votes.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Votes.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var votes Votes
	xml.Unmarshal(byteValue, &votes)
	for i := 0; i < len(votes.Votes); i++ {
		addVotes(votes.Votes[i])
	}
	defer xmlFile.Close()

}
