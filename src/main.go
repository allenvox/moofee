package main

import (
	"io/ioutil"
	"log"
	"runtime"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var start_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔐 Шифры", "code"),
		tgbotapi.NewInlineKeyboardButtonData("🔍 Поиск", "picsearch"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🆘 Помощь", "help"),
	),
)

var code_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Цезарь", "caesar"),
		tgbotapi.NewInlineKeyboardButtonData("Виженер", "vijener"),
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
		if update.CallbackQuery != nil {
			chatID := update.CallbackQuery.Message.Chat.ID
			data := update.CallbackQuery.Data
			msg := tgbotapi.NewMessage(chatID, "Вы нажали на "+data)
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
			}

			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		}
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID
		msg := tgbotapi.NewMessage(chatID, update.Message.Text)
		if update.Message.IsCommand() {
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
		} else {

		}
		msg.ReplyMarkup = start_keyboard
		if _, err = bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
