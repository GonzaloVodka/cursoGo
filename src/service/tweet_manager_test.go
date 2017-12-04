package service_test

import (
	"testing"

	"github.com/cursoGo/src/domain"
	"github.com/cursoGo/src/service"
)

func isValidTweet(t *testing.T, publishedTweet domain.Tweet, user domain.User, text string) bool {
	if publishedTweet.GetUser().Name != user.Name && publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user.Name, text, publishedTweet.GetUser().Name, publishedTweet.GetText())
		return false
	}

	if publishedTweet.GetDate() == nil {
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

	user := domain.NewUser("root", "root", "root", "root")
	manager.Register(user)
	manager.Login(user)

	//Operation
	err := manager.Login(user)

	//Validation
	validateExpectedError(t, err, "Already logged in")
}

// func TestCantLogInWithUnregisteredUser(t *testing.T) {
// 	//Initialization
// 	var manager service.Manager

// 	manager.InitializeService()
// 	user := domain.NewUser("root", "", "", "")

// 	//Operation
// 	err := manager.Login(user)

// 	//Validation
// 	validateExpectedError(t, err, "The user is not registered")

// }

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

	var tweet domain.Tweet
	user := domain.NewUser("root", "", "", "")
	manager.Register(user)

	text := "This is my first tweet"
	tweet, _ = domain.NewTextTweet(user, text)
	//Operation
	err := manager.PublishTweet(tweet)
	validateExpectedError(t, err, "You must be logged in to tweet")

}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	//Initialization
	var manager service.Manager
	manager.InitializeService()

	var tweet domain.Tweet

	user := domain.NewUser("Gonzalo", "", "", "")
	manager.Register(user)
	manager.Login(user)
	var text string

	tweet, _ = domain.NewTextTweet(user, text)

	//Operation
	err := manager.PublishTweet(tweet)

	//Validation
	validateExpectedError(t, err, "Text is required")
}

// func TestCanPublishAndRetriveMoreThanOneTweet(t *testing.T) {

// 	//Initialization
// 	var manager service.Manager
// 	manager.InitializeService()

// 	var tweet, secondTweet domain.Tweet

// 	user := domain.NewUser("Manuel", "", "", "")
// 	manager.Register(user)
// 	manager.Login(user)
// 	text := "This is my first tweet"
// 	secondText := "This is my second tweet"

// 	tweet, _ = domain.NewTextTweet(user, text)
// 	secondTweet, _ = domain.NewTextTweet(user, secondText)

// 	//Operation
// 	manager.PublishTweet(tweet)
// 	manager.PublishTweet(secondTweet)

// 	//Validation
// 	publishedTweets := manager.GetTweets()

// 	if len(publishedTweets) != 2 {
// 		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
// 		return
// 	}
// 	firstPublishedTweet := publishedTweets[0]
// 	secondPublishedTweet := publishedTweets[1]

// 	if !isValidTweet(t, firstPublishedTweet, user, text) {
// 		return
// 	}
// 	isValidTweet(t, secondPublishedTweet, user, secondText)
// }

func TestCanRegisterUser(t *testing.T) {

	//Initialization
	var manager service.Manager
	manager.InitializeService()
	user := domain.NewUser("Gonza", "", "", "")
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
	user := domain.NewUser("Gonza", "", "", "")
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

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := domain.NewUser("Manuel", "", "Manuel", "")
	manager.Register(user)

	secondUser := domain.NewUser("Gonzalo", "", "Gonzalo", "")
	manager.Register(secondUser)

	text := "This is my first tweet"
	secondText := "This is my second tweet"
	thirdText := "This is a tweet"

	tweet, _ = domain.NewTextTweet(user, text)
	secondTweet, _ = domain.NewTextTweet(user, secondText)
	thirdTweet, _ = domain.NewTextTweet(secondUser, thirdText)

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
		if tweet.GetUser().Name != user.Name {
			t.Errorf("Expected user is %s but was %s", user.Name, tweet.GetUser().Name)
		}
	}
}

func TestCantLoginWithAnIncorrectPassword(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()

	user := domain.NewUser("Gonzalo", "pass", "Gonzalo", "")

	user2 := domain.NewUser("Gonzalo", "incorrectPass", "Gonzalo", "")
	//Operation
	manager.Register(user)

	err := manager.Login(user2)

	//Validation
	validateExpectedError(t, err, "The password is incorrect")
}

