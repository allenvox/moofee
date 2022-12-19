package main

const (
	no_flag = iota
	caesar_phrase
	caesar_shift
	vigenere_phrase
	vigenere_key
)

func main() {
	bot := initBot()
	updates := getUpdates(bot)
	var flag = no_flag
	for update := range updates {
		handleMessage(bot, update, &flag)
	}
}
