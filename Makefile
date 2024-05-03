## Top-level tasks

default: check

.PHONY: check
check: pre-commit-check

## homebrew

.PHONY: homebrew-install
homebrew-install:
	brew bundle install --file=./Brewfile

## pre-commit

.PHONY: pre-commit-check
pre-commit-check:
	pre-commit run --all-files

.PHONY: pre-commit-clean
pre-commit-clean:
	pre-commit gc

.PHONY: pre-commit-update
pre-commit-update:
	pre-commit autoupdate
