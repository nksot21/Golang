package models

type Chat struct {
	ID    string   `firestore:"id"`
	Users []string `firestore:"users"`
}
