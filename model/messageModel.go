package models

import (
	"mental-health-api/pkg/const/firestoreCol"
	"mental-health-api/pkg/firebase"
	"time"
)

type Message struct {
	CreatedAt time.Time
	Sender    string `firestore:"sender"`
	Content   string `firestore:"content"`
}

func NewMessage(receiverID, senderID string, content []byte) (string, error) {
	contentStr := string(content)
	newMessg := Message{CreatedAt: time.Now(), Sender: senderID, Content: contentStr}
	chatCol := firebase.FirebaseApp.Db.Collection(firestoreCol.CHAT_COLLECTION)
	chatid, err := GetChatID(senderID, receiverID)
	if err != nil {
		return "", err
	}
	chat := chatCol.Doc(chatid)
	messgCol := chat.Collection("messages")
	newMessage := messgCol.NewDoc()
	_, err = newMessage.Create(firebase.Ctx, newMessg)
	if err != nil {
		return "", err
	}
	return newMessage.ID, nil
}
