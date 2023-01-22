all: main

.PHONY: main
main: src/*.go
	go build $^

run:
	make
	./main

clean:
	rm -rf main