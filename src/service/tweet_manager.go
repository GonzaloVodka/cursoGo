package service

import (
	"fmt"

	"github.com/cursoGo/src/domain"
)

//Manager of tweets
type Manager struct {
	users        []*domain.User
	tweets       map[string][]domain.Tweet
	loggedInUser *domain.User
}

//GetLoggedInUser getter
func (manager *Manager) GetLoggedInUser() *domain.User {
	return manager.loggedInUser
}

//SetLoggedInUser a
func (manager *Manager) SetLoggedInUser(usr *domain.User) {
	manager.loggedInUser = usr
}

//InitializeService initializes the service
func (manager *Manager) InitializeService() {
	manager.users = make([]*domain.User, 0)
	manager.tweets = make(map[string][]domain.Tweet)
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

	manager.users = append(manager.users, &userToRegister)
	manager.tweets[userToRegister.Nick] = make([]domain.Tweet, 0)
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
			manager.SetLoggedInUser(u)
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
	manager.loggedInUser = nil
	return nil
}

//IsLoggedIn checks if there is a logged in user
func (manager *Manager) IsLoggedIn() (b bool) {
	if manager.GetLoggedInUser() != nil {
		b = true
	}
	return b
}

//GetTimelineFromUser returns all tweets from one user
func (manager *Manager) GetTimelineFromUser(user domain.User) ([]domain.Tweet, error) {
	if !manager.IsRegistered(user) {
		return nil, fmt.Errorf("That user is not registered")
	}

	timeline, _ := manager.tweets[user.Nick]
	return timeline, nil
}

//GetTimeline returns the loggedInUser's timeline
func (manager *Manager) GetTimeline() ([]domain.Tweet, error) {
	if !manager.IsLoggedIn() {
		return nil, fmt.Errorf("No user logged in")
	}
	return manager.GetTimelineFromUser(*manager.loggedInUser)
}

//PublishTweet Publishes a tweet
func (manager *Manager) PublishTweet(tweetToPublish domain.Tweet) error {
	if !manager.IsLoggedIn() {
		return fmt.Errorf("You must be logged in to tweet")
	}

	if tweetToPublish.GetText() == "" {
		return fmt.Errorf("Text is required")
	}

	if manager.DuplicateTweet(tweetToPublish) {
		return fmt.Errorf("Tweet duplicated")
	}

	timeline, _ := manager.tweets[manager.GetLoggedInUser().Nick]

	timeline = append(timeline, tweetToPublish)

	manager.tweets[manager.GetLoggedInUser().Nick] = timeline

	return nil
}

//DeleteTweet delete a tweet
func (manager *Manager) DeleteTweet(id int) (string, error) {
	s := "No tweet deleted"
	if !manager.IsLoggedIn() {
		return s, fmt.Errorf("No user logged in")
	}

	timeline, _ := manager.tweets[manager.GetLoggedInUser().Nick]

	var newTimeline = make([]domain.Tweet, 0)

	for _, tw := range timeline {
		if tw.GetID() != id {
			newTimeline = append(newTimeline, tw)
		} else {
			s = "A tweet was deleted"
		}
	}
	return s, nil
}

//DuplicateTweet validate if a tweet is duplicated
func (manager *Manager) DuplicateTweet(tw domain.Tweet) (b bool) {
	timeline, _ := manager.GetTimeline()
	for _, tweet := range timeline {
		if tw.Equals(tweet) {
			b = true
		}
	}
	return b
}

//Follow a user
func (manager *Manager) Follow(nick string) error {
	fakeuser := domain.NewUser("", "", nick, "")
	if !manager.IsRegistered(fakeuser) {
		return fmt.Errorf("The user does not exist")
	}

	for _, user := range manager.users {
		if user.Nick == nick {
			manager.GetLoggedInUser().Following = append(manager.GetLoggedInUser().Following, *user)
			user.Followers = append(user.Followers, *manager.GetLoggedInUser())
		}
	}

	return nil
}

//GetFollowers from the logged user
func (manager *Manager) GetFollowers() []domain.User {
	return manager.GetLoggedInUser().Followers
}

//GetFollowings from the logged user
func (manager *Manager) GetFollowings() []domain.User {
	return manager.GetLoggedInUser().Following
}
