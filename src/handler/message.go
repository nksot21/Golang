package handler

// id cuộc trò chuyện => thêm message
import (
	"chatdemo/src/firebase"
	"chatdemo/src/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
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
		fmt.Println("add message error: ", err.Error())
		return "", err
	}
	chat := chatCol.Doc(chatid)
	messgCol := chat.Collection("messages")
	newMessage := messgCol.NewDoc()
	_, err = newMessage.Create(firebase.Ctx, newMessg)
	if err != nil {
		fmt.Println("add message error 2: ", err)
		return "", err
	}
	return newMessage.ID, nil
}

func GetAllMessage(c *fiber.Ctx) error {
	var messagesResponse Messages
	senderID := c.Params("userid")
	receiverID := c.Params("id")
	chatID, err := getChatID(senderID, receiverID)
	if err != nil {
		fmt.Println("cannot get chat id")
		return err
	}
	chatCol := firebase.FirebaseApp.Db.Collection("chats")
	chat := chatCol.Doc(chatID)
	fmt.Println("chatid: ", chat.ID)
	messagesRef := chat.Collection("messages").DocumentRefs(firebase.Ctx)
	messages, err := messagesRef.GetAll()
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
	return c.Status(200).JSON(messagesResponse)
}
