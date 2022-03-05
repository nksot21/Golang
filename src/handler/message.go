package handler

import (
	"chatdemo/src/firebase"
	"chatdemo/src/models"
	"time"
)

type Messages struct {
	Message []models.Message
}

func NewMessage(receiverID, senderID string, content []byte) (string, error) {
	contentStr := string(content)
	newMessg := models.Message{CreatedAt: time.Now(), Sender: senderID, Content: contentStr}
	chatCol := firebase.FirebaseApp.Db.Collection("chats")
	chatid, err := getChatID(senderID, receiverID)
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

func GetAllMessages(senderID, receiverID string) error {
	var messagesResponse Messages
	chatID, err := getChatID(senderID, receiverID)
	if err != nil {
		return err
	}
	chatCol := firebase.FirebaseApp.Db.Collection("chats")
	chat := chatCol.Doc(chatID)
	messagesRef := chat.Collection("messages").DocumentRefs(firebase.Ctx)
	messages, err := messagesRef.GetAll()
	if err != nil {
		return err
	}
	for messgIndex := range messages {
		messageSnap, err := messages[messgIndex].Get(firebase.Ctx)
		if err != nil {
			return err
		}
		var message models.Message
		if err = messageSnap.DataTo(&message); err != nil {
			return err
		}
		messagesResponse.Message = append(messagesResponse.Message, message)
	}
	return nil
}
