package models

import "time"

type Message struct {
	Text      string    `bson:"text" json:"text"`
	Status    string    `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

type SMS struct {
	PhoneNumber string    `bson:"_id" json:"phoneNumber"`
	Messages    []Message `bson:"messages" json:"messages"`
}

type IncomingSMS struct {
	PhoneNumber string `json:"phoneNumber"`
	Message     string `json:"message"`
	Status      string `json:"status"`
}
