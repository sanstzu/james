package handlers

import (
	"github.com/sanstzu/james/models"
	tele "gopkg.in/telebot.v3"
)

func GetPack(c tele.Context) error {
	response, err := models.GetChat(c.Chat().ID)

	if err != nil {
		return err
	}
	if response == nil {
		return c.Reply("This chat have not initialize any sticker set yet. Please use /initialize or /convert first.")
	}

	stickerName := response["sticker_id"].(string)

	return c.Reply("Sticker Pack: https://t.me/addstickers/" + stickerName)
}
