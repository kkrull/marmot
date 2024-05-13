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

## manual

.PHONY: manual
manual: man/marmot.1.md

man/marmot.1.md: man/marmot.1.mdoc.troff
	mandoc -T markdown ./man/marmot.1.mdoc.troff > ./man/marmot.1.md

.PHONY: manual-preview
manual-preview:
	man -mdoc ./man/marmot.1.mdoc.troff

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
