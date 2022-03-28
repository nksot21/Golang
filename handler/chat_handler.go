package handler

import (
	models "mental-health-api/model"
	"mental-health-api/pkg/const/firestoreCol"
	"mental-health-api/pkg/firebase"

	"time"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"

	"fmt"
)

type messageResponse struct {
	ID        string
	CreatedAt time.Time
	SenderID  string
	Sender    models.User
	Content   string
}

type MessagesResponse struct {
	Message []messageResponse
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
	var messagesResponse MessagesResponse

	senderID := ctx.Params("userid")
	receiverID := ctx.Params("id")
	chatID, err := models.GetChatID(senderID, receiverID)
	if err != nil {
		return err
	}
	chatCol := firebase.FirebaseApp.Db.Collection(firestoreCol.CHAT_COLLECTION)
	chat := chatCol.Doc(chatID)
	messagesDocIter := chat.Collection(firestoreCol.MESSAGE_COLLECTION).OrderBy("CreatedAt", firestore.Asc).Documents(firebase.Ctx)
	messages, err := messagesDocIter.GetAll()
	if err != nil {
		return err
	}
	for messgIndex := range messages {
		if err != nil {
			return err
		}
		var message models.Message
		if err = messages[messgIndex].DataTo(&message); err != nil {
			return err
		}

		//get sender-info
		var sender models.User
		if err = sender.GetOne(message.Sender, ""); err != nil {
			fmt.Println("Get_user_id: ", err)
		}

		if sender.Picture == "" {
			sender.Picture = firestoreCol.DEFAULT_PICTURE
		}

		messageResponse := messageResponse{ID: messages[messgIndex].Ref.ID, CreatedAt: message.CreatedAt, SenderID: message.Sender, Sender: sender, Content: message.Content}
		messagesResponse.Message = append(messagesResponse.Message, messageResponse)
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

	fmt.Println("chatIDs: ", chatIDs)
	return c.Status(200).JSON(conversationsInfo)
}

// Change current showEmotion status
// @Summary Change current showEmotion status
// @Tags /chat
// @Accept json
// @Produce json
// @Param userid path string true "UserID"
// @Param id path string true "ID"
// @Success 200 ""
// @Router /chat/emotion/{userid}/{id} [put]
func ShowEmotion(ctx *fiber.Ctx) error {
	senderID := ctx.Params("userid")
	receiverID := ctx.Params("id")
	chatID, err := models.GetChatID(senderID, receiverID)
	if err != nil {
		return err
	}

	chatCol := firebase.FirebaseApp.Db.Collection(firestoreCol.CHAT_COLLECTION)
	chat := chatCol.Doc(chatID)
	chatSnap, err := chat.Get(firebase.Ctx)
	if err != nil {
		return err
	}
	var chatInfo models.Chat
	err = chatSnap.DataTo(&chatInfo)
	if err != nil {
		return err
	}

	showEmotionStatus := chatInfo.ShowEmotion

	fmt.Println("show emotion status: ", showEmotionStatus)
	chat.Update(firebase.Ctx, []firestore.Update{
		{Path: firestoreCol.SHOW_EMOTION, Value: !showEmotionStatus},
	})

	return ctx.Status(200).JSON(!showEmotionStatus)
}
