package main

import (
	"github.com/abiosoft/ishell"
	"github.com/cursoGo/src/domain"
	"github.com/cursoGo/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Who are you? ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet := domain.NewTweet(user, text)

			err := service.PublishTweet(tweet)

			if err != nil {
				c.Printf("Tweet not published, %s", err.Error())
			} else {
				c.Print("Tweet sent\n")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Println(domain.StringTweet(tweet))

			return
		},
	})

	shell.Run()

}
