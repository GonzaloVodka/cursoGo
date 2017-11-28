package service

import (
	"fmt"

	"github.com/cursoGo/src/domain"
)

//Manager of tweets
type Manager struct {
	tweets       []domain.Tweet
	users        []domain.User
	loggedInUser domain.User
}

//GetLoggedInUser getter
func (manager *Manager) GetLoggedInUser() domain.User {
	return manager.loggedInUser
}

//InitializeService initializes the service
func (manager *Manager) InitializeService() {
	manager.tweets = make([]domain.Tweet, 0)
	manager.users = make([]domain.User, 0)
	domain.ResetCurrentID()
	manager.Logout()
}

//Register register a user
func (manager *Manager) Register(userToRegister domain.User) error {
	if userToRegister.Name == "" {
		return fmt.Errorf("Name is required")
	}

	if manager.IsRegistered(userToRegister) {
		return fmt.Errorf("The user is already registered")
	}

	manager.users = append(manager.users, userToRegister)
	return nil
}

//IsRegistered verify that a user is registered
func (manager *Manager) IsRegistered(user domain.User) bool {
	for _, u := range manager.users {
		if u.Name == user.Name {
			return true
		}
	}
	return false
}

//Login logs the user in
func (manager *Manager) Login(user domain.User) error {
	if manager.IsLoggedIn() {
		return fmt.Errorf("Already logged in")
	}
	if !manager.IsRegistered(user) {
		return fmt.Errorf("The user is not registered")
	}

	manager.loggedInUser = user
	return nil
}

//Logout logs the user out
func (manager *Manager) Logout() error {
	if !manager.IsLoggedIn() {
		return fmt.Errorf("Not logged in")
	}
	manager.loggedInUser = domain.User{Name: ""}
	return nil
}

//IsLoggedIn checks if there is a logged in user
func (manager *Manager) IsLoggedIn() bool {
	return manager.loggedInUser.Name != ""
}

//GetTweets returns all tweets.
func (manager *Manager) GetTweets() []domain.Tweet {
	return manager.tweets
}

//GetTweet returns the last published Tweet
// func GetTweet() domain.Tweet {
// 	return tweets[len(tweets)-1]
// }

//GetTweetByID returns the tweet that has that ID
func (manager *Manager) GetTweetByID(id int) (*domain.Tweet, error) {
	for _, tweet := range manager.tweets {
		if tweet.ID == id {
			return &tweet, nil
		}
	}
	return nil, fmt.Errorf("A tweet with that ID does not exist")
}

//GetTimelineFromUser returns all tweets from one user
func (manager *Manager) GetTimelineFromUser(user domain.User) ([]domain.Tweet, error) {
	if !manager.IsRegistered(user) {
		return nil, fmt.Errorf("That user is not registered")
	}

	var timeline []domain.Tweet
	for _, t := range manager.tweets {
		if t.User.Name == user.Name {
			timeline = append(timeline, t)
		}
	}
	return timeline, nil
}

//GetTimeline returns the loggedInUser's timeline
func (manager *Manager) GetTimeline() ([]domain.Tweet, error) {
	if !manager.IsLoggedIn() {
		return nil, fmt.Errorf("No user logged in")
	}
	return manager.GetTimelineFromUser(manager.loggedInUser)
}

//PublishTweet Publishes a tweet
func (manager *Manager) PublishTweet(tweetToPublish *domain.Tweet) error {
	if manager.loggedInUser.Name != tweetToPublish.User.Name {
		return fmt.Errorf("You must be logged in to tweet")
	}

	if tweetToPublish.Text == "" {
		return fmt.Errorf("Text is required")
	}

	manager.tweets = append(manager.tweets, *tweetToPublish)
	return nil
}
