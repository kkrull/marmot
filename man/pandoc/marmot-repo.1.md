% MARMOT-REPO(1) Version 0.5 | Meta Repo Management Tool
% Kyle Krull
% May 2024

# NAME

**marmot repo** - Work with repositories

- [`marmot repo list`](./repo-list.md)
- [`marmot repo register`](./repo-register.md)

# SYNOPSIS

| **marmot repo** [\-\-help]
| **marmot repo** *sub-command* [*options*...]

# DESCRIPTION

**marmot repo** runs commands that operate upon the Git repositories that have been registered with
**marmot**.

# OPTIONS

-\-help

: Show help

# SUB-COMMANDS

**list**

: List repositories

**register**

: Register repositories to manage

# ENVIRONMENT VARIABLES

**MARMOT_META_REPO**

: Path to the Meta Repo (default: $HOME/meta)

# FILES

*$MARMOT_META_REPO/.marmot/meta-repo.json*

: Repositories that **marmot** knows about

# EXIT STATUS

0

: Success

1+

: Invalid command

# SEE ALSO

*marmot(1)*, *marmot-repo-list(1)*, *marmot-repo-register(1)*

*marmot(7)*
