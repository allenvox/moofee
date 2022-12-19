package main

import (
	"runtime"
	"strconv"
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
		if update.Message.Text != "" && !update.Message.IsCommand() {
			handleText(update, msg, flag)
		} else if update.Message.IsCommand() {
			if *flag != no_flag {
				*flag = no_flag
			}
			handleCommands(update, msg)
		}
	}
	if update.CallbackQuery != nil {
		handleKeyboards(bot, update, msg, flag)
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

func handleKeyboards(bot *tgbotapi.BotAPI, update tgbotapi.Update, msg *tgbotapi.MessageConfig, flag *int) {
	data := update.CallbackQuery.Data
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}
	msg.ReplyMarkup = nil
	switch data {
	case "code":
		msg.ReplyMarkup = code_keyboard
	case "picsearch":
		msg.Text = "Отправь фото для поиска"
	case "help":
		msg.Text = "Команды:\n\n/author — разработчик бота\n/v — runtime environment\n/time — время в Новосибирске (GMT+7)\n/date —  сегодняшняя дата"
	case "caesar":
		msg.Text = "Введите фразу для шифрования"
		*flag = caesar_phrase
	case "vigenere":
		msg.Text = "Введите фразу для шифрования"
		*flag = vigenere_phrase
	case "chess":
		msg.ReplyMarkup = chess_keyboard
	case "chords":
		msg.ReplyMarkup = chords_keyboard
	case "nervy":
		msg.ReplyMarkup = nervy_keyboard
	case "strykalo":
		msg.ReplyMarkup = strykalo_keyboard
	case "other":
		msg.ReplyMarkup = other_keyboard
	case "nashe_leto", "kayen", "funk", "deshovye_dramy", "molchi":
		msg.Text = getSong(data)
		msg.ReplyMarkup = start_keyboard
	default:
		msg.ReplyMarkup = start_keyboard
	}
}

var phrase string
var shift int
var key string

func handleText(update tgbotapi.Update, msg *tgbotapi.MessageConfig, flag *int) {
	switch *flag {
	case caesar_phrase:
		phrase = update.Message.Text
		msg.Text = "Введите символьное смещение (например, для преобразования А в Б — 1, А в Я — 32)"
		*flag = caesar_shift
	case caesar_shift:
		num, err := strconv.Atoi(update.Message.Text)
		if err != nil {
			msg.Text = "Значение смещения должно быть числовым"
			panic(err)
		}
		if num < 0 {
			msg.Text = "Значение смещения должно быть больше нуля"
		} else {
			shift = num
			//todo
			msg.Text = ""
			*flag = no_flag
		}
	case vigenere_phrase:
		phrase = update.Message.Text
		msg.Text = "Введите ключ для шифрования строки"
		*flag = vigenere_key
	case vigenere_key:
		key = update.Message.Text
		//todo
		msg.Text = ""
		*flag = no_flag
	default:
		msg.Text = "Выбери то, что тебе нужно"
	}
}
