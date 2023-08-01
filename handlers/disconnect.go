package handlers

import (
	"errors"

	"github.com/sanstzu/james/models"
	tele "gopkg.in/telebot.v3"
)

func Disconnect(c tele.Context) error {
	chatDetails, err := models.GetChat(c.Chat().ID)
	if err != nil {
		return err
	}

	if chatDetails == nil {
		return c.Reply("This chat has no connected sticker pack.")
	}

	stickerSetDetails, err := models.GetSticker(chatDetails["sticker_id"].(string))
	if err != nil {
		return err
	}

	if stickerSetDetails == nil {
		newError := errors.New("Sticker with the that particular name is not found.")
		return newError
	}

	if stickerSetDetails["owner_id"] != c.Sender().ID {
		return c.Reply("You are not the owner of this sticker set (the person who initializes it).")
	}

	err = models.DeleteChat(c.Chat().ID)
	if err != nil {
		return err
	}

	return c.Send("Sticker pack disconnected!")

}
