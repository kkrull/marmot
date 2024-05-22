---
author:
- Kyle Krull
date: May 2024
title: MARMOT-CATEGORY-CREATE(1) Version 0.5.4 \| Meta Repo Management
  Tool
---

# NAME

**marmot category create** - Create a category

# SYNOPSIS

**marmot category create** \[**\--help**\]\
**marmot category create** *category* \[*sub-category* ...\]

# DESCRIPTION

**marmot category create** creates a new *category* and each given
*sub-category*, then adds a directory for each to the Meta Repo.

# OPTIONS

**\--help**  
Show help

# ENVIRONMENT VARIABLES

See [*marmot-category(1)*](./marmot-category.1.md).

# FILES

See [*marmot-category(1)*](./marmot-category.1.md).

# EXIT STATUS

0  
Success

1+  
Invalid command or command failure

# EXAMPLE

Create a "lang" category with sub-categories "java" and "typescript":

``` sh
marmot category create lang java typescript
```

Create a "platform" category with sub-categories "beam,"clr", "jvm", and
"node":

``` sh
marmot category create platform beam clr jvm node
```

Create a "project" category with sub-categories "dotnet-8-migration" and
"skunkworks":

``` sh
marmot category create project dotnet-8-migration skunkworks
```

# SEE ALSO

[*marmot(1)*](./marmot.1.md),
[*marmot-category(1)*](./marmot-category.1.md)

[*marmot(7)*](./marmot.7.md)
