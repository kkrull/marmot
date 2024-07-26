# Go

.PHONY: default
default: all

## Environment

### Artifact

executable ?= marmot
ifeq ($(OS),Windows_NT)
	executable ?= marmot.exe
endif

### Paths

### Programs

### Sources

# NB: the wildcard function is not recursive: https://stackoverflow.com/a/2483203/112682
source_main := main.go
sources := $(shell find . -type f -name '*.go' | sort)

#. STANDARD TARGETS

.PHONY: all
all: $(executable)
	@:

.PHONY: clean
clean:
	$(RM) $(executable)

.PHONY: install
install: $(executable)
	@:

.PHONY: test
test:
	ginkgo run -r

#. OTHER TARGETS

.PHONY: debug
debug:
	$(info OS: $(OS))
	$(info - executable: $(executable))
	$(info - source_main: $(source_main))
	$(info - sources: $(sources))
	@:

# https://stackoverflow.com/a/47107132/112682
.PHONY: help
help: #> Show this help
	@sed -n \
		-e '/@sed/!s/#[.] */_margin_\n/p' \
		-e '/@sed/!s/:.*#> /:/p' \
		$(MAKEFILE_LIST) \
	| column -ts : | sed -e 's/_margin_//'

.PHONY: install-tools
install-tools:
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install github.com/onsi/ginkgo/v2/ginkgo
	go install github.com/spf13/cobra-cli@latest
	go install golang.org/x/tools/cmd/godoc@latest
	go install mvdan.cc/gofumpt@latest

.PHONY: test-watch
test-watch: ginkgo-watch
	ginkgo watch -r

#. GO TARGETS

.PHONY: format
format:
	gofumpt -l -w .

.PHONY: run
run:
	go run .

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: update
update:
	go get -t -u -v

$(executable): $(sources)
	go build -o $(executable) $(source_main)
