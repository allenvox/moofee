package main

import (
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const ( // flags
	no_flag = iota
	caesar_phrase
	caesar_shift
	vigenere_phrase
	vigenere_key
	m1
	m2
	m3
	m4
	//quest
)

func initBot() *tgbotapi.BotAPI {
	tokenContent, _ := ioutil.ReadFile("token.txt") // get token from txt file
	token := string(tokenContent)
	bot, _ := tgbotapi.NewBotAPI(token) // initialize bot
	bot.Debug = true
	log.Printf("Authorized on account @%s", bot.Self.UserName)
	return bot
}

func getUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	return updates
}

func send(bot *tgbotapi.BotAPI, c tgbotapi.Chattable) { // send anything sendable
	if _, err := bot.Send(c); err != nil {
		panic(err)
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
		if _, err := bot.Send(text); err != nil {
			panic(err)
		}
	}
}
