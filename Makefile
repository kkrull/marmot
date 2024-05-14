## Sources and artifacts

# https://www.gnu.org/software/make/manual/make.html#Directory-Variables
prefix := /usr/local
exec_prefix := $(prefix)
bindir := $(exec_prefix)/bin

datarootdir := $(prefix)/share
mandir := $(datarootdir)/man
man1dir := $(mandir)/man1

srcdir := $(realpath src)

## Top-level tasks

default: check

.PHONY: check
check: pre-commit-check

.NOTPARALLEL: clean
.PHONY: clean
clean: pre-commit-clean manual-clean

.NOTPARALLEL: install
.PHONY: install
install: link-install

.NOTPARALLEL: install-dependencies
.PHONY: install-dependencies
install-dependencies: brew-install-runtime-deps pre-commit-install

.NOTPARALLEL: remove
.PHONY: remove
remove: link-remove

## homebrew

.PHONY: brew-install-runtime-deps
brew-install-runtime-deps:
	brew bundle install --file=./Brewfile

## links

.PHONY: link-install
link-install:
	ln -s $(srcdir)/marmot.zsh $(bindir)/marmot

.PHONY: link-remove
link-remove:
	$(RM) $(bindir)/marmot

## manuals

# Guide: https://eddieantonio.ca/blog/2015/12/18/authoring-manpages-in-markdown-with-pandoc/
# man-pages reference: https://linux.die.net/man/7/man-pages

# TODO KDK: Install the manual pages too
# https://stackoverflow.com/a/33049378/112682

PANDOC := pandoc
PANDOCFLAGS := -f markdown+definition_lists+line_blocks

manual_sources := $(wildcard man/pandoc/**.md)
# $(info manual_sources is $(manual_sources))

.PHONY: manual-clean
manual-clean: groff-manual-clean
	$(RM) man/markdown/*.md

.PHONY: manual-preview
manual-preview:
	$(PANDOC) $(manual_sources) $(PANDOCFLAGS) -s -t man \
		| mandoc

### groff manuals (a.k.a man pages)

groff_manual_objects := $(patsubst man/pandoc/%.md,man/groff/%.groff,$(manual_sources))
# $(info groff_manual_objects is $(groff_manual_objects))

.PHONY: groff-manual
groff-manual: $(groff_manual_objects)

.PHONY: groff-manual-clean
groff-manual-clean:
	$(RM) man/groff/**.groff

man/groff/%.groff: man/pandoc/%.md
	$(PANDOC) $< $(PANDOCFLAGS) -o $@ -s -t man

### markdown manuals

.PHONY: manual-markdown
manual-markdown: man/markdown/marmot.1.md

man/markdown/%.md: man/pandoc/%.md
	$(PANDOC) $< $(PANDOCFLAGS) -o $@ -s \
		-t markdown-definition_lists-line_blocks

## pre-commit

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
