package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	no_flag = iota
	caesar_phrase
	caesar_shift
	vigenere_phrase
	vigenere_key
)

func main() {
	bot := initBot()
	updates := getUpdates(bot)
	var flag = no_flag
	for update := range updates {
		var chatID int64
		if update.Message != nil {
			chatID = update.Message.Chat.ID
		} else if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.Message.Chat.ID
		}
		msg := tgbotapi.NewMessage(chatID, "")
		msg.ReplyMarkup = start_keyboard
		handleMessage(bot, update, &msg, &flag)
		send(bot, msg)
	}
}
