package domain

import (
	"fmt"
	"strconv"
	"time"
)

//Tweet is a tweet
type Tweet struct {
	User User
	Text string
	Date *time.Time
	ID   int
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

//NewTweet creates a tweet
func NewTweet(usr User, txt string) (*Tweet, error) {
	if usr.Name == "" {
		return nil, fmt.Errorf("You must be logged in")
	}
	now := time.Now()
	if len(txt) > 140 {
		return nil, fmt.Errorf("Can't have more than 140 characters")
	}
	tw := Tweet{User: usr, Text: txt, Date: &now, ID: getNextID()}
	return &tw, nil
}

//StringTweet returns a tweet as a formatted string
func (tw Tweet) StringTweet() string {
	date := tw.Date.Format("Mon Jan _2 15:04:05 2006")
	st := "TweetID: " + strconv.Itoa(tw.ID) + " Nick: " + tw.User.Nick + ": " + tw.Text + ", " + date
	return st
}
