package models

type Chat struct {
	ID     uint `firestore:"id"`
	Userst int  `firestore:"userst"`
	Usernd int  `firestore:"usernd"`
}
