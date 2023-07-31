package models

import (
	"context"
	"strconv"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type AddChatParams struct {
	ChatId    int64
	StickerId string
	OwnerId   int64
}

func AddChat(params AddChatParams) error {
	client := GetClient()
	_, err := client.Collection("chats").Doc(strconv.FormatInt(params.ChatId, 10)).Set(context.Background(), map[string]interface{}{
		"chat_id":    params.ChatId,
		"sticker_id": params.StickerId,
		"owner_id":   params.OwnerId,
	})
	return err
}

func GetChat(chatId int64) (map[string]interface{}, error) {
	client := GetClient()
	iter := client.Collection("chats").Where("chat_id", "==", chatId).Documents(context.Background())

	var isExist bool = false
	var doc *firestore.DocumentSnapshot
	for {
		data, err := iter.Next()
		if err == iterator.Done || isExist {
			break
		}
		if err != nil {
			return nil, err
		}
		isExist = true
		doc = data
	}

	if !isExist {
		return nil, nil
	} else {
		return doc.Data(), nil
	}
}
