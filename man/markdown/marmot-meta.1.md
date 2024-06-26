# NAME

**marmot meta** - Information about the meta repo

# SYNOPSIS

**marmot meta** \[**--help**\]  
**marmot meta** *sub-command* \[*args* …\]

# DESCRIPTION

**marmot meta** runs the given *sub-command* with any *args* to do
something on the Meta Repo.

# OPTIONS

  - **--help**  
    Show help

# SUB-COMMANDS

  - [**home**](./marmot-meta-home.1.md)  
    Show the base directory of the Meta Repo

# ENVIRONMENT VARIABLES

  - **MARMOT\_META\_REPO**  
    Path to the Meta Repo (default: $HOME/meta)

# FILES

  - *$MARMOT\_META\_REPO/.marmot/meta-repo.json*  
    Each category and references to their repositories

# EXIT STATUS

  - 0  
    Success

  - 1+  
    Invalid command

# SEE ALSO

[*marmot(1)*](./marmot.1.md),
[*marmot-meta-home(1)*](./marmot-meta-home.1.md)

[*marmot(7)*](./marmot.7.md)
