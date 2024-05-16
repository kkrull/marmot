# Manuals (e.g. "man pages")
# Guide: https://eddieantonio.ca/blog/2015/12/18/authoring-manpages-in-markdown-with-pandoc/
# man-pages conventions: https://linux.die.net/man/7/man-pages

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
man7dir := $(mandir)/man7

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
	$(info - man7dir: $(man7dir))

### Programs

FSWATCH ?= fswatch
MANDOC ?= mandoc
PANDOC ?= pandoc
PANDOCFLAGS := -f markdown+definition_lists+line_blocks

### Sources

sources := $(wildcard pandoc/**.md)
sources_man1 := $(filter %.1.md,$(sources))
sources_man7 := $(filter %.7.md,$(sources))

.PHONY: source-debug
source-debug:
	$(info Sources:)
	$(info - sources: $(sources))
	$(info - sources_man1: $(sources_man1))
	$(info - sources_man7: $(sources_man7))

## Standard Targets

.PHONY: all clean install test uninstall
all:
	@:

clean: groff-manual-clean markdown-manual-clean

install:
	@:

test:
	@:

uninstall: groff-manual-uninstall

## Other Targets

.PHONY: debug
.NOTPARALLEL: debug
debug: path-debug source-debug groff-manual-debug markdown-manual-debug
	@:

.PHONY:
install-man: groff-manual-install

.PHONY: man
man: groff-manual markdown-manual
	@:

### groff manual (e.g. man pages) targets

groff_manual_installed := $(wildcard $(man1dir)/marmot* $(man7dir)/marmot*)
groff_man1_objects := $(patsubst pandoc/%.md,groff/%,$(sources_man1))
groff_man7_objects := $(patsubst pandoc/%.md,groff/%,$(sources_man7))

.PHONY: groff-manual
groff-manual: $(groff_man1_objects) $(groff_man7_objects)

.PHONY: groff-manual-clean
groff-manual-clean:
	$(RM) groff/**

.PHONY: groff-manual-debug
groff-manual-debug:
	$(info Groff Manual (man pages):)
	$(info - groff_manual_installed: $(groff_manual_installed))
	$(info - groff_man1_objects: $(groff_man1_objects))
	$(info - groff_man7_objects: $(groff_man7_objects))

.PHONY: groff-manual-install
groff-manual-install: $(groff_man1_objects) $(groff_man7_objects)
	mkdir -p $(man1dir) $(man7dir)
	install -g 0 -o 0 -m 0644 $(groff_man1_objects) $(man1dir)
	install -g 0 -o 0 -m 0644 $(groff_man7_objects) $(man7dir)

.PHONY: groff-manual-preview
groff-manual-preview:
	$(PANDOC) $(sources) $(PANDOCFLAGS) -s -t man \
		| $(MANDOC)

.PHONY: groff-manual-uninstall
groff-manual-uninstall:
	$(RM) $(groff_manual_installed)

.PHONY: groff-manual-watch
groff-manual-watch:
	$(FSWATCH) $(sources) \
		| xargs -I {} echo "$(PANDOC) {} $(PANDOCFLAGS) -s -t man | $(MANDOC)" \
		| sh

groff/%: pandoc/%.md
	$(PANDOC) $< $(PANDOCFLAGS) -o $@ -s -t man

### markdown manual targets

markdown_manual_objects := $(patsubst pandoc/%.md,markdown/%.md,$(sources))

.PHONY: markdown-manual
markdown-manual: $(markdown_manual_objects)

.PHONY: markdown-manual-clean
markdown-manual-clean:
	$(RM) markdown/**.md

.PHONY: markdown-manual-debug
markdown-manual-debug:
	$(info Markdown Manual:)
	$(info - markdown_manual_objects: $(markdown_manual_objects))

markdown/%.md: pandoc/%.md
	$(PANDOC) $< $(PANDOCFLAGS) -o $@ -s \
		-t markdown-definition_lists-line_blocks
