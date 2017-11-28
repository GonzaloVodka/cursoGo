package main

import (
	"github.com/abiosoft/ishell"
	"github.com/cursoGo/src/domain"
	"github.com/cursoGo/src/service"
)

func main() {

	shell := ishell.New()
	var manager service.Manager
	manager.InitializeService()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "register",
		Help: "Registers a new user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Pick a name: ")
			user := domain.NewUser(c.ReadLine())
			err := manager.Register(user)
			if err != nil {
				c.Printf("Invalid name, %s", err.Error())
				return
			}
			if manager.IsRegistered(user) {
				c.Print("Added successfully")
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet, err := domain.NewTweet(manager.GetLoggedInUser(), text)

			if err != nil {
				c.Printf("Tweet not published, %s", err.Error())
				return
			}

			err = manager.PublishTweet(tweet)

			if err != nil {
				c.Printf("Tweet not published, %s", err.Error())
			} else {
				c.Print("Tweet sent\n")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "timeline",
		Help: "Shows timeline from logged in user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets, err := manager.GetTimeline()
			if err != nil {
				c.Printf("Can't retrieve timeline, %s", err.Error())
				return
			}
			for _, t := range tweets {
				c.Println(t.StringTweet())
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "Log a user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Enter your name: ")
			user := domain.NewUser(c.ReadLine())

			err := manager.Login(user)

			if err != nil {
				c.Print(err.Error())
				return
			}
			c.Print("Logged")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "logout",
		Help: "Logout a user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			manager.Logout()
		},
	})

	shell.Run()

}
