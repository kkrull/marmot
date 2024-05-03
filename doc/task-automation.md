# Task Automation

Task automation works with GNU Make.

## Top-level tasks

### `make check`

Run all checks.

### `make install`

Install all tools and Git hooks.

## Homebrew tasks

### `make homebrew-install`

Install homebrew dependencies to provide tools that are used here.

## `pre-commit` tasks

### `make pre-commit-check`

Check all files with `pre-commit`.

### `make pre-commit-clean`

Remove unused tools that were installed by `pre-commit`.

### `make pre-commit-install`

Install Git hooks for `pre-commit`.

### `make pre-commit-update`

Update `pre-commit` plugins.
