package main

import (
	"fmt"
	"time"
)

type Client struct {
	ID                     int
	Username               string
	IncomingMessageChannel chan string
	Listening              bool
}

type Chatroom interface {
	Join(client *Client)
	Leave(client *Client)
	Broadcast(client *Client, message string)
}

type ChatroomImpl struct {
	clients map[int]*Client
}

func (c *ChatroomImpl) Join(client *Client) {
	c.clients[client.ID] = client
}

func (c *ChatroomImpl) Leave(client *Client) {
	delete(c.clients, client.ID)
}

func (c *ChatroomImpl) Broadcast(client *Client, message string) {
	for _, cl := range c.clients {
		if cl.ID != client.ID {
			cl.IncomingMessageChannel <- message
		}
	}
}

func (c *Client) CloseMessageChannel() {
	close(c.IncomingMessageChannel)
}

func (c *Client) PrintMessages() {
	if c.Listening {
		return
	}

	go func(messageChannel <-chan string) {
		// Note: this only works because we never write to "Listening" multiple times.
		// otherwise we would have to use a mutex!
		c.Listening = true
		for mesage := range messageChannel {
			fmt.Printf("%s got message: %s\n", c.Username, mesage)
		}
		c.Listening = false
	}(c.IncomingMessageChannel)
}

func NewClient(id int, username string) *Client {
	return &Client{
		ID:                     id,
		Username:               username,
		IncomingMessageChannel: make(chan string),
	}
}

func JoinChatroom(chatroom Chatroom, client *Client) {
	chatroom.Join(client)
}

func LeaveChatroom(chatroom Chatroom, client *Client) {
	client.CloseMessageChannel()
	chatroom.Leave(client)
}

func SendMessage(chatroom Chatroom, client *Client, message string) {
	chatroom.Broadcast(client, message)
}

func main() {
	chatroom := &ChatroomImpl{clients: make(map[int]*Client)}

	client1 := NewClient(1, "Alice")
	client2 := NewClient(2, "Bob")
	JoinChatroom(chatroom, client1)
	JoinChatroom(chatroom, client2)

	client1.PrintMessages()
	client2.PrintMessages()

	SendMessage(chatroom, client1, "Hello, Bob!")
	SendMessage(chatroom, client2, "Hi, Alice!")

	LeaveChatroom(chatroom, client1)
	LeaveChatroom(chatroom, client2)

	for client1.Listening || client2.Listening {
		fmt.Println("All clients are still listening, waiting for them to leave...")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("All leaved. Closing")
}
