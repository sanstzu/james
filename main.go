package main

import (
	"log"

	"github.com/sanstzu/james/config"
	"github.com/sanstzu/james/handlers"
	"github.com/sanstzu/james/models"
)

func main() {

	config.InitializeBot()
	bot := config.GetBot()
	log.Printf("Authorized on account %s", bot.Me.Username)

	models.InitializeFirebase()
	log.Printf("Initialized Firebase")

	bot.Handle("/convert", handlers.Convert)
	bot.Handle("/initialize", handlers.Initialize)
	bot.Handle("/getpack", handlers.GetPack)
	bot.Handle("/help", handlers.Help)

	bot.Start()
}
