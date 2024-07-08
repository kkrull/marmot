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
test: ginkgo-run
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

.PHONY: format
format:
	gofumpt -l -w .

.PHONY: install-tools
install-tools:
	go install github.com/onsi/ginkgo/v2/ginkgo
	go install mvdan.cc/gofumpt@latest

.PHONY: run
run:
	go run .

.PHONY: ginkgo-run
ginkgo-run:
	ginkgo run -r

.PHONY: ginkgo-watch
ginkgo-watch:
	ginkgo watch -r

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: update
update:
	go get -t -u -v

$(executable): $(sources)
	go build -o $(executable) $(source_main)
