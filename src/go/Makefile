# Go

.PHONY: default
default: all

## Environment

### Artifact

autocomplete := _marmot

# Indentation must be spaces, when not part of a target's commands
# https://stackoverflow.com/a/4713737/112682
executable ?= marmot
ifeq ($(OS),Windows_NT)
  executable ?= marmot.exe
endif

.PHONY: artifact-info
artifact-info:
	$(info Artifacts:)
	$(info - autocomplete: $(autocomplete))
	$(info - executable: $(executable))
	$(info - OS: $(OS))
	@:

### Paths

prefix ?= /usr/local
exec_prefix ?= $(prefix)

# Executable programs that users can run (including symlinks)
bindir := $(exec_prefix)/bin
$(bindir):
	mkdir -p $(bindir)

# Executable programs to be run by other programs, in a subdirectory thereof
libexecdir := $(exec_prefix)/libexec

libexecdirpkg := $(libexecdir)/marmot
$(libexecdirpkg):
	mkdir -p $(libexecdirpkg)

# Read-only architecture-independent data files, in a subdirectory thereof
datarootdir := $(prefix)/share
datadir := $(datarootdir)

datadirpkg := $(datadir)/marmot
$(datadirpkg):
	mkdir -p $(datadirpkg)

# Autocomplete scripts
# https://unix.stackexchange.com/a/607810/37734
datadirzshfn := $(datadir)/zsh/site-functions
$(datadirzshfn):
	mkdir -p $(datadirzshfn)

.PHONY: path-info
path-info:
	$(info Paths:)
	$(info - bindir: $(bindir))
	$(info - datadir: $(datadir))
	$(info - datadirpkg: $(datadirpkg))
	$(info - datadirzshfn: $(datadirzshfn))
	$(info - datarootdir: $(datarootdir))
	$(info - exec_prefix: $(exec_prefix))
	$(info - libexecdir: $(libexecdir))
	$(info - libexecdirpkg: $(libexecdirpkg))
	$(info - prefix: $(prefix))
	@:

### Programs

INSTALLFLAGS := -g 0 -o 0

.PHONY: program-info
program-info:
	$(info Programs:)
	$(info - INSTALLFLAGS: $(INSTALLFLAGS))
	@:

### Sources

# NB: the wildcard function is not recursive: https://stackoverflow.com/a/2483203/112682
source_main := main.go
sources := $(shell find . -type f -name '*.go' | sort)
versionfile := version

.PHONY: source-info
source-info:
	$(info Sources:)
	$(info - source_main: $(source_main))
	$(info - sources: $(sources))
	$(info - versionfile: $(versionfile))
	@:

#. STANDARD TARGETS

.NOTPARALLEL: all
.PHONY: all
all: $(executable) $(autocomplete) #> Build all sources
	@:

.PHONY: clean
clean: #> Remove local build files
	$(RM) $(autocomplete) coverage.html coverage.out coverage-gaps.out $(executable)

.PHONY: install
install: $(executable) | $(bindir) $(datadirpkg) $(libexecdirpkg) #> Install program
	ln -f -s $(libexecdirpkg)/$(executable) $(bindir)/marmot
	install $(INSTALLFLAGS) -m 0644 $(versionfile) $(datadirpkg)
	install $(INSTALLFLAGS) -m 0755 $(executable) $(libexecdirpkg)

.PHONY: test
test: #> Run tests
	ginkgo run -r

.PHONY: uninstall
uninstall: #> Uninstall program
	$(RM) $(bindir)/marmot
	$(RM) -R $(datadirpkg)
	$(RM) $(datadirzshfn)/$(autocomplete)
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
info: artifact-info path-info program-info source-info #> Show build information
	@:

.PHONY: install-autocomplete-zsh
install-autocomplete-zsh: $(autocomplete) | $(datadirzshfn) #> Install autocomplete script for zsh
	install $(INSTALLFLAGS) -m 0644 $(autocomplete) $(datadirzshfn)
	@echo 'Run `source $(datadirzshfn)/$(autocomplete)` to enable auto-completion in zsh.'

.PHONY: install-tools
install-tools: #> Install Go development tools
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install github.com/gregoryv/uncover/cmd/uncover@latest
	go install github.com/onsi/ginkgo/v2/ginkgo
	go install github.com/spf13/cobra-cli@latest
	go install golang.org/x/tools/cmd/godoc@latest
	go install mvdan.cc/gofumpt@latest

.PHONY: test-coverage
test-coverage: coverage.html #> Run tests with code coverage
	@:

.PHONY: test-watch
test-watch: #> Run tests in watch mode
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

$(autocomplete): $(executable)
	./$(executable) completion zsh > $(autocomplete)

coverage.html: coverage.out
	go tool cover -html $< -o $@

# Use `go test` instead of `ginkgo` to include packages that have no coverage
coverage.out: $(sources)
	go test -coverpkg=./... -coverprofile $@ ./...

coverage-gaps.out: coverage.out
	uncover $< > $@

$(executable): $(sources)
	go build -o $(executable) $(source_main)
