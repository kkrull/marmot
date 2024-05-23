% MARMOT-REPO-PRUNE(1) Version 0.6 | Meta Repo Management Tool
% Kyle Krull
% May 2024

# NAME

**marmot repo prune** - Prune references to missing repositories

# SYNOPSIS

| **marmot repo prune** [**\-\-help**]
| **marmot repo prune**

# DESCRIPTION

**marmot repo prune** removes references to repositories that are no longer where they used to be,
when they were registered with [*marmot-repo-register(1)*](./marmot-repo-register.1.md).  It does
this by checking each registered repository path.  Any path that is not a directory is removed from
registered repositories and from any categories that included it.

# OPTIONS

**-\-help**

: Show help

# ENVIRONMENT VARIABLES

See [*marmot-repo(1)*](./marmot-repo.1.md).

# FILES

See [*marmot-repo(1)*](./marmot-repo.1.md).

# EXIT STATUS

0

: Success

1+

: Invalid command or command failure

# SEE ALSO

[*marmot(1)*](./marmot.1.md), [*marmot-category(1)*](./marmot-category.1.md),
[*marmot-repo(1)*](./marmot-repo.1.md), [*marmot-repo-list(1)*](./marmot-repo-list.1.md),
[*marmot-repo-register(1)*](./marmot-repo-register.1.md)

[*marmot(7)*](./marmot.7.md)
