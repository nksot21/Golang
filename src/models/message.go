package models

import (
	"time"
	//"cloud.google.com/go/firestore"
)

type Message struct {
	CreatedAt time.Time
	Sender    string `firestore:"sender"`
	Content   string `firestore:"content"`
}
