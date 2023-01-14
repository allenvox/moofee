package main

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var chess_puzzle_number = 0

func puzzleDescription(name string, flag *int) string {
	var result string = chess_task_locale[language] + string([]rune(name)[1]) + move_locale[language]
	if *flag >= m2 {
		result += plural_locale[language]
	}
	result += "\n\n"
	chess_puzzle_number, _ = strconv.Atoi(string([]rune(name)[3]))
	return result + chess_move_locale[language]
}

var move int = 1
var chess_solutions = [][][]string{
	{ // mate in 1
		{"g5-f7"},
		{"f5-g3"},
	},

	{ // m2
		{"b5-a5", "a5-a1"},
		{"f2-f8", "f8-e8"},
	},

	{ // m3
		{"f8-h8", "d8-d6", "h8-f8"},
		{"c8-g4", "g4-d1", "d1-e1"},
	},

	{ //m4
		{"e1-e6", "d6-e6", "g5-f7", "e6-d6"},
		{"e6-e1", "d4-g1", "g1-f2", "f2-g3"},
	},
}

func handlePuzzle(update tgbotapi.Update, flag *int) string {
	text := update.Message.Text
	var result string = "std"
	move_solution := chess_solutions[*flag-4][chess_puzzle_number-1][move-1]
	switch move {
	case 1:
		if text != move_solution {
			result = chess_phrases[0][language]
		} else {
			if *flag < m2 {
				result = chess_phrases[3][language]
				*flag = no_flag
			} else {
				result = chess_phrases[1][language]
				move++
			}
		}
	case 2:
		if text != move_solution {
			result = chess_phrases[0][language]
		} else {
			if *flag < m3 {
				result = chess_phrases[3][language]
				*flag = no_flag
				move--
			} else {
				result = chess_phrases[2][language]
				move++
			}
		}
	case 3:
		if text != move_solution {
			result = chess_phrases[0][language]
		} else {
			if *flag < m4 {
				result = chess_phrases[3][language]
				*flag = no_flag
				move = 1
			} else {
				result = chess_phrases[4][language]
				move++
			}
		}
	case 4:
		if text != move_solution {
			result = chess_phrases[0][language]
		} else {
			result = chess_phrases[3][language]
			*flag = no_flag
			move = 1
		}
	}
	return result
}
