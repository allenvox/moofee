package main

import (
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func initBot() *tgbotapi.BotAPI {
	tokenContent, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatal(err)
	}
	token := string(tokenContent)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account @%s", bot.Self.UserName)
	return bot
}
