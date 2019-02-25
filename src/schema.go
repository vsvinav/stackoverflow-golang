package main

import "encoding/xml"

type badges struct {
	XMLName xml.Name `xml:"badges"`
	badges  []badge  `xml:"row"`
}

type badge struct {
	XMLName  xml.Name `xml:"row"`
	ID       int64    `xml:"Id,attr"`
	UserID   int64    `xml:"UserId,attr"`
	Name     string   `xml:"Name,attr"`
	Date     string   `xml:"Date,attr"`
	Class    int64    `xml:"Class,attr"`
	TagBased bool     `xml:"TagBased,attr"`
}
