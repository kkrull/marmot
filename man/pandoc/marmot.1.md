% MARMOT(1) Version 0.5 | Meta Repo Management Tool
% Kyle Krull
% May 2024

<!---
man-pages reference: https://linux.die.net/man/7/man-pages
-->

# NAME

**marmot** - Meta Repo Management Tool

# SYNOPSIS

| **marmot** [-\-help] [-\-version]
| **marmot** *command* [*options*...]

# DESCRIPTION

Run the **marmot** *command*, which interacts with the Meta Repo in some way.  See COMMANDS.

A Meta Repo, for the purposes of this program, is a set of references to other Git repositories that
are grouped into 1 or more categories.  **marmot** helps you work with all the repositories in a
category, as if they are a single unit.  See *marmot(7)* to get started.

# OPTIONS

-\-help

: Show help

-\-version

: Prints the **marmot** suite version that the program came from

# COMMANDS

## Meta Repo Commands

**init**

: Create a new meta repo

**meta**

: Information about the meta repo itself

## Category commands

**category**

: Work with categories

## Repository Commands

**exec**

: Execute a shell command in multiple repositories

**repo**

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
