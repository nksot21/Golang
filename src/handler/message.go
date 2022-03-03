package handler

// id cuộc trò chuyện => thêm message
import (
	"chatdemo/src/firebase"
	"chatdemo/src/models"
	"fmt"
	"time"
)

/*type Message struct {
	senderID  string `firestore:"senderid"`
	createdAt time.Time
	content   []byte `firestore:"content"`
}*/

func NewMessage(receiverID, senderID string, content []byte) (string, error) {
	contentStr := string(content)
	newMessg := models.Message{CreatedAt: time.Now(), Sender: senderID, Content: contentStr}
	fmt.Println(newMessg)
	fmt.Println(receiverID)
	fmt.Println(senderID)
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
