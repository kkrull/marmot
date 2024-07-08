# Artifact

executable ?= marmot
ifeq ($(OS),Windows_NT)
	executable ?= marmot.exe
endif

# Standard targets

.PHONY: all
all: $(executable)
	@:

.PHONY: clean
clean:
	$(RM) $(executable)

.PHONY: test
test:
	@:

# Go module targets

sources := $(wildcard *.go **/*.go)
source_main := main.go

.PHONY: debug
debug:
	$(info OS: $(OS))
	$(info - executable: $(executable))
	$(info - sources: $(sources))
	@:

.PHONY: run
run:
	go run .

.PHONY: tidy
tidy:
	go mod tidy

$(executable): $(sources)
	go build -o $(executable) $(source_main)
