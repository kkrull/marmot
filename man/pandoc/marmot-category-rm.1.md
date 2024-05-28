% MARMOT-CATEGORY-RM(1) Version 0.6.1 | Meta Repo Management Tool
% Kyle Krull
% May 2024

# NAME

**marmot category rm** - Remove repositories from a category

# SYNOPSIS

| **marmot category rm** [**\-\-help**]
| **marmot category rm** *category* *repository* [...]
| **marmot category rm** *category*/*sub-category* *repository* [...]

# DESCRIPTION

**marmot category rm** removes each *repository* from a *category* or *sub-category*.  Add a **/**
between *category* and *sub-category*, to work within a sub-category.

# OPTIONS

**-\-help**

: Show help

# ENVIRONMENT VARIABLES

See [*marmot-category(1)*](./marmot-category.1.md).

# FILES

See [*marmot-category(1)*](./marmot-category.1.md).

# EXIT STATUS

0

: Success

1+

: Invalid command or command failure

# EXAMPLE

Remove a newly-categorized repository from the "inbox" category:

```sh
marmot category add lang/shiny ~/git/what-is-this
marmot category rm inbox ~/git/what-is-this
```

# SEE ALSO

[*marmot(1)*](./marmot.1.md), [*marmot-category(1)*](./marmot-category.1.md)

[*marmot(7)*](./marmot.7.md)
