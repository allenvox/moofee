package main

import (
	"runtime"
	"strconv"
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

var vigenere_switch = 0

func handleKeyboards(bot *tgbotapi.BotAPI, update tgbotapi.Update, flag *int) {
	data := update.CallbackQuery.Data
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}
	switch data {
	/*case "quest":
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, quest_phrases[0])
	send(bot, msg)
	*flag = quest*/
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
		editText(bot, update, "Шифр Виженера")
		editKeyboard(bot, update, vigenere_keyboard)
	case "vigenere_encode", "vigenere_decode":
		text := "Введите фразу для "
		if data == "vigenere_decode" {
			text += "де"
			vigenere_switch = 1
		}
		text += "шифрования"
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
		send(bot, msg)
		*flag = vigenere_phrase
	case "chess":
		editText(bot, update, "Выбери тип шахматной задачи")
		editKeyboard(bot, update, chess_keyboard)
	case "mate_in1":
		editText(bot, update, "🧩 "+mate_locale[language]+" 1 "+move_locale[language])
		editKeyboard(bot, update, mate_in1_keyboard)
	case "mate_in2":
		editText(bot, update, "🧩 "+mate_locale[language]+" 2 "+moves_locale[language])
		editKeyboard(bot, update, mate_in2_keyboard)
	case "mate_in3":
		editText(bot, update, "🧩 "+mate_locale[language]+" 3 "+moves_locale[language])
		editKeyboard(bot, update, mate_in3_keyboard)
	case "mate_in4":
		editText(bot, update, "🧩 "+mate_locale[language]+" 4 "+moves_locale[language])
		editKeyboard(bot, update, mate_in4_keyboard)
	case "m1_1", "m1_2", "m2_1", "m2_2", "m3_1", "m3_2", "m4_1", "m4_2":
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
	case "kiskis":
		editText(bot, update, "кис-кис")
		editKeyboard(bot, update, kiskis_keyboard)
	case "date":
		currentTime := time.Now()
		day := time.Now().Format("Monday")
		if language == ru {
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
		}
		editText(bot, update, currentTime.Format(today_locale[language]+" is 02.01.2006, "+day))
	case "time":
		editText(bot, update, time.Now().Format(time_locale[language]+" "+in_nsk_locale[language]+" (GMT+7): 15:04"))
	case "author":
		editText(bot, update, author_locale[language]+" — @allenvox")
	case "version":
		editText(bot, update, working_locale[language]+" "+runtime.Version())
	default:
		editText(bot, update, getSong(data))
		editKeyboard(bot, update, chords_keyboard)
	}
}

var phrase string

func handleText(bot *tgbotapi.BotAPI, update tgbotapi.Update, flag *int) {
	switch *flag {
	case caesar_phrase:
		phrase = update.Message.Text
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, shift_message_locale[language])
		send(bot, msg)
		*flag = caesar_shift
	case caesar_shift:
		shift, err := strconv.Atoi(update.Message.Text)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, shift_value_mustbe_locale[language]+" "+numeric_locale[language])
			send(bot, msg)
		}
		if shift < 0 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, shift_value_mustbe_locale[language]+" "+natural_locale[language])
			send(bot, msg)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, caesar(phrase, shift))
			send(bot, msg)
			*flag = no_flag
		}
	case vigenere_phrase:
		phrase = update.Message.Text
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите ключ шифрования строки")
		send(bot, msg)
		*flag = vigenere_key
	case vigenere_key:
		key := update.Message.Text
		text := result_locale[language] + ":\n"
		if vigenere_switch > 0 {
			text += vigenereDecode(phrase, key)
		} else {
			text += vigenereEncode(phrase, key)
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		send(bot, msg)
		*flag = no_flag
	case m1, m2, m3, m4:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, handlePuzzle(update, flag))
		send(bot, msg)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, choose_locale[language])
		msg.ReplyMarkup = start_keyboard
		send(bot, msg)
	}
}
