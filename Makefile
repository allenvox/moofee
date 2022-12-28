all: main

.PHONY: main
main: src/main.go src/bot.go src/handler.go src/keyboards.go src/songs.go src/cipher.go src/quest.go src/chess.go
	go build $^

run:
	make
	./main

clean:
	rm -rf main