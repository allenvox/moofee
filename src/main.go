package main

import (
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("✍ Шифрование", "code"),
		tgbotapi.NewInlineKeyboardButtonData("📃 Поиск по картинке", "picsearch"),
	),
)

func main() {
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
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			msg := tgbotapi.NewMessage(chatID, "Выбери")
			msg.ReplyMarkup = keyboard

			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			// callback (e.g. keyboard) handler
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			chatID := update.CallbackQuery.Message.Chat.ID
			data := update.CallbackQuery.Data

			msg := tgbotapi.NewMessage(chatID, "Вы нажали на "+data)

			switch data {
			case "code":

			case "picsearch":

			}

			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
