package main

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func puzzleDescription(name string, flag *int) string {
	var result string = "Задача — поставить мат за "
	if strings.HasPrefix(name, "m2") {
		result += "2"
	} else {
		result += "3"
	}
	result += " хода\n\n"
	switch name {
	case "m2_1":
		*flag = m2_1
	case "m2_2":
		*flag = m2_2
	case "m3_1":
		*flag = m3_1
	case "m3_2":
		*flag = m3_2
	}
	return result + "Формат хода (в сообщении): прошлое_место_фигуры-новое_место_фигуры\nПример: e2-e4"
}

var move int = 1
var chess_phrases = []string{"❌ Неверный ход\nПопробуйте другой", "✅ Отличное начало. Введите второй ход", "✅ Замечательно! Введите третий ход", "🎂 Поздравляю!\n✅ Вы выполнили шахматную задачу ♟"}
var solution_m2_1 = []string{"b5-a5", "a5-a1"}
var solution_m2_2 = []string{"f2-f8", "f8-e8"}
var solution_m3_1 = []string{"-", "-", "-"}
var solution_m3_2 = []string{"-", "-", "-"}

func handlePuzzle(update tgbotapi.Update, flag *int) string {
	text := update.Message.Text
	var result string = "std"

	var move_solution = []string{}
	switch *flag {
	case m2_1:
		move_solution = solution_m2_1
	case m2_2:
		move_solution = solution_m2_2
	case m3_1:
		move_solution = solution_m3_1
	case m3_2:
		move_solution = solution_m3_2
	}

	switch move {
	case 1:
		if text != move_solution[0] {
			result = chess_phrases[0]
		} else {
			result = chess_phrases[1]
			move++
		}
	case 2:
		if text != move_solution[1] {
			result = chess_phrases[0]
		} else {
			if *flag < m3_1 {
				result = chess_phrases[3]
				*flag = no_flag
				move--
			} else {
				result = chess_phrases[2]
				move++
			}
		}
	case 3:
		if text != move_solution[2] {
			result = chess_phrases[0]
		} else {
			result = chess_phrases[3]
			*flag = no_flag
			move = 1
		}
	}
	return result
}
