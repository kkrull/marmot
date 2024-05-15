% MARMOT-INIT(1) Version 0.5 | Meta Repo Management Tool
% Kyle Krull
% May 2024

# NAME

**marmot init** - Initialize a meta repo

# SYNOPSIS

| **marmot init** [**\-\-help**]
| **marmot init**

# DESCRIPTION

Initialize a blank meta repo in the configured directory, if none is already present.

# OPTIONS

**-\-help**

: Show help

# ENVIRONMENT VARIABLES

**MARMOT_META_REPO**

: Path in which to create the Meta Repo (default: $HOME/meta)

# FILES

*$MARMOT_META_REPO/.marmot/meta-repo.json*

: Blank metadata with no registered repositories or categories

# EXIT STATUS

0

: Success

1+

: Invalid command, command failure, or meta repo already exists

# SEE ALSO

*marmot(1)*, *marmot(7)*
