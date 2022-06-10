all: main

main: src/main.go src/bot.go src/handler.go src/keyboards.go
	go build $^

run:
	make
	./main

clean:
	rm -rf main