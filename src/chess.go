package main

import "strings"

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
