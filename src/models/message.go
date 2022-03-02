package models

import (
	"time"
	//"cloud.google.com/go/firestore"
)

type Message struct {
	ID        string `firestore:"id"`
	CreatedAt time.Time
	Sender    string `firestore:"sender"`
	Content   []byte `firestore:"content"`
}
