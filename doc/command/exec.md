# `marmot exec`

`marmot exec` - Execute a command in multiple repositories

## SYNOPSIS

```sh
marmot exec --help
marmot exec
  [--category <category|sub-category>]
  [--direnv] [--print]
  <shell command> [args...]
```

## DESCRIPTION

This command repeats a given shell command on all repositories matching a
(sub-)category.

marmot exec changes directories to each repository before running the
shell command, to ensure that any path-specific environment settings are
applied.  This is helpful for directory-based tools such as
`direnv`, `fnm`, and `rvm`, which update the shell's path and other
parts of its environment when changing directories.  The usefulness of the
shell command may depend upon it, for example when checking if all
repositories in a project use the same version of Node.js.

## OPTIONS

```text
--direnv        Suppress direnv output when changing directories
--help          Show help
--repo-names    Print repository names `inline` prior to or as a `heading`
                above shell command output
```

### `--direnv`

Suppress `direnv` output by setting `DIRENV_LOG_FORMAT=`.

```sh
$ marmot exec --direnv node --version
/Users/developer/git/website-api: v14.18.1
/Users/developer/git/website-app: v20.11.1
```

Source: <https://github.com/direnv/direnv/wiki/Quiet-or-Silence-direnv>

## TIPS

### Git

- Add `--no-pager` to git commands that pipe to less (and pause for input)

## EXAMPLES

### Scanning

Node: List version of Node.js used in repositories that use direnv+nvm:

```sh
$ marmot exec --category platform/node --direnv \
  node --version
```

### Searching and tracing

Git: Grep for matching source code in all repositories:

```sh
$ marmot exec --category project/robot-masters --repo-names heading \
  git --no-pager grep dungeonType
```

### Unified work

Git: Check which branches are checked out right now:

```sh
$ marmot exec --category project/too-many-microservices \
  git branch --show-current
```

Git: Pull all the things!

```sh
$ marmot exec --repo-names heading \
  git pull --ff-only origin
```

Git: Push all the things!

```sh
$ marmot exec --repo-names heading \
  git push
```
