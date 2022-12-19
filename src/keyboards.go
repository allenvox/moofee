package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var start_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔐 Шифры", "code"),
		tgbotapi.NewInlineKeyboardButtonData("🔍 Поиск", "picsearch"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🎸 Аккорды", "chords"),
		tgbotapi.NewInlineKeyboardButtonData("🆘 Помощь", "help"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("♟️ Шахматы", "chess"),
		tgbotapi.NewInlineKeyboardButtonData("🆘 Помощь", "help"),
	),
)

var code_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Цезарь", "caesar"),
		tgbotapi.NewInlineKeyboardButtonData("Виженер", "vigenere"),
	),
)

var chess_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🧩 Головоломка", "puzzle"),
	),
)

var chords_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Стрыкало", "strykalo"),
		tgbotapi.NewInlineKeyboardButtonData("Нервы", "nervy"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Другое", "other"),
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
)

var nervy_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Слишком влюблён", "slishkom_vlyublon"),
	),
)

//todo
var other_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Молчи", "molchi"),
	),
)
