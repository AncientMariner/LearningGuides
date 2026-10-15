package behavioral

import "fmt"

// Mediator: centralises communication between components.
// Channels are often simpler for goroutine-based scenarios.

type ChatMediator interface {
	Send(message, from string)
	Register(user *ChatUser)
}

type ChatRoom struct {
	users []*ChatUser
}

func (c *ChatRoom) Register(user *ChatUser) {
	c.users = append(c.users, user)
	user.room = c
}

func (c *ChatRoom) Send(message, from string) {
	for _, u := range c.users {
		if u.name != from {
			u.Receive(message, from)
		}
	}
}

type ChatUser struct {
	name string
	room *ChatRoom
}

func NewChatUser(name string) *ChatUser { return &ChatUser{name: name} }

func (u *ChatUser) Send(message string) {
	fmt.Printf("%s sends: %s\n", u.name, message)
	u.room.Send(message, u.name)
}

func (u *ChatUser) Receive(message, from string) {
	fmt.Printf("%s receives from %s: %s\n", u.name, from, message)
}

func ExampleMediator() {
	room := &ChatRoom{}
	alice := NewChatUser("Alice")
	bob := NewChatUser("Bob")
	room.Register(alice)
	room.Register(bob)
	alice.Send("hi Bob")
	// Output:
	// Alice sends: hi Bob
	// Bob receives from Alice: hi Bob
}
