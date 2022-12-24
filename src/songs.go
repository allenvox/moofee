package main

import (
	"bufio"
	"os"
)

const (
	s        = "\nПорядок струн: 654321\nАккорды:\n"
	Am       = "Am: 002210\n"
	E        = "E: 022100\n"
	Em       = "Em: 022000\n"
	G        = "G: 210022\n"
	G3       = "G3: 355433\n"
	Gm3      = "Gm3: 355333\n"
	Gsharpm4 = "G#m4: 466444\n"
	D        = "D: 000232\n"
	Dsharp6  = "D#6: 668686\n"
	Dm       = "Dm: 000231\n"
	A        = "A: 002220\n"
	F        = "F: 133211\n"
	F7       = "F7: 88(10)8(10)8\n"
	E7       = "E7: 779797\n"
	Am5      = "Am5: 577555\n"
	Fm       = "Fm: 133111\n"
	Fsharp   = "F#: 244322\n"
	Fsharpm  = "F#m: 244222\n"
	C        = "C: 032010\n"
	Csharpm4 = "C#m4: 446654\n"
	B        = "B: 224442\n"
	Bm       = "Bm: 224432\n"
)

func getText(filename string) string {
	var text string = ""
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		text += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return text
}

func getSong(data string) string {
	var song string
	text := getText("songs/" + data + ".txt")
	switch data {
	case "molchi":
		song = "кис-кис — Молчи\n" + s + A + Gsharpm4 + Csharpm4 + B
	case "nashe_leto":
		song = "Валентин Стрыкало — Наше лето\n" + s + Am + F + Dm + E
	case "slishkom_vlyublon":
		song = "Нервы — Слишком влюблён\n" + s + Csharpm4 + Dsharp6 + Gsharpm4 + E
	case "kayen":
		song = "Валентин Стрыкало — Кайен\n" + s + Am + Dm + E
	case "funk":
		song = "Валентин Стрыкало — Фанк\n" + s + F7 + E7 + Am5 + G3
	case "deshovye_dramy":
		song = "Валентин Стрыкало — Дешёвые драмы\n" + s + Am + F + E + G + Dm
	case "middle":
		song = "Jimmy Eat World — The Middle\n" + s + D + A + G
	case "may_bye":
		song = "Нервы — Май bye\n" + s + Bm + G3 + Am5 + Fsharp
	case "samy_dorogoy":
		song = "Нервы — Самый дорогой человек\n" + s + Em + C + G + D
	case "batarei":
		song = "Нервы — Батареи\n" + s + "+ palm mute\n" + Dm + E + Am + F
	}
	return song + "\n" + text
}
