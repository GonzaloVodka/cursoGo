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

//GetText return the text of the tweet
func (t TextTweet) GetText() string {
	return t.text
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

//StringTweet returns a tweet as a formatted string
// func (tw Tweet) StringTweet() string {
// 	date := tw.Date.Format("Mon Jan _2 15:04:05 2006")
// 	st := "TweetID: " + strconv.Itoa(tw.ID) + " Nick: " + tw.User.Nick + ": " + tw.Text + ", " + date
// 	return st
// }
