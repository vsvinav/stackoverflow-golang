package main

type Badges struct {
	XMLName xml.name `xml:"Badges.xml"`
	Badges  []Badge  `xml:""`
}
