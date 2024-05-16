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

## Standard Targets

.PHONY: all clean install test uninstall
all:
	$(MAKE) -C man all
	$(MAKE) -C src all

clean: pre-commit-clean
	$(MAKE) -C man clean
	$(MAKE) -C src clean

install:
	$(MAKE) -C man install
	$(MAKE) -C src install

test: pre-commit-run
	$(MAKE) -C man test
	$(MAKE) -C src test

uninstall:
	$(MAKE) -C man uninstall
	$(MAKE) -C src uninstall

## Other Targets

.PHONY: debug
.NOTPARALLEL: debug
debug: path-debug
	$(MAKE) -C man debug
	$(MAKE) -C src debug

.NOTPARALLEL: install-dependencies
.PHONY: install-dependencies
install-dependencies: brew-developer-install brew-user-install pre-commit-install
	@:

.PHONY: install-man
install-man:
	$(MAKE) -C man install-man

### homebrew targets

.PHONY: brew-developer-install
brew-developer-install:
	$(BREW) bundle install --file=./Brewfile.developer

.PHONY: brew-user-install
brew-user-install:
	$(BREW) bundle install --file=./Brewfile.user

### pre-commit targets

.PHONY: pre-commit-clean
pre-commit-clean:
	$(PRECOMMIT) gc

.PHONY: pre-commit-install
pre-commit-install:
	$(PRECOMMIT) install

.PHONY: pre-commit-run
pre-commit-run:
	$(PRECOMMIT) run --all-files

.PHONY: pre-commit-update
pre-commit-update:
	$(PRECOMMIT) autoupdate
