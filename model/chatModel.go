package models

import (
	"mental-health-api/pkg/const/firestoreCol"
	"mental-health-api/pkg/firebase"

	"fmt"

	"cloud.google.com/go/firestore"
)

type Chat struct {
	ID    string   `firestore:"id"`
	Users []string `firestore:"users"`
}

func FindChatID(chatIDst, chatIDnd string, chatCol *firestore.CollectionRef) (string, error) {
	fmt.Println("chatID ", chatIDst)
	query := chatCol.Where("id", "==", chatIDst)
	chat, err := query.Documents(firebase.Ctx).Next()
	if err != nil {
		query = chatCol.Where("id", "==", chatIDnd)
		chat, err = query.Documents(firebase.Ctx).Next()
		if err != nil {
			return "", err
		}
	}
	return chat.Ref.ID, nil
}

func NewChat(userstID string, userndID string, chatID string, chatCol *firestore.CollectionRef) (string, error) {
	usersID := append(make([]string, 0), userstID, userndID)
	chat := Chat{
		ID:    chatID,
		Users: usersID,
	}

	newChat := chatCol.Doc(chatID)
	_, err := newChat.Set(firebase.Ctx, chat)
	if err != nil {
		return "", err
	}
	return newChat.ID, nil
}

func GetChatID(userstID string, userndID string) (string, error) {
	chatIDst := userstID + userndID
	chatIDnd := userndID + userstID
	chatCol := firebase.FirebaseApp.Db.Collection(firestoreCol.CHAT_COLLECTION)
	chatID, err := FindChatID(chatIDst, chatIDnd, chatCol)
	if err != nil {
		chatID, err = NewChat(userstID, userndID, chatIDst, chatCol)
		if err != nil {
			return "", err
		}
	}
	return chatID, nil
}
