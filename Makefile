# Marmot

.PHONY: default
default: all

## Environment

### Installation directories

# https://www.gnu.org/software/make/manual/make.html#Directory-Variables
prefix ?= /usr/local
exec_prefix ?= $(prefix)
bindir := $(exec_prefix)/bin

datarootdir := $(prefix)/share
mandir := $(datarootdir)/man
man1dir := $(mandir)/man1

.PHONY: install-info
install-info:
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

## Targets

### Main targets

.PHONY: all clean info install test uninstall

all:
	$(MAKE) -C man all
	$(MAKE) -C src all

clean: pre-commit-clean
	$(MAKE) -C man clean
	$(MAKE) -C src clean

.NOTPARALLEL: info
info: install-info
	$(MAKE) -C man info
	$(MAKE) -C src info

install:
	$(MAKE) -C man install
	$(MAKE) -C src install

.NOTPARALLEL: install-dependencies
.PHONY: install-dependencies
install-dependencies: brew-developer-install brew-user-install pre-commit-install
	@:

test: pre-commit-run
	$(MAKE) -C man test
	$(MAKE) -C src test

uninstall:
	$(MAKE) -C man uninstall
	$(MAKE) -C src uninstall

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
