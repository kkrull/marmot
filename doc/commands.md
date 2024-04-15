# Commands

Marmot is a Command Line Interface with multiple commands, much like Git.

## `marmot exec [--direnv] --project-file <file> <command> [args...]`

Run the given command in each repository registered for a project.

`marmot` switches to each repository's root directory before running the command, to capture any
repository-specific settings that may affect the command.  For example, tools like `direnv` and
`rvm` update environment variables when you change in or out of a directory managed by those tools.

Example:

```sh
$ marmot exec --project-file website.conf node --version
/Users/developer/git/website-api: v20.11.1 #runs from website-api/
/Users/developer/git/website-app: v14.18.1 #runs from website-app/
```

### `--direnv` option

Suppress `direnv` output by setting `DIRENV_LOG_FORMAT=`.

Source: <https://github.com/direnv/direnv/wiki/Quiet-or-Silence-direnv>

```sh
$ marmot exec --direnv --project-file website.conf node --version
/Users/developer/git/website-api: v14.18.1
/Users/developer/git/website-app: v20.11.1
```

### `--project-file <file>` option

Path to the Marmot project file, which lists local repositories associated with the project.

This is a text file, containing the absolute path to each repository on its own line.  Example:

```text
/Users/developer/git/website-api
/Users/developer/git/website-app
```

### Example

```sh
$ marmot exec --direnv --project-file website.conf git branch --show-current
/Users/developer/git/website-api: main
/Users/developer/git/website-app: develop
```

### Interaction with `dotfiles`

My own `dotfiles` are noisy.  I needed a way to turn that off:

```sh
DOTFILES_SILENT='' marmot exec --project-file website.conf wc -l README.md
/Users/developer/git/website-api: 125 README.md
/Users/developer/git/website-app: 128 README.md
```
