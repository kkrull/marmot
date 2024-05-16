# Task Automation

This project uses GNU Make to automate tasks.

## Conventions

- Use conventional directory variables like `prefix` with reasonable defaults.  This enables one to
  install to different paths by setting environment variables when running `make`.
- Use variables for the names of external programs like `fswatch`, with reasonable defaults.  This
  enables one to use a different program by setting an environment variable when running `make`.
- Use [standard
  targets](https://www.gnu.org/software/make/manual/html_node/Standard-Targets.html#Standard-Targets)
  like `all` and `install`.  Create separate `Makefiles` for sub-directories that have their own
  work to do for these targets, to avoid the clutter of separate targets for separate artifacts
  (`install-manuals` and `install-program`, for example)

## Main targets

Each `Makefile` contains a set of main targets.  Each of the main targets in the root `Makefile`
calls `make` with the same target, in sub-directories that have their own `Makefile`.

### `make all`

Build everything except documentation.

### `make clean`

Remove files that were built by running `make` earlier.

### `make info`

Print debugging information, such as the values of variables that affect the build.

### `make install`

Install all programs and manuals that are sourced in this repository.

### `make test`

Run all tests and checks.

### `make uninstall`

Uninstall all programs and manuals that are sourced in this repository.

## Homebrew tasks

### `make brew-developer-install`

Install homebrew packages in `Brewfile.developer` that developers need to work on this project.

### `make brew-user-install`

Install homebrew packages in `Brewfile.user` that users need to run the programs built here.

## Manual Page tasks

A separate `man/Makefile` builds manuals.  It converts Pandoc sources to man pages (e.g. `groff` or
`troff`) and to basic Markdown, in `man/groff` and `man/markdown`, respectively.  It includes
conventional targets that install manuals to `$(mandir)`.

It also has some custom targets:

### `make groff-manual-preview`

Convert and render manuals as man pages, without installing them anywhere.

### `make groff-manual-watch`

Watch Pandoc source files and render previews of them when they change.

## `marmot` tasks

A separate `src/Makefile` handles program sources.  It includes conventional targets that install
symlinks to programs in `$(bindir)`.

## `pre-commit` tasks

### `make pre-commit-clean`

Remove unused tools that were installed by `pre-commit`.

### `make pre-commit-install`

Install Git hooks for `pre-commit`.

### `make pre-commit-run`

Run all `pre-commit` checks on all repository files.

### `make pre-commit-update`

Update `pre-commit` plugins.
