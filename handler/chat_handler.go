package handler

import (
	models "mental-health-api/model"
	"mental-health-api/pkg/const/firestoreCol"
	"mental-health-api/pkg/firebase"

	"github.com/gofiber/fiber/v2"
)

type Messages struct {
	Message []models.Message
}

func GetAllMessages(ctx *fiber.Ctx) error {
	var messagesResponse Messages

	senderID := ctx.Params("userid")
	receiverID := ctx.Params("id")
	chatID, err := models.GetChatID(senderID, receiverID)
	if err != nil {
		return err
	}
	chatCol := firebase.FirebaseApp.Db.Collection(firestoreCol.CHAT_COLLECTION)
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
	return ctx.Status(fiber.StatusCreated).JSON(messagesResponse)
}
