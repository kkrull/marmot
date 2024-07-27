# Go

.PHONY: default
default: all

## Environment

### Artifact

executable ?= marmot
ifeq ($(OS),Windows_NT)
	executable ?= marmot.exe
endif

artifact-info:
	$(info Artifacts:)
	$(info - executable: $(executable))
	$(info - OS: $(OS))
	@:

### Paths

prefix ?= /usr/local
exec_prefix ?= $(prefix)

# Executable programs that users can run (including symlinks)
bindir := $(exec_prefix)/bin

# Executable programs to be run by other programs, in a subdirectory thereof
libexecdir := $(exec_prefix)/libexec
libexecdirpkg := $(libexecdir)/marmot

# Read-only architecture-independent data files, in a subdirectory thereof
datarootdir := $(prefix)/share
datadir := $(datarootdir)
datadirpkg := $(datadir)/marmot

.PHONY: path-info
path-info:
	$(info Paths:)
	$(info - bindir: $(bindir))
	$(info - datadir: $(datadir))
	$(info - datadirpkg: $(datadirpkg))
	$(info - datarootdir: $(datarootdir))
	$(info - exec_prefix: $(exec_prefix))
	$(info - libexecdir: $(libexecdir))
	$(info - libexecdirpkg: $(libexecdirpkg))
	$(info - prefix: $(prefix))

### Programs

### Sources

# NB: the wildcard function is not recursive: https://stackoverflow.com/a/2483203/112682
source_main := main.go
sources := $(shell find . -type f -name '*.go' | sort)
versionfile := version

source-info:
	$(info Sources:)
	$(info - source_main: $(source_main))
	$(info - sources: $(sources))
	$(info - versionfile: $(versionfile))
	@:

#. STANDARD TARGETS

.PHONY: all
all: $(executable) #> Build all sources
	@:

.PHONY: clean
clean: #> Remove local build files
	$(RM) $(executable)

.PHONY: install
install: $(executable) #> Install program
	mkdir -p $(libexecdirpkg)
	install -g 0 -o 0 -m 0755 $(executable) $(libexecdirpkg)

	mkdir -p $(bindir)
	ln -f -s $(libexecdirpkg)/$(executable) $(bindir)/marmot

	mkdir -p $(datadirpkg)
	install -g 0 -o 0 -m 0644 $(versionfile) $(datadirpkg)

.PHONY: test
test: #> Run tests
	ginkgo run -r

.PHONY: uninstall
uninstall: #> Uninstall program
	$(RM) $(bindir)/marmot
	$(RM) -R $(datadirpkg)
	$(RM) -R $(libexecdirpkg)

#. OTHER TARGETS

# https://stackoverflow.com/a/47107132/112682
.PHONY: help
help: #> Show this help
	@sed -n \
		-e '/@sed/!s/#[.] */_margin_\n/p' \
		-e '/@sed/!s/:.*#> /:/p' \
		$(MAKEFILE_LIST) \
	| column -ts : | sed -e 's/_margin_//'

.PHONY: info
.NOTPARALLEL: info
info: artifact-info path-info source-info #> Show build information
	@:

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
