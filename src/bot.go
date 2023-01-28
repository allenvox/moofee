package main

import (
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const ( // locales
	en = iota // english = 0
	ru        // russian = 1
)

var language int = en

const ( // flags
	no_flag         = iota // initialize no_flag as 0
	caesar_phrase          // handling phrase for caesar cipher
	caesar_shift           // handling caesar cipher shift
	vigenere_phrase        // handling phrase for vigenere cipher
	vigenere_key           // handling vigenere cipher's key
	m1                     // mate in 1 move
	m2                     // mate in 2 moves
	m3                     // ...
	m4
)

func initBot() *tgbotapi.BotAPI {
	tokenContent, _ := ioutil.ReadFile("token.txt")            // get token from txt file in the main folder (where bot executable is located)
	token := string(tokenContent)                              // convert bytes gotten from txt to string
	bot, _ := tgbotapi.NewBotAPI(token)                        // initialize bot with gotten token via API using @go-telegram-bot-api
	bot.Debug = true                                           // enable debug logging
	log.Printf("Authorized on account @%s", bot.Self.UserName) // log that initialization completed
	return bot                                                 // return initialized bot for the future uses
}

func getUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)       // last update (offset = 0)
	u.Timeout = 60                   // set timeout to 60 seconds for last update if not getting new ones
	updates := bot.GetUpdatesChan(u) // get updates channel from bot via API and our new update
	return updates
}

func send(bot *tgbotapi.BotAPI, c tgbotapi.Chattable) { // send anything sendable
	if _, err := bot.Send(c); err != nil { // send the chattable objects (messages, images etc.)
		panic(err) // if wasn't succesfully sent - panic & print error
	}
}

func editKeyboard(bot *tgbotapi.BotAPI, update tgbotapi.Update, keyboard tgbotapi.InlineKeyboardMarkup) {
	markup := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, keyboard)
	if _, err := bot.Send(markup); err != nil {
		panic(err)
	}
}

func editText(bot *tgbotapi.BotAPI, update tgbotapi.Update, s string) {
	if s != update.CallbackQuery.Message.Text { // update (edit) text only if it's different from previous
		text := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, s, *update.CallbackQuery.Message.ReplyMarkup)
		send(bot, text)
	}
}
