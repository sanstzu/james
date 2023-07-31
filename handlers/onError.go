package handlers

import (
	"log"

	tele "gopkg.in/telebot.v3"
)

func OnError(err error, c tele.Context) error {
	log.Printf("Error: %s", err.Error())
	return c.Reply("An internal error has occured. Please try again later.")
}
