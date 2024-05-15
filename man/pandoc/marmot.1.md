% MARMOT(1) Version 0.5 | Meta Repo Management Tool
% Kyle Krull
% May 2024

<!---
man-pages reference: https://linux.die.net/man/7/man-pages
-->

# NAME

**marmot** - Meta Repo Management Tool

# SYNOPSIS

| **marmot** [**-\-help**] [**-\-version**]
| **marmot** *command* [*args* ...]

# DESCRIPTION

Run the **marmot** *command* with any *args* to interact with the Meta Repo or the repositories it
tracks, in some way.  See COMMANDS.

A Meta Repo, for the purposes of this program, is a set of references to other Git repositories that
are grouped into 1 or more categories.  **marmot** helps you work with all the repositories in a
category, as if they are a single unit.  See *marmot(7)* to get started.

# OPTIONS

**-\-help**

: Show help

**-\-version**

: Prints the **marmot** suite version that the program came from

# COMMANDS

## Meta Repo Commands

[**init**](./marmot-init.1.md)

: Create a new meta repo

[**meta**](./marmot-meta.1.md)

: Information about the meta repo itself

## Category commands

[**category**](./marmot-category.1.md)

: Work with categories

## Repository Commands

[**exec**](./marmot-exec.1.md)

: Execute a shell command in multiple repositories

[**repo**](./marmot-repo.1.md)

: Work with repositories

# ENVIRONMENT VARIABLES

**MARMOT_META_REPO**

: Path to the Meta Repo (default: $HOME/meta)

# FILES

*$MARMOT_META_REPO/.marmot/meta-repo.json*

: Registered repositories and how they relate to one another

# EXIT STATUS

0

: Success

1+

: Invalid command or command failure

# SEE ALSO

*marmot-category(1)*, *marmot-exec(1)*, *marmot-init(1)*, *marmot-meta(1)*, *marmot-repo(1)*

*marmot(7)*
