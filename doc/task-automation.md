# Task Automation

This project uses GNU Make to automate tasks.

## Conventions

- Use conventional [directory variables][gnu-directory-variables] with reasonable defaults.
- Use variables for the names of external programs like `fswatch`, with reasonable defaults.
- Use [standard targets][gnu-standard-targets].  Create separate Makefiles for sub-directories that
  have their own work.

[gnu-directory-variables]: https://www.gnu.org/software/make/manual/make.html#Directory-Variables
[gnu-standard-targets]:
    https://www.gnu.org/software/make/manual/html_node/Standard-Targets.html#Standard-Targets

## Standard Targets

Each `Makefile` contains a set of standard targets.  Each of these targets in the root `Makefile`
calls `make` with the same target, in sub-directories that have their own `Makefile`.

### `make all`

Build everything except documentation.

### `make clean`

Remove files that were built by running `make` earlier.

### `make install`

Install all programs and manuals that are made here.

### `make test`

Run all tests and checks.

### `make uninstall`

Uninstall all programs and manuals made in this repository.

## Homebrew Targets

### `make brew-developer-install`

Install homebrew packages in `Brewfile.developer` that developers need to work on this project.

### `make brew-user-install`

Install homebrew packages in `Brewfile.user` that users need to run the programs built here.

## Manual Page Targets

A separate `man/Makefile` builds manuals.  It converts Pandoc sources to man pages (e.g. `groff` or
`troff`) and to basic Markdown, in `man/groff` and `man/markdown`, respectively.  It has some custom
targets:

### `make groff-manual-preview`

Convert and render manuals as man pages, without installing them anywhere.

### `make groff-manual-watch`

Watch Pandoc source files and render previews of them when they change.

### `make install`

Install man pages to `$(mandir)`.

### `make man`

Build all manuals.

## Other Targets

### `make debug`

Print debugging information, such as the values of variables that affect the build.

## `pre-commit` Targets

### `make pre-commit-clean`

Remove unused tools that were installed by `pre-commit`.

### `make pre-commit-install`

Install Git hooks for `pre-commit`.

### `make pre-commit-run`

Run all `pre-commit` checks on all repository files.

### `make pre-commit-update`

Update `pre-commit` plugins.

## Program Targets

A separate `src/Makefile` handles program sources.  It includes conventional targets that install
symlinks to programs in `$(bindir)`.
