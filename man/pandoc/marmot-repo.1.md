% MARMOT-REPO(1) Version 0.6 | Meta Repo Management Tool
% Kyle Krull
% May 2024

# NAME

**marmot repo** - Work with repositories

# SYNOPSIS

| **marmot repo** [**\-\-help**]
| **marmot repo** *sub-command* [*args* ...]

# DESCRIPTION

**marmot repo** runs the given *sub-command* with any *args* to do something with repositories.

# OPTIONS

**-\-help**

: Show help

# SUB-COMMANDS

[**list**](./marmot-repo-list.1.md)

: List repositories

[**prune**](./marmot-repo-prune.1.md)

: Prune references to missing repositories

[**register**](./marmot-repo-register.1.md)

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

[*marmot(1)*](./marmot.1.md), [*marmot-repo-list(1)*](./marmot-repo-list.1.md),
[*marmot-repo-register(1)*](./marmot-repo-register.1.md)

[*marmot(7)*](./marmot.7.md)
