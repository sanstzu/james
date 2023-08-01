package handlers

import (
	"strings"

	"github.com/sanstzu/james/models"
	tele "gopkg.in/telebot.v3"
)

func Connect(c tele.Context) error {
	rawArgs := strings.Split(c.Message().Text, " ")
	if len(rawArgs) < 2 {
		return c.Reply("Usage: /connect <sticker_name> \nsticker_name: ID of the sticker (alphanumeric and underscore only)")
	}

	args := make(map[string]string)

	args["name"] = rawArgs[1]

	stickerName := args["name"] + "_by_" + c.Bot().Me.Username

	chatDetails, err := models.GetChat(c.Chat().ID)
	if err != nil {
		return err
	}

	if chatDetails != nil {
		return c.Reply("This chat already has a sticker set.")
	}

	stickerSetDetails, err := models.GetSticker(stickerName)
	if err != nil {
		return err
	}

	if stickerSetDetails == nil {
		return c.Reply("Sticker with the that particular name is not found.")
	}

	if stickerSetDetails["owner_id"] != c.Sender().ID {
		return c.Reply("You are not the owner of this sticker set (the person who initializes it).")
	}

	addChat := &models.AddChatParams{
		ChatId:    c.Message().Chat.ID,
		StickerId: stickerName,
	}

	err = models.AddChat(*addChat)
	if err != nil {
		return err
	}

	return c.Send("Sticker pack connected!\nLink: https://t.me/addstickers/" + stickerName)

}
