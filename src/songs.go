package main

import (
	"bufio"
	"os"
)

const (
	strings  = "\nПорядок струн: 654321\nАккорды:\n"
	Am       = "Am: 002210\n"
	E        = "E: 022100\n"
	Em       = "Em: 022000\n"
	G        = "G: 210022\n"
	G3       = "G3: 355433\n"
	Gm3      = "Gm3: 355333\n"
	Gsharpm4 = "G#m4: 466444\n"
	D        = "D: 000232\n"
	Dm       = "Dm: 000231\n"
	A        = "A: 002220\n"
	F        = "F: 133211\n"
	Fm       = "Fm: 133111\n"
	Fsharp   = "F#: 244322\n"
	Fsharpm  = "F#m: 244222\n"
	C        = "C: 032010\n"
	Csharpm4 = "C#m4: 446654\n"
	B        = "B: 224442\n"
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
		song = "*кис-кис — Молчи*\n" + strings + A + Gsharpm4 + Csharpm4 + B
	case "nashe_leto":
		song = "*Валентин Стрыкало — Наше лето*\n" + strings + Am + F + Dm + E
	default:
		return "Упс, песни не существует"
	}
	return song + text
}
