package middleware

import (
	"log"
	"os"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func CustomLogger(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		logger := log.New(os.Stdout, "["+c.Message().Sender.Username+"] ", log.LUTC)
		logger.Printf(strings.Split(c.Message().Text, " ")[0])
		return next(c)
	}
}
