package models

import (
	"mental-health-api/pkg/const/firestoreCol"
	"mental-health-api/pkg/firebase"

	"fmt"

	"cloud.google.com/go/firestore"
)

type Chat struct {
	ID    string   `firestore:"id"`
	Users []string `firestore:"users"`
}

type ChatSummary struct {
	ChatID      string
	Friend      User
	LastMessage string
}

type Conversation struct {
	users []string
}

func FindChatID(chatIDst, chatIDnd string, chatCol *firestore.CollectionRef) (string, error) {
	fmt.Println("chatID ", chatIDst)
	query := chatCol.Where("id", "==", chatIDst)
	chat, err := query.Documents(firebase.Ctx).Next()
	if err != nil {
		query = chatCol.Where("id", "==", chatIDnd)
		chat, err = query.Documents(firebase.Ctx).Next()
		if err != nil {
			return "", err
		}
	}
	return chat.Ref.ID, nil
}

func NewChat(userstID string, userndID string, chatID string, chatCol *firestore.CollectionRef) (string, error) {
	usersID := append(make([]string, 0), userstID, userndID)
	chat := Chat{
		ID:    chatID,
		Users: usersID,
	}

	newChat := chatCol.Doc(chatID)
	_, err := newChat.Set(firebase.Ctx, chat)
	if err != nil {
		return "", err
	}
	return newChat.ID, nil
}

func GetChatID(userstID string, userndID string) (string, error) {
	chatIDst := userstID + userndID
	chatIDnd := userndID + userstID
	chatCol := firebase.FirebaseApp.Db.Collection(firestoreCol.CHAT_COLLECTION)
	chatID, err := FindChatID(chatIDst, chatIDnd, chatCol)
	if err != nil {
		chatID, err = NewChat(userstID, userndID, chatIDst, chatCol)
		if err != nil {
			return "", err
		}
	}
	return chatID, nil
}

//GET CONVERSATIONS' INFO BY USERID
func ConversationsInfo(chatsSnap []*firestore.DocumentSnapshot, userID string) ([]ChatSummary, error) {
	var conversationsInfo []ChatSummary

	for chatIndex := range chatsSnap {
		chatSnap := chatsSnap[chatIndex]
		conversationInfo, err := ConversationInfo(chatSnap, userID)
		if err != nil {
			return conversationsInfo, err
		}
		conversationsInfo = append(conversationsInfo, conversationInfo)
	}

	return conversationsInfo, nil
}

//GET CONVERSATION'S INFO BY USERID (receiverID => userinfo, last message)
func ConversationInfo(chatSnap *firestore.DocumentSnapshot, userID string) (ChatSummary, error) {

	var chatSummary ChatSummary
	var converInfo Chat
	var friendID string
	err := chatSnap.DataTo(&converInfo)
	if err != nil {
		fmt.Println(err)
		return chatSummary, err
	}

	//get receiverID
	var friend User
	usersID := converInfo.Users
	if userID == usersID[0] {
		friendID = usersID[1]
	} else {
		friendID = usersID[0]
	}
	if err = friend.GetOne(friendID, ""); err != nil {
		fmt.Println("Get_user_id: ", err)
		//return chatSummary, err
	}

	//get last message
	var lastMessage Message
	chatRef := chatSnap.Ref
	messageDocIter := chatRef.Collection(firestoreCol.MESSAGE_COLLECTION).OrderBy("CreatedAt", firestore.Desc).Limit(1).Documents(firebase.Ctx)
	messageSnap, err := messageDocIter.Next()
	if err != nil {
		return chatSummary, err
	}
	err = messageSnap.DataTo(&lastMessage)
	if err != nil {
		return chatSummary, err
	}

	conversationInfo := ChatSummary{
		ChatID:      converInfo.ID,
		Friend:      friend,
		LastMessage: lastMessage.Content,
	}
	return conversationInfo, nil
}
