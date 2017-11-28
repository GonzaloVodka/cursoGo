package service_test

import (
	"testing"

	"github.com/cursoGo/src/domain"
	"github.com/cursoGo/src/service"
)

func isValidTweet(t *testing.T, publishedTweet domain.Tweet, user domain.User, text string) bool {
	if publishedTweet.User.Name != user.Name && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user.Name, text, publishedTweet.User.Name, publishedTweet.Text)
		return false
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}
	return true
}

func validateExpectedError(t *testing.T, err error, expectedError string) {
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error is '%s', but was %s", expectedError, err.Error())
		return
	}
}

func TestCantLoginIfAlreadyLoggedIn(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()

	user := domain.NewUser("root")
	manager.Register(user)
	manager.Login(user)

	//Operation
	err := manager.Login(user)

	//Validation
	validateExpectedError(t, err, "Already logged in")
}
func TestCantLogInWithUnregisteredUser(t *testing.T) {
	//Initialization
	var manager service.Manager

	manager.InitializeService()
	user := domain.NewUser("root")

	//Operation
	err := manager.Login(user)

	//Validation
	validateExpectedError(t, err, "The user is not registered")

}

// func TestPublishedTweetIsSaved(t *testing.T) {
// 	//Initialization
// 	var manager service.Manager
// 	manager.InitializeService()

// 	var tweet *domain.Tweet
// 	user := domain.NewUser("root")
// 	manager.Register(user)
// 	manager.Login(user)
// 	text := "This is my first tweet"
// 	tweet, _ = domain.NewTweet(user, text)
// 	//Operation
// 	err := manager.PublishTweet(tweet)

// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}

// 	//Validation
// 	publishedTweet := service.GetTweet()
// 	isValidTweet(t, publishedTweet, user, text)
// }

func TestMustBeLoggedInToPublishTweet(t *testing.T) {
	//Initalization
	var manager service.Manager
	manager.InitializeService()

	var tweet *domain.Tweet
	user := domain.NewUser("root")
	manager.Register(user)

	text := "This is my first tweet"
	tweet, _ = domain.NewTweet(user, text)
	//Operation
	err := manager.PublishTweet(tweet)
	validateExpectedError(t, err, "You must be logged in to tweet")

}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	//Initialization
	var manager service.Manager
	manager.InitializeService()

	var tweet *domain.Tweet

	user := domain.NewUser("Gonzalo")
	manager.Register(user)
	manager.Login(user)
	var text string

	tweet, _ = domain.NewTweet(user, text)

	//Operation
	err := manager.PublishTweet(tweet)

	//Validation
	validateExpectedError(t, err, "Text is required")
}

func TestCanPublishAndRetriveMoreThanOneTweet(t *testing.T) {

	//Initialization
	var manager service.Manager
	manager.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := domain.NewUser("Manuel")
	manager.Register(user)
	manager.Login(user)
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet, _ = domain.NewTweet(user, text)
	secondTweet, _ = domain.NewTweet(user, secondText)

	//Operation
	manager.PublishTweet(tweet)
	manager.PublishTweet(secondTweet)

	//Validation
	publishedTweets := manager.GetTweets()

	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, user, text) {
		return
	}
	isValidTweet(t, secondPublishedTweet, user, secondText)
}

func TestCanRegisterUser(t *testing.T) {

	//Initialization
	var manager service.Manager
	manager.InitializeService()
	user := domain.NewUser("Gonza")
	//Operation
	manager.Register(user)
	//Validation
	if !manager.IsRegistered(user) {
		t.Error("User did not get registered")
	}
}

func TestCantRegisterInvalidUser(t *testing.T) {
	//Initalization
	var manager service.Manager
	manager.InitializeService()
	var user domain.User
	//Operation
	err := manager.Register(user)
	//Validation
	validateExpectedError(t, err, "Name is required")
}

func TestCantRegisterSameUserMoreThanOnce(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()
	user := domain.NewUser("Gonza")
	//Operation
	manager.Register(user)
	err := manager.Register(user)
	//Validation
	validateExpectedError(t, err, "The user is already registered")
}

func TestCanRetrieveTimeline(t *testing.T) {

	//Initialization
	var manager service.Manager
	manager.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := domain.NewUser("Manuel")
	manager.Register(user)

	secondUser := domain.NewUser("Gonzalo")
	manager.Register(secondUser)

	text := "This is my first tweet"
	secondText := "This is my second tweet"
	thirdText := "This is a tweet"

	tweet, _ = domain.NewTweet(user, text)
	secondTweet, _ = domain.NewTweet(user, secondText)
	thirdTweet, _ = domain.NewTweet(secondUser, thirdText)

	manager.Login(secondUser)
	manager.PublishTweet(thirdTweet)
	manager.Logout()

	manager.Login(user)
	manager.PublishTweet(tweet)
	manager.PublishTweet(secondTweet)

	//Operation
	publishedTweets, _ := manager.GetTimeline()

	//Validation
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	for _, tweet := range publishedTweets {
		if tweet.User.Name != user.Name {
			t.Errorf("Expected user is %s but was %s", user.Name, tweet.User.Name)
		}
	}
}

func TestCantRetrieveTimelineWithoutLoggingIn(t *testing.T) {

	//Initialization
	var manager service.Manager
	manager.InitializeService()

	var tweet *domain.Tweet

	user := domain.NewUser("Manuel")
	manager.Register(user)
	manager.Login(user)

	text := "This is my first tweet"
	tweet, _ = domain.NewTweet(user, text)

	manager.PublishTweet(tweet)
	manager.Logout()

	//Operation
	_, err := manager.GetTimeline()

	//Validation
	validateExpectedError(t, err, "No user logged in")
}

func TestCantRetrieveTimelineOfUnregisteredUser(t *testing.T) {

	//Initialization
	var manager service.Manager
	manager.InitializeService()

	user := domain.NewUser("Manuel")

	//Operation
	_, err := manager.GetTimelineFromUser(user)

	//Validation
	validateExpectedError(t, err, "That user is not registered")
}

func TestCanRetrieveTweetById(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()

	var tweet *domain.Tweet
	user := domain.NewUser("root")
	manager.Register(user)
	manager.Login(user)

	text := "This is my first tweet"

	tweet, _ = domain.NewTweet(user, text)
	//Operations
	manager.PublishTweet(tweet)

	//Validation
	publishedTweet, err := manager.GetTweetByID(0)
	if err != nil {
		t.Errorf("Did not expect error, but got %s", err.Error())
	}
	isValidTweet(t, *publishedTweet, user, text)
}

func TestCantRetrieveTweetByNonExistentID(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()

	var tweet *domain.Tweet
	user := domain.NewUser("root")
	manager.Register(user)
	manager.Login(user)

	text := "This is my first tweet"

	tweet, _ = domain.NewTweet(user, text)
	//Operations
	err := manager.PublishTweet(tweet)
	_, err = manager.GetTweetByID(5)

	validateExpectedError(t, err, "A tweet with that ID does not exist")
}

func TestCantCreateTweetWithMoreThan140Characters(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()

	user := domain.NewUser("root")
	manager.Register(user)
	manager.Login(user)
	text := "Este es un texto muy largo que se supone" +
		"que haga fallar al test del tweet, ya que en el" +
		"tweeter que estamos haciendo no se puede tweetear" +
		"algo que tenga mas de 140 caracteres."

	//Operation
	_, err := domain.NewTweet(user, text)

	//Validation
	validateExpectedError(t, err, "Can't have more than 140 characters")
}
