package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot := initBot()
	updates := getUpdates(bot)
	for update := range updates {
		var chatID int64
		if update.Message != nil {
			chatID = update.Message.Chat.ID
		} else if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.Message.Chat.ID
		}
		msg := tgbotapi.NewMessage(chatID, "Выбери")
		msg.ReplyMarkup = start_keyboard
		handleCommands(update, &msg)
		handleKeyboards(bot, update, &msg)
		send(bot, msg)
	}
}
