% MARMOT-CATEGORY(1) Version 0.5 | Meta Repo Management Tool
% Kyle Krull
% May 2024

# NAME

**marmot category** - Work with categories

- [`marmot category add`](./category-add.md)
- [`marmot category create`](./category-create.md)
- [`marmot category list`](./category-list.md)

# SYNOPSIS

| **marmot category** [**\-\-help**]
| **marmot category** *sub-command* [*args* ...]

# DESCRIPTION

**marmot category** runs the given *sub-command* with any *args*, to do something with categories.

# OPTIONS

**-\-help**

: Show help

# SUB-COMMANDS

**add**

: Add repositories to a category

**create**

: Create a new category

**list**

: List categories

# ENVIRONMENT VARIABLES

**MARMOT_META_REPO**

: Path to the Meta Repo (default: $HOME/meta)

# FILES

*$MARMOT_META_REPO/.marmot/meta-repo.json*

: Each category and references to their repositories

# EXIT STATUS

0

: Success

1+

: Invalid command

# SEE ALSO

*marmot(1)*, *marmot-category-add(1)*, *marmot-category-create(1)*, *marmot-category-list(1)*

*marmot(7)*
