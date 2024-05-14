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

## manual

# Guide: https://eddieantonio.ca/blog/2015/12/18/authoring-manpages-in-markdown-with-pandoc/
# man-pages reference: https://linux.die.net/man/7/man-pages

# TODO KDK: Install the manual pages too
# https://stackoverflow.com/a/33049378/112682

.PHONY: manual-clean
manual-clean:
	$(RM) -r man/groff/*.groff man/markdown/*.md

.PHONY: manual-groff
manual-groff: man/groff/marmot.1.groff

man/groff/marmot.1.groff: man/pandoc/marmot.1.md
	pandoc ./man/pandoc/marmot.1.md \
		-f markdown+definition_lists+line_blocks \
		-o ./man/groff/marmot.1.groff \
		-s \
		-t man

.PHONY: manual-markdown
manual-markdown: man/markdown/marmot.1.md

man/markdown/marmot.1.md: man/pandoc/marmot.1.md
	pandoc ./man/pandoc/marmot.1.md \
		-f markdown+definition_lists+line_blocks \
		-o ./man/markdown/marmot.1.md \
		-s \
		-t markdown-definition_lists-line_blocks

.PHONY: manual-preview
manual-preview:
	pandoc ./man/pandoc/marmot.1.md \
		-f markdown+definition_lists+line_blocks \
		-s \
		-t man \
		| mandoc

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
