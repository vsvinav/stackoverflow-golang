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
	var b badges
	xml.Unmarshal(byteValue, &b)
	for i := 0; i < len(b.badges); i++ {
		fmt.Println(b.badges)
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
	defer xmlFile.Close()

}
