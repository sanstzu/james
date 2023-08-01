package handlers

import (
	tele "gopkg.in/telebot.v3"
)

func Help(c tele.Context) error {
	return c.Reply("/getpack: Getting the current sticker pack in the group\n/initialize: Start a sticker pack in the group and convert a photo (not a document) to the first sticker\n/convert: Import a photo (not a document) to a sticker\n/connect: Connect the current's chat sticker pack to an existing one (must be disconnected first)\n/disconnect: Disconnect the current's chat sticker pack\n/help: Show this message")
}
