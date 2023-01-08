package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var start_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔐 "+cipher_locale[language], "code"),
		tgbotapi.NewInlineKeyboardButtonData("🎸 "+chords_locale[language], "chords"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("♟️ "+chess_locale[language], "chess"),
		tgbotapi.NewInlineKeyboardButtonData("🆘 "+help_locale[language], "help"),
	),
)

var help_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("✍ "+author_locale[language], "author"),
		tgbotapi.NewInlineKeyboardButtonData("1️⃣0️⃣1️⃣ "+version_locale[language], "version"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🕦 "+time_locale[language], "time"),
		tgbotapi.NewInlineKeyboardButtonData("📅 "+date_locale[language], "date"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+back_locale[language], "start"),
	),
)

var code_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(caesar_locale[language], "caesar"),
		tgbotapi.NewInlineKeyboardButtonData(vigenere_locale[language], "vigenere"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+back_locale[language], "start"),
	),
)

var vigenere_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(encode_locale[language], "vigenere_encode"),
		tgbotapi.NewInlineKeyboardButtonData(decode_locale[language], "vigenere_decode"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+cipher_locale[language], "code"),
	),
)

var chess_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🧩 "+mate_locale[language]+" 1"+mate_in_1_locale[language], "mate_in1"),
		tgbotapi.NewInlineKeyboardButtonData("🧩 "+mate_locale[language]+" 2 "+move_locale[language], "mate_in2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🧩 "+mate_locale[language]+" 3 "+move_locale[language], "mate_in3"),
		tgbotapi.NewInlineKeyboardButtonData("🧩 "+mate_locale[language]+" 4 "+move_locale[language], "mate_in4"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_locale[language], "chess"),
	),
)

var mate_in1_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "m1_1"),
		tgbotapi.NewInlineKeyboardButtonData("2", "m1_2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_puzzles_locale[language], "chess"),
	),
)

var mate_in2_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "m2_1"),
		tgbotapi.NewInlineKeyboardButtonData("2", "m2_2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_puzzles_locale[language], "chess"),
	),
)

var mate_in3_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "m3_1"),
		tgbotapi.NewInlineKeyboardButtonData("2", "m3_2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_puzzles_locale[language], "chess"),
	),
)

var mate_in4_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "m4_1"),
		tgbotapi.NewInlineKeyboardButtonData("2", "m4_2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_puzzles_locale[language], "chess"),
	),
)

var chords_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Стрыкало", "strykalo"),
		tgbotapi.NewInlineKeyboardButtonData("Нервы", "nervy"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("кис-кис", "kiskis"),
		tgbotapi.NewInlineKeyboardButtonData(other_locale[language], "other"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+back_locale[language], "start"),
	),
)

var strykalo_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Наше лето", "nashe_leto"),
		tgbotapi.NewInlineKeyboardButtonData("Дешёвые драмы", "deshovye_dramy"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Кайен", "kayen"),
		tgbotapi.NewInlineKeyboardButtonData("Фанк", "funk"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+artists_locale[language], "chords"),
	),
)

var nervy_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Слишком влюблён", "slishkom_vlyublon"),
		tgbotapi.NewInlineKeyboardButtonData("Май bye", "may_bye"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Самый дорогой человек", "samy_dorogoy"),
		tgbotapi.NewInlineKeyboardButtonData("Батареи", "batarei"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+artists_locale[language], "chords"),
	),
)

var kiskis_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Молчи", "molchi"),
		tgbotapi.NewInlineKeyboardButtonData("Мелочь", "meloch"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Кирилл", "kirill"),
		tgbotapi.NewInlineKeyboardButtonData("ЛБТД", "lbtd"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+artists_locale[language], "chords"),
	),
)

var other_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("The Middle", "middle"),
		tgbotapi.NewInlineKeyboardButtonData("Вахтёрам", "vahteram"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ "+artists_locale[language], "chords"),
	),
)
