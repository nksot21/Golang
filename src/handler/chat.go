package handler

import (
	"chatdemo/src/firebase"
	"chatdemo/src/models"
	"errors"
	"fmt"

	"cloud.google.com/go/firestore"
)

// userid + userid =>
// tạo cuộc trò chuyện
// userid + userid => lấy ra id cuộc trò chuyện
// id cuộc trò chuyện => lấy ra cuộc trò chuyện

func findChatID(chatIDst, chatIDnd string, chatCol *firestore.CollectionRef) (string, error) {
	query := chatCol.Where("id", "==", chatIDst)
	chat, err := query.Documents(firebase.Ctx).Next()
	if err != nil {
		query = chatCol.Where("id", "==", chatIDnd)
		chat, err = query.Documents(firebase.Ctx).Next()
		if err != nil {
			fmt.Println("find err: ", err)
			return "", err
		}
	}
	return chat.Ref.ID, nil
}

func getChatID(userstID string, userndID string) (string, error) {
	// nếu đoạn chat tồn tại (id) => sử dụng
	// nếu chưa tồn tại => tạo + sử dụng
	//userSlicend := append(make([]string, 0), userndID)

	chatIDst := userstID + userndID
	chatIDnd := userndID + userstID
	chatCol := firebase.FirebaseApp.Db.Collection("chats")
	chatID, err := findChatID(chatIDst, chatIDnd, chatCol)
	if err != nil {
		chatID, err = NewChat(userstID, userndID, chatIDst, chatCol)
		if err != nil {
			return "", errors.New("cannot create new chat")
		}
	}
	return chatID, nil
}

func NewChat(userstID string, userndID string, chatID string, chatCol *firestore.CollectionRef) (string, error) {
	userSlice := append(make([]string, 0), userstID, userndID)
	chat := models.Chat{
		ID:    chatID,
		Users: userSlice,
	}
	fmt.Println("chatID:", chatID)

	newChat := chatCol.Doc(chatID)
	wr, err := newChat.Set(firebase.Ctx, chat)
	if err != nil {
		fmt.Println("newChat err: ", err)
		return "", err
	}
	fmt.Println(wr)
	return newChat.ID, nil
}
