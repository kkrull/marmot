# Manuals (e.g. "man pages")
# Guide: https://eddieantonio.ca/blog/2015/12/18/authoring-manpages-in-markdown-with-pandoc/
# man-pages conventions: https://linux.die.net/man/7/man-pages

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

PANDOC ?= pandoc
PANDOCFLAGS := -f markdown+definition_lists+line_blocks

### Sources

sources := $(wildcard pandoc/**.md)

.PHONY: source-info
source-info:
	$(info Sources:)
	$(info - sources: $(sources))

## Targets

### Main targets

.PHONY: all check clean info install remove

all: groff-manual markdown-manual

check: ;

clean: groff-manual-clean markdown-manual-clean

.NOTPARALLEL: info
info: install-info source-info groff-manual-info markdown-manual-info

install: groff-manual-install

remove: groff-manual-remove

### groff manual (e.g. man pages) targets

groff_manual_installed := $(wildcard $(man1dir)/marmot*)
groff_manual_objects := $(patsubst pandoc/%.md,groff/%,$(sources))

.PHONY: groff-manual
groff-manual: $(groff_manual_objects)

.PHONY: groff-manual-clean
groff-manual-clean:
	$(RM) groff/**.groff

.PHONY: groff-manual-info
groff-manual-info:
	$(info Groff Manual (man pages):)
	$(info - groff_manual_installed: $(groff_manual_installed))
	$(info - groff_manual_objects: $(groff_manual_objects))

.PHONY: groff-manual-install
groff-manual-install: $(groff_manual_objects)
	mkdir -p $(man1dir)
	install -g 0 -o 0 -m 0644 $(groff_manual_objects) $(man1dir)

.PHONY: groff-manual-preview
groff-manual-preview:
	$(PANDOC) $(sources) $(PANDOCFLAGS) -s -t man \
		| mandoc

.PHONY: groff-manual-remove
groff-manual-remove:
	$(RM) $(groff_manual_installed)

groff/%: pandoc/%.md
	$(PANDOC) $< $(PANDOCFLAGS) -o $@ -s -t man

### markdown manual targets

markdown_manual_objects := $(patsubst pandoc/%.md,markdown/%.md,$(sources))

.PHONY: markdown-manual
markdown-manual: $(markdown_manual_objects)

.PHONY: markdown-manual-clean
markdown-manual-clean:
	$(RM) markdown/**.md

.PHONY: markdown-manual-info
markdown-manual-info:
	$(info Markdown Manual:)
	$(info - markdown_manual_objects: $(markdown_manual_objects))

markdown/%.md: pandoc/%.md
	$(PANDOC) $< $(PANDOCFLAGS) -o $@ -s \
		-t markdown-definition_lists-line_blocks
