package main

func main() {
	bot := initBot()
	updates := getUpdates(bot)
	var flag = no_flag
	for update := range updates {
		handleMessage(bot, update, &flag)
	}
}
