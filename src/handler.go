package main

import (
	"runtime"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	return updates
}

func handleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, msg *tgbotapi.MessageConfig, flag *int) {
	if update.Message != nil {
		handleText(update, msg, flag)
		if update.Message.IsCommand() {
			handleCommands(update, msg)
		}
	}
	if update.CallbackQuery != nil {
		handleKeyboards(bot, update, msg)
	}
}

func handleCommands(update tgbotapi.Update, msg *tgbotapi.MessageConfig) {
	switch update.Message.Command() {
	case "help":
		msg.Text = "Команды:\n\n/author — разработчик бота\n/v — runtime environment\n/time — время в Новосибирске (GMT+7)\n/date —  сегодняшняя дата"
	case "start":
		msg.Text = "Привет!\nЯ простой бот-информатор, но меня быстро учат всему новому!\n\nВведи /help, чтобы посмотреть мои команды :)"
	case "date":
		currentTime := time.Now()
		day := time.Now().Format("Monday")
		switch day {
		case "Monday":
			day = "понедельник"
		case "Tuesday":
			day = "вторник"
		case "Wednesday":
			day = "среда"
		case "Thursday":
			day = "четверг"
		case "Friday":
			day = "пятница"
		case "Saturday":
			day = "суббота"
		case "Sunday":
			day = "воскресенье"
		}
		msg.Text = currentTime.Format("Сегодня 02.01.2006, " + day)
	case "time":
		msg.Text = time.Now().Format("Время в Новосибирске (GMT+7): 15:04")
	case "v":
		msg.Text = "Я работаю на " + runtime.Version()
	case "author":
		msg.Text = "Разработчик — @allenvox"
	default:
		msg.Text = "Что-то я не нашёл такой команды.\nПосмотри в /help"
	}
}

func handleKeyboards(bot *tgbotapi.BotAPI, update tgbotapi.Update, msg *tgbotapi.MessageConfig) {
	data := update.CallbackQuery.Data
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}
	switch data {
	case "code":
		msg.ReplyMarkup = code_keyboard
	case "picsearch":
		msg.Text = "Отправь фото для поиска"
	case "help":
		msg.Text = "Команды:\n\n/author — разработчик бота\n/v — runtime environment\n/time — время в Новосибирске (GMT+7)\n/date —  сегодняшняя дата"
	case "caesar":

	case "vigenere":

	case "chess":
		msg.ReplyMarkup = chess_keyboard
	default:
		msg.ReplyMarkup = start_keyboard
	}
}

func handleText(update tgbotapi.Update, msg *tgbotapi.MessageConfig, flag *int) {
	if update.Message != nil && !update.Message.IsCommand() {

	}
}
