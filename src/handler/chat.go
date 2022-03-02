package handler

import (
	"chatdemo/src/firebase"
	"errors"
	"fmt"
)

// userid + userid =>
// tạo cuộc trò chuyện
// userid + userid => lấy ra id cuộc trò chuyện
// id cuộc trò chuyện => lấy ra cuộc trò chuyện

type Chat struct {
	Users []string `firestore:"users"`
}

func isChatExist(userstID string, userndID string) (string, error) {
	// nếu đoạn chat tồn tại (id) => sử dụng
	// nếu chưa tồn tại => tạo + sử dụng
	userSlicest := append(make([]string, 0), userstID, userndID)
	//userSlicend := append(make([]string, 0), userndID)

	fmt.Println(userstID)
	fmt.Println(userndID)
	chatCol := firebase.FirebaseApp.Db.Collection("chats")
	query := chatCol.Where("users", "in", userSlicest)
	chat, err := query.Documents(firebase.Ctx).Next()
	if err != nil {
		fmt.Println("get chat id err: ", err)
		chatid, err := NewChat(userstID, userndID)
		if err != nil {
			return "", errors.New("chat doesnot exist! Cannot create new chat")
		}
		return chatid, nil
	}
	return chat.Ref.ID, nil
}

func NewChat(userstID string, userndID string) (string, error) {
	userSlice := append(make([]string, 0), userstID, userndID)
	chat := Chat{Users: userSlice}
	fmt.Println(chat)

	chatCol := firebase.FirebaseApp.Db.Collection("chats")
	newChat := chatCol.NewDoc()
	wr, err := newChat.Create(firebase.Ctx, chat)
	if err != nil {
		fmt.Println("newChat err: ", err)
		return "", err
	}
	fmt.Println(wr)
	return newChat.ID, nil
}
