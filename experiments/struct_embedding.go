package main

import (
	"fmt"
	"time"
)

type UserData struct {
	SenderId string `json:"sender_id"`
	Sender   string `json:"sender"`
}

type Message struct {
	*UserData
	Message string `json:"message"`
	Time    int64  `json:"time"`
}

type Event struct {
	Type string
	*Message
}

func main() {
	var event = &Event{Type: "my_type", Message: &Message{
		Message: "Hello  world",
		Time:    time.Now().UnixMilli(),
		UserData: &UserData{
			SenderId: "123",
			Sender:   "Bob",
		},
	}}

	fmt.Println(event)
	fmt.Println(event.Time)
}
