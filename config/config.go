package config

import (
	"log"
	"time"

	"github.com/sanstzu/james/consts"
	custMiddleware "github.com/sanstzu/james/middleware"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

var bot *tele.Bot

func InitializeBot() {
	pref := tele.Settings{
		Token:  consts.ENV("TELEGRAM_APITOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, c tele.Context) {
			log.Println(err)
			//c.Reply("An internal error has occured. Please try again later.")
		},
	}

	newBot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	newBot.Use(custMiddleware.CustomLogger)
	newBot.Use(middleware.Recover())
	//newBot.Use(middleware.Logger())
	bot = newBot
}

func GetBot() *tele.Bot {
	return bot
}
