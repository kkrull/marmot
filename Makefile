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

srcdir := $(realpath src)

.PHONY: source-info
source-info:
	$(info Sources:)
	$(info - srcdir: $(srcdir))

## Targets

### Main targets

.PHONY: all
all:
	$(MAKE) -C man all

.PHONY: check
check: pre-commit-check

.NOTPARALLEL: clean
.PHONY: clean
clean: pre-commit-clean
	$(MAKE) -C man clean

.PHONY: info
info: install-info source-info
	$(MAKE) -C man info

.NOTPARALLEL: install
.PHONY: install
install: marmot-install
	$(MAKE) -C man install

.NOTPARALLEL: install-dependencies
.PHONY: install-dependencies
install-dependencies: brew-developer-install brew-user-install pre-commit-install

.NOTPARALLEL: remove
.PHONY: remove
remove: marmot-remove
	$(MAKE) -C man remove

### homebrew targets

.PHONY: brew-developer-install
brew-developer-install:
	brew bundle install --file=./Brewfile.developer

.PHONY: brew-user-install
brew-user-install:
	brew bundle install --file=./Brewfile.user

### marmot targets

.PHONY: marmot-install
marmot-install:
	mkdir -p $(bindir)
	ln -f -s $(srcdir)/marmot.zsh $(bindir)/marmot

.PHONY: marmot-remove
marmot-remove:
	$(RM) $(bindir)/marmot

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
