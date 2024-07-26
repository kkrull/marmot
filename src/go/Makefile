# Go

.PHONY: default
default: all

## Environment

### Artifact

executable ?= marmot
ifeq ($(OS),Windows_NT)
	executable ?= marmot.exe
endif

debug-artifacts:
	$(info Artifacts:)
	$(info - executable: $(executable))
	$(info - OS: $(OS))
	@:

### Paths

prefix ?= /usr/local
exec_prefix ?= $(prefix)
bindir := $(exec_prefix)/bin

.PHONY: debug-paths
debug-paths:
	$(info Installation paths:)
	$(info - prefix: $(prefix))
	$(info - exec_prefix: $(exec_prefix))
	$(info - bindir: $(bindir))

### Programs

### Sources

# NB: the wildcard function is not recursive: https://stackoverflow.com/a/2483203/112682
source_main := main.go
sources := $(shell find . -type f -name '*.go' | sort)

debug-sources: #> Show debugging information
	$(info Sources:)
	$(info - source_main: $(source_main))
	$(info - sources: $(sources))
	@:

#. STANDARD TARGETS

.PHONY: all
all: $(executable) #> Build all sources
	@:

.PHONY: clean
clean: #> Remove local build files
	$(RM) $(executable)

# TODO KDK: Implement install
.PHONY: install
install: $(executable) #> Install program
	@:

.PHONY: test
test: #> Run tests
	ginkgo run -r

# TODO KDK: Implement uninstall
.PHONY: uninstall
uninstall: #> Uninstall program
	$(RM) $(bindir)/marmot

#. OTHER TARGETS

.PHONY: debug
debug: debug-artifacts debug-paths debug-sources #> Show debugging information
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
install-tools: #> Install Go development tools
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install github.com/onsi/ginkgo/v2/ginkgo
	go install github.com/spf13/cobra-cli@latest
	go install golang.org/x/tools/cmd/godoc@latest
	go install mvdan.cc/gofumpt@latest

.PHONY: test-watch
test-watch: ginkgo-watch #> Run tests in watch mode
	ginkgo watch -r

#. GO TARGETS

.PHONY: format
format: #> Format sources
	gofumpt -l -w .

.PHONY: run
run: #> Run the program
	go run .

.PHONY: tidy
tidy: #> Tidy Go module dependencies
	go mod tidy

.PHONY: update
update: #> Update Go dependencies
	go get -t -u -v

$(executable): $(sources)
	go build -o $(executable) $(source_main)
