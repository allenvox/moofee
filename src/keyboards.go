package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var start_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔐 "+cipher_locale[0], "code"),
			tgbotapi.NewInlineKeyboardButtonData("🎸 "+chords_locale[0], "chords"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("♟️ "+chess_locale[0], "chess"),
			tgbotapi.NewInlineKeyboardButtonData("🆘 "+help_locale[0], "help"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔐 "+cipher_locale[1], "code"),
			tgbotapi.NewInlineKeyboardButtonData("🎸 "+chords_locale[1], "chords"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("♟️ "+chess_locale[1], "chess"),
			tgbotapi.NewInlineKeyboardButtonData("🆘 "+help_locale[1], "help"),
		),
	),
}

var current_language string

var help_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✍ "+author_locale[0], "author"),
			tgbotapi.NewInlineKeyboardButtonData("1️⃣0️⃣1️⃣ "+version_locale[0], "version"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🕦 "+time_locale[0], "time"),
			tgbotapi.NewInlineKeyboardButtonData("📅 "+date_locale[0], "date"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(current_language+" "+language_locale[0], "language"),
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+back_locale[0], "start"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✍ "+author_locale[1], "author"),
			tgbotapi.NewInlineKeyboardButtonData("1️⃣0️⃣1️⃣ "+version_locale[1], "version"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🕦 "+time_locale[1], "time"),
			tgbotapi.NewInlineKeyboardButtonData("📅 "+date_locale[1], "date"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(current_language+" "+language_locale[1], "language"),
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+back_locale[1], "start"),
		),
	),
}

var code_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(caesar_locale[0], "caesar"),
			tgbotapi.NewInlineKeyboardButtonData(vigenere_locale[0], "vigenere"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+back_locale[0], "start"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(caesar_locale[1], "caesar"),
			tgbotapi.NewInlineKeyboardButtonData(vigenere_locale[1], "vigenere"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+back_locale[1], "start"),
		),
	),
}

var vigenere_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(encode_locale[0], "vigenere_encode"),
			tgbotapi.NewInlineKeyboardButtonData(decode_locale[0], "vigenere_decode"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+cipher_locale[0], "code"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(encode_locale[1], "vigenere_encode"),
			tgbotapi.NewInlineKeyboardButtonData(decode_locale[1], "vigenere_decode"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+cipher_locale[1], "code"),
		),
	),
}

var chess_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🧩 "+mate_locale[0]+" 1 "+move_locale[0], "mate_in1"),
			tgbotapi.NewInlineKeyboardButtonData("🧩 "+mate_locale[0]+" 2 "+moves_locale[0], "mate_in2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🧩 "+mate_locale[0]+" 3 "+moves_locale[0], "mate_in3"),
			tgbotapi.NewInlineKeyboardButtonData("🧩 "+mate_locale[0]+" 4 "+moves_locale[0], "mate_in4"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_locale[0], "chess"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🧩 "+mate_locale[1]+" 1 "+move_locale[1], "mate_in1"),
			tgbotapi.NewInlineKeyboardButtonData("🧩 "+mate_locale[1]+" 2 "+moves_locale[1], "mate_in2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🧩 "+mate_locale[1]+" 3 "+moves_locale[1], "mate_in3"),
			tgbotapi.NewInlineKeyboardButtonData("🧩 "+mate_locale[1]+" 4 "+moves_locale[1], "mate_in4"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_locale[1], "chess"),
		),
	),
}

var mate_in1_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "m1_1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "m1_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_puzzles_locale[0], "chess"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "m1_1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "m1_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_puzzles_locale[1], "chess"),
		),
	),
}

var mate_in2_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "m2_1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "m2_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_puzzles_locale[0], "chess"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "m2_1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "m2_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_puzzles_locale[1], "chess"),
		),
	),
}

var mate_in3_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "m3_1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "m3_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_puzzles_locale[0], "chess"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "m3_1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "m3_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_puzzles_locale[1], "chess"),
		),
	),
}

var mate_in4_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "m4_1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "m4_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_puzzles_locale[0], "chess"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "m4_1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "m4_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+chess_puzzles_locale[1], "chess"),
		),
	),
}

var chords_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Стрыкало", "strykalo"),
			tgbotapi.NewInlineKeyboardButtonData("Нервы", "nervy"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("кис-кис", "kiskis"),
			tgbotapi.NewInlineKeyboardButtonData(other_locale[0], "other"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+back_locale[0], "start"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Стрыкало", "strykalo"),
			tgbotapi.NewInlineKeyboardButtonData("Нервы", "nervy"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("кис-кис", "kiskis"),
			tgbotapi.NewInlineKeyboardButtonData(other_locale[1], "other"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+back_locale[1], "start"),
		),
	),
}

var strykalo_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Наше лето", "nashe_leto"),
			tgbotapi.NewInlineKeyboardButtonData("Дешёвые драмы", "deshovye_dramy"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Кайен", "kayen"),
			tgbotapi.NewInlineKeyboardButtonData("Фанк", "funk"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Гори", "goris"),
			tgbotapi.NewInlineKeyboardButtonData("Всё решено", "vsyo_resheno"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+artists_locale[0], "chords"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Наше лето", "nashe_leto"),
			tgbotapi.NewInlineKeyboardButtonData("Дешёвые драмы", "deshovye_dramy"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Кайен", "kayen"),
			tgbotapi.NewInlineKeyboardButtonData("Фанк", "funk"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Гори", "goris"),
			tgbotapi.NewInlineKeyboardButtonData("Всё решено", "vsyo_resheno"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+artists_locale[1], "chords"),
		),
	),
}

var nervy_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Слишком влюблён", "slishkom_vlyublon"),
			tgbotapi.NewInlineKeyboardButtonData("Май bye", "may_bye"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Самый дорогой человек", "samy_dorogoy"),
			tgbotapi.NewInlineKeyboardButtonData("Батареи", "batarei"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+artists_locale[0], "chords"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Слишком влюблён", "slishkom_vlyublon"),
			tgbotapi.NewInlineKeyboardButtonData("Май bye", "may_bye"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Самый дорогой человек", "samy_dorogoy"),
			tgbotapi.NewInlineKeyboardButtonData("Батареи", "batarei"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+artists_locale[1], "chords"),
		),
	),
}

var kiskis_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Молчи", "molchi"),
			tgbotapi.NewInlineKeyboardButtonData("Мелочь", "meloch"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Кирилл", "kirill"),
			tgbotapi.NewInlineKeyboardButtonData("ЛБТД", "lbtd"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+artists_locale[0], "chords"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Молчи", "molchi"),
			tgbotapi.NewInlineKeyboardButtonData("Мелочь", "meloch"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Кирилл", "kirill"),
			tgbotapi.NewInlineKeyboardButtonData("ЛБТД", "lbtd"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+artists_locale[1], "chords"),
		),
	),
}

var other_keyboard = []tgbotapi.InlineKeyboardMarkup{
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("The Middle", "middle"),
			tgbotapi.NewInlineKeyboardButtonData("Вахтёрам", "vahteram"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Районы-кварталы", "rayoni"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+artists_locale[0], "chords"),
		),
	),
	tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("The Middle", "middle"),
			tgbotapi.NewInlineKeyboardButtonData("Вахтёрам", "vahteram"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Районы-кварталы", "rayoni"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇ "+artists_locale[1], "chords"),
		),
	),
}

var language_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇬🇧", "en"),
		tgbotapi.NewInlineKeyboardButtonData("🇷🇺", "ru"),
	),
)
