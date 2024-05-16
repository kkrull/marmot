---
author:
- Kyle Krull
date: May 2024
title: MARMOT-REPO-LIST(1) Version 0.5.1 \| Meta Repo Management Tool
---

# NAME

**marmot repo list** - List repositories

# SYNOPSIS

**marmot repo list** \[**\--help**\]\
**marmot repo list** \[**\--category** *category*\|*sub-category*\]

# DESCRIPTION

**marmot repo list** lists repositories that have been registered with
Marmot. Given options, this lists only the repositories that match the
given criteria.

# OPTIONS

**\--category**  
List repositories that have been added to the given *category* or
*sub-category*

**\--help**  
Show help

# ENVIRONMENT VARIABLES

See [*marmot-repo(1)*](./marmot-repo.1.md).

# FILES

See [*marmot-repo(1)*](./marmot-repo.1.md).

# EXIT STATUS

0  
Success

1+  
Invalid command or command failure

# EXAMPLE

List registered TypeScript repositories

``` sh
marmot repo list --category lang/typescript
```

# SEE ALSO

[*marmot(1)*](./marmot.1.md), [*marmot-repo(1)*](./marmot-repo.1.md)

[*marmot(7)*](./marmot.7.md)
