## Top-level tasks

default: check

.PHONY: check
check: pre-commit-check

# TODO KDK: Make install task that installs to /usr/local/{bin,man,share}
# Manual page installation: https://stackoverflow.com/a/33049378/112682

.NOTPARALLEL: install-dependencies
.PHONY: install-dependencies
install-dependencies: homebrew-install pre-commit-install

## homebrew

.PHONY: homebrew-install
homebrew-install:
	brew bundle install --file=./Brewfile

## pandoc

# Guide: https://eddieantonio.ca/blog/2015/12/18/authoring-manpages-in-markdown-with-pandoc/
# man-pages reference: https://linux.die.net/man/7/man-pages

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
