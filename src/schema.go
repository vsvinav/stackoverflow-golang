package main

import "encoding/xml"

type Badges struct {
	XMLName xml.Name `xml:"badges"`
	Badges  []Badge  `xml:"row"`
}

type Badge struct {
	XMLName  xml.Name `xml:"row"`
	ID       int64    `xml:"Id,attr"`
	UserID   int64    `xml:"UserId,attr"`
	Name     string   `xml:"Name,attr"`
	Date     string   `xml:"Date,attr"`
	Class    int64    `xml:"Class,attr"`
	TagBased bool     `xml:"TagBased,attr"`
}

//************ 	Comments Structure

type Comments struct {
	XMLName  xml.Name  `xml:"comments"`
	Comments []Comment `xml:"row"`
}

type Comment struct {
	XMLName      xml.Name `xml:"row"`
	ID           int64    `xml:"Id,attr"`
	PostID       int64    `xml:"PostId,attr"`
	Score        int      `xml:"Score,attr"`
	Text         string   `xml:"Text,attr"`
	CreationDate string   `xml:"CreationDate,attr"`
	UserID       int64    `xml:"UserId,attr"`
}

//********** Post History Structure

type Posthistory struct {
	XMLName     xml.Name   `xml:"posthistory"`
	Posthistory []Posthist `xml:"row"`
}

type Posthist struct {
	XMLName           xml.Name `xml:"row"`
	ID                int64    `xml:"Id,attr"`
	PostHistoryTypeID int64    `xml:"PostHistoryTypeId,attr"`
	PostID            int      `xml:"PostId,attr"`
	RevisionGUID      string   `xml:"RevisionGUID,attr"`
	CreationDate      string   `xml:"CreationDate,attr"`
	UserID            int64    `xml:"UserId,attr"`
	Text              string   `xml:"Text,attr"`
	Comment           string   `xml:"Comment,attr"`
}

// ************ Post Links Structure
type Postlinks struct {
	Postlinks []Postlink `xml:"row"`
}

type Postlink struct {
	ID            int    `xml:"Id,attr"`
	PostID        int    `xml:"PostId,attr"`
	CreationDate  string `xml:"CreationDate,attr"`
	RelatedPostID int    `xml:"RelatedPostId,attr"`
	LinkTypeID    int    `xml:"LinkTypeId,attr"`
}
