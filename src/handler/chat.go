package handler

import (
	"chatdemo/src/firebase"
	"chatdemo/src/models"

	"cloud.google.com/go/firestore"
)

func findChatID(chatIDst, chatIDnd string, chatCol *firestore.CollectionRef) (string, error) {
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
	userst, err := GetUserByID(userstID)
	if err != nil {
		return "", err
	}
	usernd, err := GetUserByID(userndID)
	if err != nil {
		return "", err
	}
	usersRef := append(make([]firestore.DocumentRef, 0), *userst, *usernd)
	chat := models.Chat{
		ID:    chatID,
		Users: &usersRef,
	}

	newChat := chatCol.Doc(chatID)
	_, err = newChat.Set(firebase.Ctx, chat)
	if err != nil {
		return "", err
	}
	return newChat.ID, nil
}

func getChatID(userstID string, userndID string) (string, error) {
	chatIDst := userstID + userndID
	chatIDnd := userndID + userstID
	chatCol := firebase.FirebaseApp.Db.Collection("chats")
	chatID, err := findChatID(chatIDst, chatIDnd, chatCol)
	if err != nil {
		chatID, err = NewChat(userstID, userndID, chatIDst, chatCol)
		if err != nil {
			return "", err
		}
	}
	return chatID, nil
}
