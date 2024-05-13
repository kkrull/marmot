## Top-level tasks

default: check

.PHONY: check
check: pre-commit-check

.NOTPARALLEL: install-dependencies
.PHONY: install-dependencies
install-dependencies: homebrew-install pre-commit-install

## homebrew

.PHONY: homebrew-install
homebrew-install:
	brew bundle install --file=./Brewfile

## manual

.PHONY: manual-preview
manual-preview:
	man -mdoc ./man/marmot.1.mdoc

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
