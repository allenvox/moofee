package main

const (
	no_flag = iota
	caesar_phrase
	caesar_shift
	vigenere_phrase
	vigenere_key
	quest
	m2_1
	m2_2
	m3_1
	m3_2
	m4_1
	m4_2
)

func main() {
	bot := initBot()
	updates := getUpdates(bot)
	var flag = no_flag
	for update := range updates {
		handleMessage(bot, update, &flag)
	}
}
