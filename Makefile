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
install: groff-manual-install link-install

.NOTPARALLEL: install-dependencies
.PHONY: install-dependencies
install-dependencies: brew-install-runtime-deps pre-commit-install

.NOTPARALLEL: remove
.PHONY: remove
remove: groff-manual-remove link-remove

## homebrew

.PHONY: brew-install-runtime-deps
brew-install-runtime-deps:
	brew bundle install --file=./Brewfile

## links

.PHONY: link-install
link-install:
	ln -f -s $(srcdir)/marmot.zsh $(bindir)/marmot

.PHONY: link-remove
link-remove:
	$(RM) $(bindir)/marmot

## manuals

# Guide: https://eddieantonio.ca/blog/2015/12/18/authoring-manpages-in-markdown-with-pandoc/
# man-pages reference: https://linux.die.net/man/7/man-pages

PANDOC := pandoc
PANDOCFLAGS := -f markdown+definition_lists+line_blocks

manual_sources := $(wildcard man/pandoc/**.md)
# $(info manual_sources is $(manual_sources))

.PHONY: manual
manual: groff-manual markdown-manual

.PHONY: manual-clean
manual-clean: groff-manual-clean markdown-manual-clean

.PHONY: manual-preview
manual-preview:
	$(PANDOC) $(manual_sources) $(PANDOCFLAGS) -s -t man \
		| mandoc

### groff manuals (a.k.a man pages)

groff_manual_installed := $(wildcard $(man1dir)/marmot*)
# $(info groff_manual_installed is $(groff_manual_installed))

groff_manual_objects := $(patsubst man/pandoc/%.md,man/groff/%,$(manual_sources))
# $(info groff_manual_objects is $(groff_manual_objects))

.PHONY: groff-manual
groff-manual: $(groff_manual_objects)

.PHONY: groff-manual-clean
groff-manual-clean:
	$(RM) man/groff/**.groff

.PHONY: groff-manual-install
groff-manual-install: $(groff_manual_objects)
	mkdir -p $(man1dir)
	install -g 0 -o 0 -m 0644 $(groff_manual_objects) $(man1dir)

.PHONY: groff-manual-remove
groff-manual-remove:
	$(RM) $(groff_manual_installed)

man/groff/%: man/pandoc/%.md
	$(PANDOC) $< $(PANDOCFLAGS) -o $@ -s -t man

### markdown manuals

markdown_manual_objects := $(patsubst man/pandoc/%.md,man/markdown/%.md,$(manual_sources))
# $(info markdown_manual_objects is $(markdown_manual_objects))

.PHONY: markdown-manual
markdown-manual: $(markdown_manual_objects)

.PHONY: markdown-manual-clean
markdown-manual-clean:
	$(RM) man/markdown/**.md

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
