# Marmot

Marmot: the Multi Repo Management Tool.

Marmot helps regular people make sense of the many Git repositories around them, in a mad world that
can't stop creating them.  It helps developers create a Meta Repo, which refers to the individual
repositories that make up a large project and provides common configuration and documentation for
them all.

Marmot gives power back to individual developers who are often fighting against entropy, by helping
them organize, relate, and locate information spread out over multiple repositories.  Projects are
encouraged to publish their meta repo and share it with their teams, but individual developers can
create their own meta repo without needing support or buy-in from anyone else.

Taking even one more step back, Marmot might even help developers locate files and information
contained in all the repositories they have ever worked in.  This is not unlike how functional
programmers sometimes store a life-long REPL session or how Obsidian Notes helps people organize
documents.

## Commands

Marmot is a Command Line Interface with multiple commands, much like Git.

### `marmot exec <command> [args...]`

Run the given command in each repository registered for the project, switching to each repository's
root directory first.

Example:

```sh
$ marmot exec node --version
website-api: v20.11.1
website-app: v14.18.1
```

Another example, where there is some extra output from the interactive shell and from `direnv`.

```sh
$ ./marmot-exec.sh node --version 2>/dev/null
...output from starting interactive shell...
website-api: v14.18.1
website-app: v20.11.1
```

## Interactions

### `direnv`

`direnv` outputs a lot of extra information.  Suppress output by setting `DIRENV_LOG_FORMAT=''`:

```sh
$ DIRENV_LOG_FORMAT='' ./marmot-exec.sh node --version
website-api: v14.18.1
website-app: v20.11.1
```

Source: <https://github.com/direnv/direnv/wiki/Quiet-or-Silence-direnv>

### `dotfiles`

My own `dotfiles` are noisy.  I needed a way to turn that off:

```sh
DOTFILES_SILENT='' ./marmot-exec.sh wc -l README.md
website-api:      125 README.md
website-app:      128 README.md
```
