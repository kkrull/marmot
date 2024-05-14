# Marmot

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

### Sources

## Targets

### Main targets

.PHONY: all check clean info install remove

all:
	$(MAKE) -C man all
	$(MAKE) -C src all

check: pre-commit-check
	$(MAKE) -C man check
	$(MAKE) -C src check

clean: pre-commit-clean
	$(MAKE) -C man clean
	$(MAKE) -C src clean

NOTPARALLEL: info
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

remove:
	$(MAKE) -C man remove
	$(MAKE) -C src remove

### homebrew targets

.PHONY: brew-developer-install
brew-developer-install:
	brew bundle install --file=./Brewfile.developer

.PHONY: brew-user-install
brew-user-install:
	brew bundle install --file=./Brewfile.user

### pre-commit targets

.PHONY: pre-commit-check
pre-commit-check:
	pre-commit run --all-files

.PHONY: pre-commit-clean
pre-commit-clean:
	pre-commit gc

.PHONY: pre-commit-install
pre-commit-install:
	pre-commit install

.PHONY: pre-commit-update
pre-commit-update:
	pre-commit autoupdate
