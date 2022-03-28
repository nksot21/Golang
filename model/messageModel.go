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
	//authenticate reeceiverID, senderID
	var receiver User
	var sender User
	if err := receiver.GetOne(receiverID, ""); err != nil {
		return "", err
	}

	if err := sender.GetOne(senderID, ""); err != nil {
		return "", err
	}

	contentStr := string(content)
	chatCol := firebase.FirebaseApp.Db.Collection(firestoreCol.CHAT_COLLECTION)
	chatid, err := GetChatID(senderID, receiverID)
	if err != nil {
		return "", err
	}
	chat := chatCol.Doc(chatid)
	messgCol := chat.Collection(firestoreCol.MESSAGE_COLLECTION)
	messageRef := messgCol.NewDoc()
	newMessage := Message{CreatedAt: time.Now(), Sender: senderID, Content: contentStr}
	_, err = messageRef.Create(firebase.Ctx, newMessage)
	if err != nil {
		return "", err
	}
	return messageRef.ID, nil
}
