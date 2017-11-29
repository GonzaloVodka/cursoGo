package service

import (
	"fmt"

	"github.com/cursoGo/src/domain"
)

//Manager of tweets
type Manager struct {
	users        []domain.User
	tweets       map[domain.User][]domain.Tweet
	loggedInUser domain.User
}

//GetLoggedInUser getter
func (manager *Manager) GetLoggedInUser() domain.User {
	return manager.loggedInUser
}

//InitializeService initializes the service
func (manager *Manager) InitializeService() {
	manager.users = make([]domain.User, 0)
	manager.tweets = make(map[domain.User][]domain.Tweet)
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
	manager.tweets[userToRegister] = make([]domain.Tweet, 0)
	return nil
}

//IsRegistered verify that a user is registered
func (manager *Manager) IsRegistered(user domain.User) bool {
	for _, u := range manager.users {
		if u.Nick == user.Nick {
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

	for _, u := range manager.users {
		if u.Nick == user.Nick && u.Pass == user.Pass {
			user = u
			manager.loggedInUser = user
			return nil
		}
	}

	return fmt.Errorf("The password is incorrect")
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

//GetTimelineFromUser returns all tweets from one user
func (manager *Manager) GetTimelineFromUser(user domain.User) ([]domain.Tweet, error) {
	if !manager.IsRegistered(user) {
		return nil, fmt.Errorf("That user is not registered")
	}

	timeline, _ := manager.tweets[user]
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
	if !manager.IsLoggedIn() {
		return fmt.Errorf("No user logged in")
	}

	if tweetToPublish.Text == "" {
		return fmt.Errorf("Text is required")
	}

	timeline, _ := manager.tweets[manager.GetLoggedInUser()]

	timeline = append(timeline, *tweetToPublish)

	manager.tweets[manager.GetLoggedInUser()] = timeline

	return nil
}

//DeleteTweet delete a tweet
func (manager *Manager) DeleteTweet(id int) error {
	if !manager.IsLoggedIn() {
		return fmt.Errorf("No user logged in")
	}

	timeline, _ := manager.tweets[manager.GetLoggedInUser()]
	var newTimeline = make([]domain.Tweet, 0)
	for _, tw := range timeline {
		if tw.ID != id {
			newTimeline = append(newTimeline, tw)
		}
	}
	return nil
}
