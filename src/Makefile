# Marmot program

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

srcdir := $(realpath .)

.PHONY: source-info
source-info:
	$(info Sources:)
	$(info - srcdir: $(srcdir))

## Targets

### Main targets

.PHONY: all check clean info install remove
all: ;
check: ;
clean: ;

.NOTPARALLEL: info
info: install-info source-info ;

install:
	mkdir -p $(bindir)
	ln -f -s $(srcdir)/marmot.zsh $(bindir)/marmot

remove:
	$(RM) $(bindir)/marmot
