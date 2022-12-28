package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var start_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔐 Шифры", "code"),
		tgbotapi.NewInlineKeyboardButtonData("🎸 Аккорды", "chords"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("♟️ Шахматы", "chess"),
		tgbotapi.NewInlineKeyboardButtonData("🆘 Помощь", "help"),
	),
)

var help_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("✍ Автор", "author"),
		tgbotapi.NewInlineKeyboardButtonData("1️⃣0️⃣1️⃣ Версия", "version"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🕦 Время", "time"),
		tgbotapi.NewInlineKeyboardButtonData("📅 Дата", "date"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("???", "quest"),
		tgbotapi.NewInlineKeyboardButtonData("⬇ Обратно", "start"),
	),
)

var code_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Цезарь", "caesar"),
		tgbotapi.NewInlineKeyboardButtonData("Виженер", "vigenere"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ Обратно", "start"),
	),
)

var chess_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🧩 Головоломки", "puzzle"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ Шахматы", "chess"),
	),
)

var puzzle_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "puzzle1"),
		tgbotapi.NewInlineKeyboardButtonData("2", "puzzle2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "puzzle3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ Головоломки", "chess"),
	),
)

var chords_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Стрыкало", "strykalo"),
		tgbotapi.NewInlineKeyboardButtonData("Нервы", "nervy"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("кис-кис", "kiskis"),
		tgbotapi.NewInlineKeyboardButtonData("Другое", "other"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ Обратно", "start"),
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
		tgbotapi.NewInlineKeyboardButtonData("⬇ Исполнители", "chords"),
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
		tgbotapi.NewInlineKeyboardButtonData("⬇ Исполнители", "chords"),
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
		tgbotapi.NewInlineKeyboardButtonData("⬇ Исполнители ", "chords"),
	),
)

var other_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("The Middle", "middle"),
		tgbotapi.NewInlineKeyboardButtonData("Вахтёрам", "vahteram"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬇ Исполнители", "chords"),
	),
)
