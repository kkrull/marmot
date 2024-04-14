# Marmot

Marmot: the **Meta Repo Management Tool**.

Marmot helps regular people make sense of the many Git repositories around them, in a mad world that
can't stop creating them.  It helps developers create a Meta Repo, which refers to the individual
repositories that make up a large project and provides common configuration and documentation for
them all.

Marmot gives power back to individual developers who are often fighting against entropy, by helping
them organize, relate, and locate information spread out over multiple repositories.  Projects are
encouraged to publish their meta repo and share it with their teams, but individual developers can
still create their own meta repo without needing support or buy-in from anyone else.

Taking even one more step back, Marmot might even help developers locate files and information
contained in all the repositories they have ever worked in.  This is not unlike how people who
develop on the REPL sometimes store their sessions and search through them later.

## Commands

Marmot is a Command Line Interface with multiple commands, much like Git.

### `marmot exec [--direnv] --project-file <file> <command> [args...]`

Run the given command in each repository registered for a project.

`marmot` switches to each repository's root directory before running the command, to capture any
repository-specific settings that may affect the command.  For example, tools like `direnv` and
`rvm` update environment variables when you change in or out of a directory managed by those tools.

Example:

```sh
$ ./marmot-exec.sh --project-file website.conf node --version
/Users/developer/git/website-api: v20.11.1 #runs from website-api/
/Users/developer/git/website-app: v14.18.1 #runs from website-app/
```

#### `--direnv` option

Suppress `direnv` output by setting `DIRENV_LOG_FORMAT=`.

Source: <https://github.com/direnv/direnv/wiki/Quiet-or-Silence-direnv>

```sh
$ ./marmot-exec.sh --direnv --project-file website.conf node --version
/Users/developer/git/website-api: v14.18.1
/Users/developer/git/website-app: v20.11.1
```

#### `--project-file <file>` option

Path to the Marmot project file, which lists local repositories associated with the project.

This is a text file, containing the absolute path to each repository on its own line.  Example:

```text
/Users/developer/git/website-api
/Users/developer/git/website-app
```

#### Example

```sh
$ ./marmot-exec.sh --direnv --project-file website.conf git branch --show-current
/Users/developer/git/website-api: main
/Users/developer/git/website-app: develop
```

## Interactions

### `dotfiles`

My own `dotfiles` are noisy.  I needed a way to turn that off:

```sh
DOTFILES_SILENT='' ./marmot-exec.sh --project-file website.conf wc -l README.md
/Users/developer/git/website-api: 125 README.md
/Users/developer/git/website-app: 128 README.md
```
