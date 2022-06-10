all: main

main: src/main.go
	go build $^

run:
	make
	./main

clean:
	rm -rf main