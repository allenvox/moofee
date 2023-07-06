EXE = moofee
SRCS = $(wildcard src/*.go)

all: $(EXE)
.PHONY: $(EXE)
$(EXE): $(SRCS)
	go build -o $@ $^

.PHONY: run
run: $(EXE)
	./$(EXE)

.PHONY: clean
clean:
	rm -f $(EXE)