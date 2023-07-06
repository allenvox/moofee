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
	m3                     // mate in 3 moves
	m4                     // mate in 4 moves
)

func initBot() *tgbotapi.BotAPI {
	tokenContent, _ := ioutil.ReadFile("token.txt") // get token from token.txt
	token := string(tokenContent)                   // convert bytes from txt to string
	if token == "replace_this_line_with_your_token" {
		log.Fatal("Please change your bot token in token.txt")
	}
	bot, err := tgbotapi.NewBotAPI(token) // initialize bot with token via go-telegram-bot-api
	if err != nil {
		log.Fatal("Invalid token")
	}
	bot.Debug = true                                           // enable debug logging
	log.Printf("Authorized on account @%s", bot.Self.UserName) // init log
	return bot                                                 // return init bot
}

func getUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)       // last update (offset = 0)
	u.Timeout = 60                   // set timeout to 60 seconds for last update if no updates
	updates := bot.GetUpdatesChan(u) // get updates channel from bot via API
	return updates
}

func send(bot *tgbotapi.BotAPI, c tgbotapi.Chattable) { // send anything sendable
	if _, err := bot.Send(c); err != nil { // send the chattable objects (msgs, imgs etc.)
		panic(err) // if not succesfully sent - panic & print error
	}
}

func editKeyboard(bot *tgbotapi.BotAPI, update tgbotapi.Update, keyboard tgbotapi.InlineKeyboardMarkup) {
	// keyboard edited markup
	markup := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, keyboard)
	send(bot, markup)
}

func editText(bot *tgbotapi.BotAPI, update tgbotapi.Update, s string) {
	if s != update.CallbackQuery.Message.Text { // update msg text only if it's different from previous
		text := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, s, *update.CallbackQuery.Message.ReplyMarkup)
		send(bot, text)
	}
}
