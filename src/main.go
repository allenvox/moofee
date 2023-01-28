package main

func main() {
	bot := initBot()              // bot initialization/authorization
	updates := getUpdates(bot)    // get updates range from updates channel
	var flag = no_flag            // flag for various scenarios
	for update := range updates { // in a range of updates
		handleMessage(bot, update, &flag) // handle each update
	}
}
