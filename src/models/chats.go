package models

import (
	"cloud.google.com/go/firestore"
)

type Chat struct {
	ID    string                   `firestore:"id"`
	Users *[]firestore.DocumentRef `firestore:"users"`
}

//*firestore.DocumentRef converts to Reference
