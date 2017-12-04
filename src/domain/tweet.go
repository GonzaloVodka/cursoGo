package domain

import (
	"fmt"
	"time"
)

//Tweet interface
type Tweet interface {
	PrintableTweet() string
	GetUser() User
	GetText() string
	GetDate() *time.Time
	GetID() int
	SetUser(User)
	SetText(string)
	String() string
	Equals(Tweet) bool
}

//TextTweet is a tweet
type TextTweet struct {
	user User
	text string
	date *time.Time
	iD   int
}

//PrintableTweet return a string to print the tweet
func (t TextTweet) PrintableTweet() string {
	s := "@" + t.user.Nick + ": " + t.text
	return s
}

//GetUser return the user of the tweet
func (t TextTweet) GetUser() User {
	return t.user
}

//SetUser is a setter
func (t TextTweet) SetUser(u User) {
	t.user = u
}

//GetText return the text of the tweet
func (t TextTweet) GetText() string {
	return t.text
}

//SetText is a setter
func (t TextTweet) SetText(text string) {
	t.text = text
}

//GetDate return the date of the tweet
func (t TextTweet) GetDate() *time.Time {
	return t.date
}

//GetID return the id of the tweet
func (t TextTweet) GetID() int {
	return t.iD
}

var currentID = -1

//getNextID returns the id of the next tweet
func getNextID() int {
	currentID++
	return (currentID)
}

//ResetCurrentID serves as an initialization, resetting the current ID
func ResetCurrentID() {
	currentID = -1
}

//NewTextTweet creates a TextTweet
func NewTextTweet(usr User, txt string) (Tweet, error) {
	if usr.Name == "" {
		return nil, fmt.Errorf("You must be logged in")
	}
	now := time.Now()
	if len(txt) > 140 {
		return nil, fmt.Errorf("Can't have more than 140 characters")
	}
	tw := TextTweet{user: usr, text: txt, date: &now, iD: getNextID()}
	return tw, nil
}

func (t TextTweet) String() string {
	return t.PrintableTweet()
}

//Equals compares tweets
func (t TextTweet) Equals(tw Tweet) (b bool) {
	if t.GetUser().Nick == tw.GetUser().Nick && t.GetText() == tw.GetText() {
		b = true
	}
	return b
}

//StringTweet returns a tweet as a formatted string
// func (tw Tweet) StringTweet() string {
// 	date := tw.Date.Format("Mon Jan _2 15:04:05 2006")
// 	st := "TweetID: " + strconv.Itoa(tw.ID) + " Nick: " + tw.User.Nick + ": " + tw.Text + ", " + date
// 	return st
// }

//ImageTweet is a tweet with a image in it
type ImageTweet struct {
	TextTweet
	url string
}

//PrintableTweet return a string to print the tweet
func (t ImageTweet) PrintableTweet() string {
	s := "@" + t.user.Nick + ": " + t.text + " " + t.url
	return s
}

//GetUser return the user of the tweet
func (t ImageTweet) GetUser() User {
	return t.user
}

//SetUser is a setter
func (t ImageTweet) SetUser(u User) {
	t.user = u
}

//GetText return the text of the tweet
func (t ImageTweet) GetText() string {
	return t.text
}

//SetText is a setter
func (t ImageTweet) SetText(text string) {
	t.text = text
}

//GetDate return the date of the tweet
func (t ImageTweet) GetDate() *time.Time {
	return t.date
}

//GetURL return the user of the tweet
func (t ImageTweet) GetURL() string {
	return t.url
}

//GetID return the id of the tweet
func (t ImageTweet) GetID() int {
	return t.iD
}

//NewImageTweet creates a TextTweet
func NewImageTweet(usr User, txt string, url string) (Tweet, error) {
	if usr.Name == "" {
		return nil, fmt.Errorf("You must be logged in")
	}

	if len(txt) > 140 {
		return nil, fmt.Errorf("Can't have more than 140 characters")
	}

	now := time.Now()

	tw := TextTweet{user: usr, text: txt, date: &now, iD: getNextID()}
	itw := ImageTweet{TextTweet: tw, url: url}
	return itw, nil
}

func (t ImageTweet) String() string {
	return t.PrintableTweet()
}

//Equals compares tweets
func (t ImageTweet) Equals(tw Tweet) (b bool) {
	c, ok := tw.(ImageTweet)
	if !ok {
		return b
	}
	if t.GetUser().Nick == c.GetUser().Nick && t.GetText() == c.GetText() && t.url == c.url {
		b = true
	}
	return b
}

//QuoteTweet is a tweet with a quoted tweet
type QuoteTweet struct {
	TextTweet
	quotedtweet Tweet
}

//PrintableTweet return a string to print the tweet
func (t QuoteTweet) PrintableTweet() string {
	s := t.TextTweet.PrintableTweet() + ` "` + t.quotedtweet.PrintableTweet() + `"`
	return s
}

//GetUser return the user of the tweet
func (t QuoteTweet) GetUser() User {
	return t.TextTweet.GetUser()
}

//SetUser is a setter
func (t QuoteTweet) SetUser(u User) {
	t.TextTweet.user = u
}

//GetText return the text of the tweet
func (t QuoteTweet) GetText() string {
	return t.TextTweet.GetText()
}

//SetText is a setter
func (t QuoteTweet) SetText(text string) {
	t.TextTweet.text = text
}

//GetDate return the date of the tweet
func (t QuoteTweet) GetDate() *time.Time {
	return t.TextTweet.GetDate()
}

//GetURL return the user of the tweet
// func (t QuoteTweet) GetURL() string {
// 	return t.Tweet.GetURL()
// }

//GetID return the id of the tweet
func (t QuoteTweet) GetID() int {
	return t.quotedtweet.GetID()
}

//NewQuoteTweet creates a QuoteTweet
func NewQuoteTweet(usr User, text string, quotedtw Tweet) (Tweet, error) {
	if usr.Name == "" {
		return nil, fmt.Errorf("You must be logged in")
	}

	if len(text) > 140 {
		return nil, fmt.Errorf("Can't have more than 140 characters")
	}

	now := time.Now()

	tw := TextTweet{user: usr, text: text, date: &now, iD: getNextID()}
	qtw := QuoteTweet{TextTweet: tw, quotedtweet: quotedtw}
	return qtw, nil
}

func (t QuoteTweet) String() string {
	return t.PrintableTweet()
}

//Equals compares tweets
func (t QuoteTweet) Equals(tw Tweet) (b bool) {
	c, ok := tw.(QuoteTweet)
	if !ok {
		return b
	}
	if t.GetUser().Nick == c.GetUser().Nick && t.GetText() == c.GetText() && t.quotedtweet.Equals(c.quotedtweet) {
		b = true
	}
	return b
}
