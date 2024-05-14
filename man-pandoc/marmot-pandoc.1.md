% MARMOT(1) Version 0.5 | Meta Repo Management Tool
% Kyle Krull
% May 2024

<!---
man-pages reference: https://linux.die.net/man/7/man-pages
-->

# Name

**marmot** - Meta Repo Management Tool

# SYNOPSIS

| **marmot** [-\-help] [-\-version]
| **marmot** *COMMAND* [*options*...]

# DESCRIPTION

Marmot creates and maintains a Meta Repository (e.g. a "meta repo"), which can
be used to group several Git repositories by 1 or more arbitrary categories.

Marmot creates a directory structure in the meta repo's file system to mirror
the way that repositories have been categorized, so that there is a
`/:category/:sub-category` directory for each (sub-)category.  Each directory
contains symbolic links back to the Git repositories that are grouped into
the same (sub-)category.

Users run commands from one of these directories in order to restrict
commands to the Git repositories that have that categorization in common.
In this fashion, users can do things like search closely-related
code with `git grep` or open an editor for those Git repositories, without
clutter and noise from irrelevant sources in unrelated repositories.

# OPTIONS

-\-help

: Show help

-\-version

: Prints the **marmot** suite version that the program came from

# COMMANDS

## META REPO COMMANDS

**init**

: Make a new meta repo in the default directory

**meta**

: Information about the meta repo (not the data it manages)

## CATEGORY AND REPOSITORY COMMANDS

**category**

: Work with categories

**exec**

: Execute a command in multiple repositories

**repo**

: Work with repositories

## INSTALLATION COMMANDS

**link**

: Add symlink so you can use this on your path

**unlink**

: Remove symlink for this script

# ENVIRONMENT VARIABLES

**MARMOT_META_REPO**

: Path to the Meta Repo (default: $HOME/meta)

# FILES

**marmot** reads meta data from $HOME/meta/.marmot/meta-repo.json.

# EXIT STATUS

0

: Success

1+

: Invalid command or command failure
