% MARMOT-CATEGORY-ADD(1) Version 0.5 | Meta Repo Management Tool
% Kyle Krull
% May 2024

# NAME

**marmot category add** - Add repositories to a category

# SYNOPSIS

| **marmot category add** [**\-\-help**]
| **marmot category add** *category* *repository* [...]
| **marmot category add** *category*/*sub-category* *repository* [...]

# DESCRIPTION

**marmot category add** adds each *repository* to a *category* or *sub-category*.  Add a **/**
between *category* and *sub-category*, to work within a sub-category.

# OPTIONS

**-\-help**

: Show help

# ENVIRONMENT VARIABLES

See *marmot-category(1)*.

# FILES

See *marmot-category(1)*.

# EXIT STATUS

0

: Success

1+

: Invalid command or command failure

# EXAMPLE

Add a repository to the "user" category:

```sh
marmot category add user ~/git/dotfiles
```

Add some repositories to the "wily" project (lookout Dr. Light):

```sh
marmot category add project/wily ~/git/robot-masters ~/git/skull-fortress
```

# SEE ALSO

*marmot(1)*, *marmot-category(1)*

*marmot(7)*
