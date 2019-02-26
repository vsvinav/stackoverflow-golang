package main

type Badges struct {
	Badges []Badge `xml:"row"`
}

type Badge struct {
	ID       int64  `xml:"Id,attr"`
	UserID   int64  `xml:"UserId,attr"`
	Name     string `xml:"Name,attr"`
	Date     string `xml:"Date,attr"`
	Class    int64  `xml:"Class,attr"`
	TagBased bool   `xml:"TagBased,attr"`
}

//************ 	Comments Structure

type Comments struct {
	Comments []Comment `xml:"row"`
}

type Comment struct {
	ID           int64  `xml:"Id,attr"`
	PostID       int64  `xml:"PostId,attr"`
	Score        int    `xml:"Score,attr"`
	Text         string `xml:"Text,attr"`
	CreationDate string `xml:"CreationDate,attr"`
	UserID       int64  `xml:"UserId,attr"`
}

//********** Post History Structure

type Posthistory struct {
	Posthistory []Posthist `xml:"row"`
}

type Posthist struct {
	ID                int64  `xml:"Id,attr"`
	PostHistoryTypeID int64  `xml:"PostHistoryTypeId,attr"`
	PostID            int    `xml:"PostId,attr"`
	RevisionGUID      string `xml:"RevisionGUID,attr"`
	CreationDate      string `xml:"CreationDate,attr"`
	UserID            int64  `xml:"UserId,attr"`
	Text              string `xml:"Text,attr"`
	Comment           string `xml:"Comment,attr"`
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

// ************* Posts Structure

type Posts struct {
	Posts []Post `xml:"row"`
}

type Post struct {
	ID               int    `xml:"Id,attr"`
	PostTypeID       int    `xml:"PostTypeId,attr"`
	CreationDate     string `xml:"CreationDate,attr"`
	Score            int    `xml:"Score,attr"`
	ViewCount        int    `xml:"ViewCount,attr"`
	Body             string `xml:"Body,attr"`
	OwnerUserID      int    `xml:"OwnerUserId,attr"`
	LastActivityDate string `xml:"LastActivityDate,attr"`
	Title            string `xml:"Title,attr"`
	Tags             string `xml:"Tags,attr"`
	AnswerCount      int    `xml:"AnswerCount,attr"`
	CommentCount     int    `xml:"CommentCount,attr"`
	FavoriteCount    int    `xml:"FavoriteCount,attr"`
	ClosedDate       string `xml:"ClosedDate,attr"`
}
