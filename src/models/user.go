package models

import "time"

type User struct {
	ID        string `firestore:"id"`
	CreatedAt time.Time
	Name      string `firestore:"name"`
	Email     string `firestore:"email"`
	Password  string `firestore:"password"`
}
