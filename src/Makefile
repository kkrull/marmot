# Marmot programs

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

### Sources

srcdir := $(realpath .)

.PHONY: source-debug
source-debug:
	$(info Sources:)
	$(info - srcdir: $(srcdir))

## Standard Targets

.PHONY: all clean install test uninstall
all:
	@:

clean:
	@:

install:
	mkdir -p $(bindir)
	ln -f -s $(srcdir)/marmot.zsh $(bindir)/marmot

test:
	@:

uninstall:
	$(RM) $(bindir)/marmot

## Other Targets

.PHONY: debug
.NOTPARALLEL: debug
debug: path-debug source-debug
	@:
