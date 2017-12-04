package domain

//User of tweeter
type User struct {
	Name      string
	Nick      string
	Mail      string
	Pass      string
	Followers []User
	Following []User
}

//NewUser Creates a new user
func NewUser(name string, pass string, nick string, mail string) User {
	return User{Name: name, Pass: pass, Nick: nick, Mail: mail, Followers: make([]User, 0), Following: make([]User, 0)}
}
