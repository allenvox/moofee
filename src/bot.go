package main

import (
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	no_flag = iota
	caesar_phrase
	caesar_shift
	vigenere_phrase
	vigenere_key
	quest
	m1_1
	m1_2
	m2_1
	m2_2
	m3_1
	m3_2
	m4_1
	m4_2
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

func getUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	return updates
}

func send(bot *tgbotapi.BotAPI, c tgbotapi.Chattable) {
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
	if s != update.CallbackQuery.Message.Text {
		text := tgbotapi.NewEditMessageTextAndMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, s, *update.CallbackQuery.Message.ReplyMarkup)
		if _, err := bot.Send(text); err != nil {
			panic(err)
		}
	}
}
