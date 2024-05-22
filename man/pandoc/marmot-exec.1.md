% MARMOT-EXEC(1) Version 0.5.4 | Meta Repo Management Tool
% Kyle Krull
% May 2024

# NAME

**marmot exec** - Execute a command in multiple repositories

# SYNOPSIS

| **marmot exec** [**\-\-help**]
| **marmot exec**
  [**\-\-category** *category*|*sub-category*]
  [**\-\-direnv**] [**\-\-repo-names** **heading**|**inline**]
  *shell-command* [*args* ...]

# DESCRIPTION

**marmot exec** repeats a given *shell-command* with any *args* on each matching repository.  The
repositories are either all registered repositories, or those matching a given *category* or
*sub-category*.

**marmot exec** changes directories to each repository before running *shell-command*, to ensure
that any path-specific environment settings are applied.  This is helpful for directory-based tools
such as `direnv`, `fnm`, and `rvm`, which update the shell's path and other parts of its environment
when changing directories.  The usefulness of the shell command may depend upon it, for example when
checking if all repositories in a project use the same version of Node.js.

# OPTIONS

**\-\-direnv**

: Suppress `direnv` output when changing directories

**\-\-help**

: Show help

**\-\-repo\-names**

: Print repository names **inline** prior to output from *shell-command*, or as a **heading** above
it

# ENVIRONMENT VARIABLES

**MARMOT_META_REPO**

: Path to the Meta Repo (default: $HOME/meta)

# FILES

*$MARMOT_META_REPO/.marmot/meta-repo.json*

: Categories and registered repositories

# EXIT STATUS

0

: Success

1+

: Invalid command or **shell-command** failure

# NOTES

## `direnv`

Suppress `direnv` output by setting `DIRENV_LOG_FORMAT=`.

```sh
$ marmot exec --direnv node --version
/Users/developer/git/website-api: v14.18.1
/Users/developer/git/website-app: v20.11.1
```

Source: <https://github.com/direnv/direnv/wiki/Quiet-or-Silence-direnv>

## Git

Add `--no-pager` to git commands that pipe to less (and pause for input).

# EXAMPLE

## Scanning

Node: List version of Node.js used in repositories that use direnv+nvm:

```sh
marmot exec --category platform/node --direnv \
  node --version
```

## Searching and Tracing

Git: Grep for matching source code in all repositories:

```sh
marmot exec --category project/robot-masters --repo-names heading \
  git --no-pager grep dungeonType
```

## Unified Work

Git: Check which branches are checked out right now:

```sh
marmot exec --category project/too-many-microservices \
  git branch --show-current
```

Git: Pull all the things!

```sh
marmot exec --repo-names heading \
  git pull --ff-only origin
```

Git: Push all the things!

```sh
marmot exec --repo-names heading \
  git push
```

# SEE ALSO

[*marmot(1)*](./marmot.1.md), [*marmot(7)*](./marmot.7.md)