func TestCantRetrieveTimelineWithoutLoggingIn(t *testing.T) {

	//Initialization
	var manager service.Manager
	manager.InitializeService()

	var tweet domain.Tweet

	user := domain.NewUser("Manuel", "", "", "")
	manager.Register(user)
	manager.Login(user)

	text := "This is my first tweet"
	tweet, _ = domain.NewTextTweet(user, text)

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

	user := domain.NewUser("Manuel", "", "", "")

	//Operation
	_, err := manager.GetTimelineFromUser(user)

	//Validation
	validateExpectedError(t, err, "That user is not registered")
}

// func TestCanRetrieveTweetById(t *testing.T) {
// 	//Initialization
// 	var manager service.Manager
// 	manager.InitializeService()

// 	var tweet domain.Tweet
// 	user := domain.NewUser("root", "", "", "")
// 	manager.Register(user)
// 	manager.Login(user)

// 	text := "This is my first tweet"

// 	tweet, _ = domain.NewTextTweet(user, text)
// 	//Operations
// 	manager.PublishTweet(tweet)

// 	//Validation
// 	publishedTweet, err := manager.GetTweetByID(0)
// 	if err != nil {
// 		t.Errorf("Did not expect error, but got %s", err.Error())
// 	}
// 	isValidTweet(t, *publishedTweet, user, text)
// }

// func TestCantRetrieveTweetByNonExistentID(t *testing.T) {
// 	//Initialization
// 	var manager service.Manager
// 	manager.InitializeService()

// 	var tweet *domain.Tweet
// 	user := domain.NewUser("root")
// 	manager.Register(user)
// 	manager.Login(user)

// 	text := "This is my first tweet"

// 	tweet, _ = domain.NewTweet(user, text)
// 	//Operations
// 	err := manager.PublishTweet(tweet)
// 	_, err = manager.GetTweetByID(5)

// 	validateExpectedError(t, err, "A tweet with that ID does not exist")
// }

func TestCantCreateTweetWithMoreThan140Characters(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()

	user := domain.NewUser("root", "", "", "")
	manager.Register(user)
	manager.Login(user)
	text := "Este es un texto muy largo que se supone" +
		"que haga fallar al test del tweet, ya que en el" +
		"tweeter que estamos haciendo no se puede tweetear" +
		"algo que tenga mas de 140 caracteres."

	//Operation
	_, err := domain.NewTextTweet(user, text)

	//Validation
	validateExpectedError(t, err, "Can't have more than 140 characters")
}

