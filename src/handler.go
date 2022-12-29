package main

import (
	"runtime"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, flag *int) {
	if update.Message != nil {
		if update.Message.Text != "" && !update.Message.IsCommand() {
			handleText(bot, update, flag)
		} else if update.Message.IsCommand() {
			if *flag != no_flag {
				*flag = no_flag
			}
			handleCommands(bot, update)
		}
	}
	if update.CallbackQuery != nil {
		handleKeyboards(bot, update, flag)
	}
}

func handleCommands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	send(bot, tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "У меня нет команд, тыкай кнопки"))
}

func handleKeyboards(bot *tgbotapi.BotAPI, update tgbotapi.Update, flag *int) {
	data := update.CallbackQuery.Data
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}
	switch data {
	case "quest":
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, quest_phrases[0])
		send(bot, msg)
		*flag = quest
	case "code":
		editText(bot, update, "Шифры")
		editKeyboard(bot, update, code_keyboard)
	case "help":
		editText(bot, update, "Помощь")
		editKeyboard(bot, update, help_keyboard)
	case "caesar":
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введите фразу для шифрования")
		send(bot, msg)
		*flag = caesar_phrase
	case "vigenere":
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введите фразу для шифрования")
		send(bot, msg)
		*flag = vigenere_phrase
	case "chess":
		editText(bot, update, "Выбери тип шахматной задачи")
		editKeyboard(bot, update, chess_keyboard)
	case "mate_in2":
		editText(bot, update, "🧩 Мат в 2 хода")
		editKeyboard(bot, update, mate_in2_keyboard)
	case "mate_in3":
		editText(bot, update, "🧩 Мат в 3 хода")
		editKeyboard(bot, update, mate_in3_keyboard)
	case "m2_1", "m2_2", "m3_1", "m3_2":
		picture := tgbotapi.NewPhoto(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("puzzles/"+data+".png"))
		picture.Caption = puzzleDescription(data, flag)
		send(bot, picture)
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введите первый ход")
		send(bot, msg)
	case "chords":
		editText(bot, update, "Аккорды")
		editKeyboard(bot, update, chords_keyboard)
	case "nervy":
		editText(bot, update, "Нервы")
		editKeyboard(bot, update, nervy_keyboard)
	case "strykalo":
		editText(bot, update, "Валентин Стрыкало")
		editKeyboard(bot, update, strykalo_keyboard)
	case "other":
		editText(bot, update, "Другое")
		editKeyboard(bot, update, other_keyboard)
	case "vahteram", "lbtd", "kirill", "meloch", "nashe_leto", "kayen", "funk", "deshovye_dramy", "molchi", "middle", "slishkom_vlyublon", "may_bye", "samy_dorogoi", "batarei":
		editText(bot, update, getSong(data))
		editKeyboard(bot, update, chords_keyboard)
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
		editText(bot, update, currentTime.Format("Сегодня 02.01.2006, "+day))
	case "time":
		editText(bot, update, time.Now().Format("Время в Новосибирске (GMT+7): 15:04"))
	case "author":
		editText(bot, update, "Разработчик — @allenvox")
	case "version":
		editText(bot, update, "Я работаю на "+runtime.Version())
	default:
		editText(bot, update, "Выбери то, что нужно")
		editKeyboard(bot, update, start_keyboard)
	}
}

var phrase string

func handleText(bot *tgbotapi.BotAPI, update tgbotapi.Update, flag *int) {
	switch *flag {
	case caesar_phrase:
		phrase = update.Message.Text
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите символьное смещение (например, для преобразования А в Б — 1, А в Я — 32)")
		send(bot, msg)
		*flag = caesar_shift
	case caesar_shift:
		shift, err := strconv.Atoi(update.Message.Text)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Значение смещения должно быть числовым")
			send(bot, msg)
		}
		if shift < 0 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Значение смещения должно быть больше нуля")
			send(bot, msg)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, caesar(phrase, shift))
			send(bot, msg)
			*flag = no_flag
		}
	case vigenere_phrase:
		phrase = update.Message.Text
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите ключ для шифрования строки")
		send(bot, msg)
		*flag = vigenere_key
	case vigenere_key:
		key := update.Message.Text
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, vigenere(phrase, key))
		send(bot, msg)
		*flag = no_flag
	case quest:
		switch strings.ToLower(update.Message.Text) {
		case quest_keywords[0]:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quest_phrases[1])
			send(bot, msg)
		case quest_keywords[1]:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quest_phrases[2])
			send(bot, msg)
		case quest_keywords[2]:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quest_phrases[3])
			send(bot, msg)
		case quest_keywords[3]:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quest_phrases[4])
			send(bot, msg)
		case quest_keywords[4]:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quest_phrases[5])
			send(bot, msg)
		case quest_keywords[5]:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, quest_phrases[6])
			send(bot, msg)
			*flag = no_flag
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "🤔 Кажется, что-то не то...\nПопробуй ввести другое.")
			send(bot, msg)
		}
	case m2_1, m2_2, m3_1, m3_2:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, handlePuzzle(update, flag))
		send(bot, msg)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выбери то, что нужно")
		msg.ReplyMarkup = start_keyboard
		send(bot, msg)
	}
}
