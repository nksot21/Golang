package handler

import (
	models "mental-health-api/model"
	"mental-health-api/pkg/const/firestoreCol"
	"mental-health-api/pkg/firebase"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"

	"fmt"
)

type Messages struct {
	Message []models.Message
}

// Get All Messages
// @Summary Get All Messages
// @Tags /chat
// @Accept json
// @Produce json
// @Param userid path string true "UserID"
// @Param id path string true "ID"
// @Success 200 ""
// @Router /chat/getall/{userid}/{id} [get]
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
	messagesDocIter := chat.Collection(firestoreCol.MESSAGE_COLLECTION).OrderBy("CreatedAt", firestore.Asc).Documents(firebase.Ctx)
	//messagesRef := chat.Collection("messages").DocumentRefs(firebase.Ctx)
	messages, err := messagesDocIter.GetAll()
	if err != nil {
		return err
	}
	for messgIndex := range messages {
		//messageSnap, err := messages[messgIndex].Get(firebase.Ctx)
		if err != nil {
			return err
		}
		var message models.Message
		//if err = messageSnap.DataTo(&message); err != nil {
		//	return err
		//}
		if err = messages[messgIndex].DataTo(&message); err != nil {
			return err
		}
		messagesResponse.Message = append(messagesResponse.Message, message)
	}

	return ctx.Status(fiber.StatusCreated).JSON(messagesResponse)
}

// Connect Chat
// @Summary Update to websocket
// @Tags /chat
// @Accept json
// @Produce json
// @Param userid path string true "UserID"
// @Success 200 ""
// @Router /chat/{userid} [get]
func ChatPage(c *fiber.Ctx) error {
	//models.GetAllMessages(c.Params("userid"), c.Params("id"))
	return c.SendFile("home.html")
}

// Get Conversations' summary
// @Summary Get conversations' summary
// @Tags /chat
// @Accept json
// @Produce json
// @Param userid path string true "UserID"
// @Success 200 {object} models.ChatSummary
// @Router /chat/conversations/{userid} [get]
func GetChatIDs(c *fiber.Ctx) error {
	var chatIDs []string
	//var chats []ChatShortCut

	userID := c.Params("userid")
	chatCol := firebase.FirebaseApp.Db.Collection(firestoreCol.CHAT_COLLECTION)
	chatDocumentIter := chatCol.Where("users", "array-contains", userID).Documents(firebase.Ctx)
	chatsSnap, err := chatDocumentIter.GetAll()
	if err != nil {
		return err
	}

	conversationsInfo, err := models.ConversationsInfo(chatsSnap, userID)
	if err != nil {
		return err
	}

	for chatIndex := range chatsSnap {
		ID := chatsSnap[chatIndex].Ref.ID
		chatIDs = append(chatIDs, ID)
	}

	fmt.Println("text: ", conversationsInfo)
	fmt.Println("chatIDs: ", chatIDs)
	return c.Status(200).JSON(conversationsInfo)
}