// TEST QUE NOS DIO SANTI

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	//tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	user := domain.NewUser("grupoesfera", "grupoesfera", "grupoesfera", "") //agregado
	tweet, _ := domain.NewTextTweet(user, "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	user := domain.NewUser("grupoesfera", "grupoesfera", "grupoesfera", "")
	tweet, _ := domain.NewImageTweet(user, "This is my image", "http://www.grupoesfera.com.ar/common/img/grupoesfera.png")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	// Initialization
	user := domain.NewUser("grupoesfera", "", "grupoesfera", "")
	quotedTweet, _ := domain.NewTextTweet(user, "This is my tweet")
	user2 := domain.NewUser("nick", "", "nick", "")
	tweet, _ := domain.NewQuoteTweet(user2, "Awesome", quotedTweet)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	user := domain.NewUser("grupoesfera", "", "grupoesfera", "")
	tweet, _ := domain.NewTextTweet(user, "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanDeleteATweet(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()
	user := domain.NewUser("Gonzalo", "", "Gonzalo", "")
	manager.Register(user)
	manager.Login(user)
	tweet, _ := domain.NewTextTweet(user, "This is my first tweet")
	manager.PublishTweet(tweet)
	tweet2, _ := domain.NewTextTweet(user, "This is my second tweet")
	manager.PublishTweet(tweet2)
	tweet3, _ := domain.NewTextTweet(user, "This is my third tweet")
	manager.PublishTweet(tweet3)

	//Operation
	s, _ := manager.DeleteTweet(2)

	//Validation

	if s != "A tweet was deleted" {
		t.Errorf("Expected A tweet was deleted but was %s", s)
	}

}

func TestCantDeleteATweetUnlogged(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()
	user := domain.NewUser("Gonzalo", "", "Gonzalo", "")
	manager.Register(user)
	tweet, _ := domain.NewTextTweet(user, "This is my tweet")
	manager.PublishTweet(tweet)

	//Operation
	_, err := manager.DeleteTweet(0)

	//Validation

	validateExpectedError(t, err, "No user logged in")
}

func TestCantDeleteATweetWithWrongId(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()
	user := domain.NewUser("Gonzalo", "", "Gonzalo", "")
	manager.Register(user)
	manager.Login(user)
	tweet, _ := domain.NewTextTweet(user, "This is my first tweet")
	manager.PublishTweet(tweet)

	//Operation
	s, _ := manager.DeleteTweet(1)

	//Validation

	if s != "No tweet deleted" {
		t.Errorf("Expected No tweet deleted but was %s", s)
	}
}

func TestCantDeleteATweetOfAnotherUser(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()

	//First user publish a tweet
	user := domain.NewUser("Gonzalo", "Gonzalo", "Gonzalo", "Gonzalo")
	manager.Register(user)
	manager.Login(user)
	tweet, _ := domain.NewTextTweet(user, "This is my first tweet")
	manager.PublishTweet(tweet)
	manager.Logout()

	//Second user publish a tweet
	user2 := domain.NewUser("Another", "Another", "Another", "Another")
	manager.Register(user2)
	manager.Login(user2)
	tweet2, _ := domain.NewTextTweet(user2, "This is my first tweet")
	manager.PublishTweet(tweet2)
	manager.Logout()

	//Operation
	manager.Login(user)
	s, _ := manager.DeleteTweet(1)

	//Validation

	if s != "No tweet deleted" {
		t.Errorf("Expected No tweet deleted but was %s", s)
	}
}

func TestCantPublishADuplicateTextTweet(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()

	user := domain.NewUser("Gonzalo", "Gonzalo", "Gonzalo", "Gonzalo")
	manager.Register(user)
	manager.Login(user)
	tweet, _ := domain.NewTextTweet(user, "This is my first tweet")
	manager.PublishTweet(tweet)

	//Operation

	tweet2, _ := domain.NewTextTweet(user, "This is my first tweet")
	err := manager.PublishTweet(tweet2)

	//Validation

	validateExpectedError(t, err, "Tweet duplicated")
}
func TestCantPublishADuplicateImageTweet(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()

	user := domain.NewUser("Gonzalo", "Gonzalo", "Gonzalo", "Gonzalo")
	manager.Register(user)
	manager.Login(user)
	tweet, _ := domain.NewImageTweet(user, "This is my first tweet", "url")
	manager.PublishTweet(tweet)

	//Operation

	tweet2, _ := domain.NewImageTweet(user, "This is my first tweet", "url")
	err := manager.PublishTweet(tweet2)

	//Validation

	validateExpectedError(t, err, "Tweet duplicated")
}
func TestCantPublishADuplicateQuoteTweet(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()

	user := domain.NewUser("Gonzalo", "Gonzalo", "Gonzalo", "Gonzalo")
	manager.Register(user)
	manager.Login(user)
	quotedtweet, _ := domain.NewTextTweet(user, "This is my first tweet")
	tweet, _ := domain.NewQuoteTweet(user, "This is my first tweet", quotedtweet)
	manager.PublishTweet(tweet)

	//Operation

	quotedtweet2, _ := domain.NewTextTweet(user, "This is my first tweet")
	tweet2, _ := domain.NewQuoteTweet(user, "This is my first tweet", quotedtweet2)
	err := manager.PublishTweet(tweet2)

	//Validation

	validateExpectedError(t, err, "Tweet duplicated")
}

func TestCanFollowAUser(t *testing.T) {
	//Initialization
	var manager service.Manager
	manager.InitializeService()

	user := domain.NewUser("Gonzalo", "Gonzalo", "Gonzalo", "Gonzalo")
	user2 := domain.NewUser("Another", "Another", "Another", "Another")
	manager.Register(user)
	manager.Register(user2)
	manager.Login(user)

	//Operation

	manager.Follow("Another")

	//Validation
	var b bool

	for _, u := range manager.GetLoggedInUser().Following {
		if u.Nick == user2.Nick {
			b = true
		}
	}

	if !b {
		t.Errorf("Not following")
	}

}
