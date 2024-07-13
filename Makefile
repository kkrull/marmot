# Marmot

.PHONY: default
default: all

## Environment

### Installation directories

prefix ?= /usr/local
exec_prefix ?= $(prefix)
bindir := $(exec_prefix)/bin

datarootdir := $(prefix)/share
mandir := $(datarootdir)/man
man1dir := $(mandir)/man1

.PHONY: path-debug
path-debug:
	$(info Installation paths:)
	$(info - prefix: $(prefix))
	$(info - exec_prefix: $(exec_prefix))
	$(info - bindir: $(bindir))

	$(info Data paths:)
	$(info - datarootdir: $(datarootdir))
	$(info - mandir: $(mandir))
	$(info - man1dir: $(man1dir))

### Programs

BREW ?= brew
PRECOMMIT ?= pre-commit

### Sources

#. STANDARD TARGETS

.PHONY: all clean install test uninstall
all: #> Build everything
	$(MAKE) -C man all
	$(MAKE) -C src/go all
	$(MAKE) -C src/zsh all

clean: pre-commit-clean #> Remove files built by running make earlier
	$(MAKE) -C man clean
	$(MAKE) -C src/go clean
	$(MAKE) -C src/zsh clean

install: #> Install programs and manuals made here
	$(MAKE) -C man install
	$(MAKE) -C src/go install
	$(MAKE) -C src/zsh install

test: pre-commit-run #> Run all tests and checks
	$(MAKE) -C man test
	$(MAKE) -C src/go test
	$(MAKE) -C src/zsh test

uninstall: #> Uninstall programs and manuals made here
	$(MAKE) -C man uninstall
	$(MAKE) -C src/go uninstall
	$(MAKE) -C src/zsh uninstall

#. OTHER TARGETS

.PHONY: debug
.NOTPARALLEL: debug
debug: path-debug #> Show debugging information
	$(MAKE) -C man debug
	$(MAKE) -C src/go debug
	$(MAKE) -C src/zsh debug

# https://stackoverflow.com/a/47107132/112682
.PHONY: help
help: #> Show this help
	@sed -n \
		-e '/@sed/!s/#[.] *//p' \
		-e '/@sed/!s/:.*#> /:/p' \
		$(MAKEFILE_LIST) \
		| column -ts :

#. HOMEBREW TARGETS

.NOTPARALLEL: brew-install
.PHONY: brew-install
brew-install: brew-developer-install brew-user-install pre-commit-install #> Install depenedencies with Homebrew
	@:

.PHONY: brew-developer-install
brew-developer-install: #> Use HomeBrew to install dependencies that developers need to work here
	$(BREW) bundle install --file=./etc/macos/Brewfile.developer

.PHONY: brew-user-install
brew-user-install: #> Use HomeBrew to install dependencies that users need to run this
	$(BREW) bundle install --file=./etc/macos/Brewfile.user

#. PRE-COMMIT TARGETS

.PHONY: pre-commit-clean
pre-commit-clean: #> Remove pre-commit files that are no longer referenced
	$(PRECOMMIT) gc

.PHONY: pre-commit-install
pre-commit-install: #> Install Git hooks
	$(PRECOMMIT) install

.PHONY: pre-commit-run
pre-commit-run: #> Run pre-commit hooks on all Git files
	$(PRECOMMIT) run --all-files

.PHONY: pre-commit-update
pre-commit-update: #> Update pre-commit plugins
	$(PRECOMMIT) autoupdate
