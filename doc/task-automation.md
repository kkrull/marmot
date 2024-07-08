# Task Automation

This project uses GNU Make to automate tasks.

## Conventions

- Use conventional [directory variables][gnu-directory-variables] with reasonable defaults.
- Use variables for the names of external programs like `fswatch`, with reasonable defaults.
- Use [standard targets][gnu-standard-targets].  Create separate Makefiles for sub-directories that
  have their own work.
- Add self-documenting [`help`](https://stackoverflow.com/a/47107132/112682), for relevant targets.

[gnu-directory-variables]: https://www.gnu.org/software/make/manual/make.html#Directory-Variables
[gnu-standard-targets]:
    https://www.gnu.org/software/make/manual/html_node/Standard-Targets.html#Standard-Targets

## Artifacts

### Targets in `man/`

Manuals are written in Markdown.  This project uses Pandoc to generate `groff` manual pages and
Markdown manuals in `man/groff/` and `man/markdown/`, respectively.  The Markdown manuals are
included in the repository, for ease of reading on GitHub.

### Targets in `src/zsh`

There isn't anything to build, since the programs are all scripts.  However, there are conventional
targets to install symlinks to the user's `PATH`.

## Development Environment

### Homebrew Targets

Makefiles include targets to ease the process of installing dependencies, via HomeBrew.

### `pre-commit` Targets

Makefiles include targets to ease the process of running various checks on project sources with
[`pre-commit`](./tools.md#pre-commit).

## Help

Use the `help` target to list relevant targets in each `Makefile`:

```sh
# List targets in root Makefile
make help

# List targets for specific sources
make -C man help
```
