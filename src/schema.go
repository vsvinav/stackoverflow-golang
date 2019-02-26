package main

type Badges struct {
	Badges []Badge `xml:"row"`
}

type Badge struct {
	ID       int    `xml:"Id,attr"`
	UserID   int    `xml:"UserId,attr"`
	Name     string `xml:"Name,attr"`
	Date     string `xml:"Date,attr"`
	Class    int    `xml:"Class,attr"`
	TagBased bool   `xml:"TagBased,attr"`
}

//************ 	Comments Structure

type Comments struct {
	Comments []Comment `xml:"row"`
}

type Comment struct {
	ID           int    `xml:"Id,attr"`
	PostID       int    `xml:"PostId,attr"`
	Score        int    `xml:"Score,attr"`
	Text         string `xml:"Text,attr"`
	CreationDate string `xml:"CreationDate,attr"`
	UserID       int    `xml:"UserId,attr"`
}

//********** Post History Structure

type Posthistory struct {
	Posthistory []Posthist `xml:"row"`
}

type Posthist struct {
	ID                int    `xml:"Id,attr"`
	PostHistoryTypeID int    `xml:"PostHistoryTypeId,attr"`
	PostID            int    `xml:"PostId,attr"`
	RevisionGUID      string `xml:"RevisionGUID,attr"`
	CreationDate      string `xml:"CreationDate,attr"`
	UserID            int    `xml:"UserId,attr"`
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
	ID                    int    `xml:"Id,attr"`
	PostTypeID            int    `xml:"PostTypeId,attr"`
	ParentID              int    `xml:"ParentId,attr"`
	AcceptedAnswerID      int    `xml:"AcceptedAnswerId,attr"`
	CreationDate          string `xml:"CreationDate,attr"`
	Score                 int    `xml:"Score,attr"`
	ViewCount             int    `xml:"ViewCount,attr"`
	Body                  string `xml:"Body,attr"`
	OwnerUserID           int    `xml:"OwnerUserId,attr"`
	OwnerDisplayName      string `xml:"OwnerDisplayName,attr"`
	LastEditorUserID      int    `xml:"LastEditorUserId,attr"`
	LastEditDate          string `xml:"LastEditDate,attr"`
	LastActivityDate      string `xml:"LastActivityDate,attr"`
	LastEditorDisplayName string `xml:"LastEditorDisplayName,attr"`
	Title                 string `xml:"Title,attr"`
	Tags                  string `xml:"Tags,attr"`
	AnswerCount           int    `xml:"AnswerCount,attr"`
	CommentCount          int    `xml:"CommentCount,attr"`
	FavoriteCount         int    `xml:"FavoriteCount,attr"`
	ClosedDate            string `xml:"ClosedDate,attr"`
	CommunityOwnedDate    string `xml:"CommunityOwnedDate,attr"`
}

// *******Tags structure

type Tags struct {
	Tags []Tag `xml:"row"`
}

type Tag struct {
	ID            int    `xml:"Id,attr"`
	TagName       string `xml:"TagName,attr"`
	Count         int    `xml:"Count,attr"`
	ExcerptPostID int    `xml:"ExcerptPostId,attr"`
	WikiPostID    int    `xml:"WikiPostId,attr"`
}

// *******Users structure
type Users struct {
	Users []User `xml:"row"`
}

type User struct {
	ID              int    `xml:"Id,attr"`
	Reputation      int    `xml:"Reputation,attr"`
	CreationDate    string `xml:"CreationDate,attr"`
	DisplayName     string `xml:"DisplayName,attr"`
	LastAccessDate  string `xml:"LastAccessDate,attr"`
	WebsiteURL      string `xml:"WebsiteUrl,attr"`
	Location        string `xml:"Location,attr"`
	AboutMe         string `xml:"AboutMe,attr"`
	Views           int    `xml:"Views,attr"`
	UpVotes         int    `xml:"UpVotes,attr"`
	DownVotes       int    `xml:"DownVotes,attr"`
	AccountID       int    `xml:"AccountId,attr"`
	ProfileImageURL string `xml:"ProfileImageUrl,attr"`
}

// *******Votes structure

type Votes struct {
	Votes []Vote `xml:"row"`
}

type Vote struct {
	ID           int    `xml:"Id,attr"`
	PostID       int    `xml:"PostId,attr"`
	CreationDate string `xml:"CreationDate,attr"`
	VoteTypeID   int    `xml:"VoteTypeId,attr"`
	UserID       int    `xml:"UserId,attr"`
	BountyAmount int    `xml:"BountyAmount,attr"`
}

// *********Customer Structure

type Customer struct {
	ID       int
	username string
	password string
	email    string
}
