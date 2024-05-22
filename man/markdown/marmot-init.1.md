---
author:
- Kyle Krull
date: May 2024
title: MARMOT-INIT(1) Version 0.5.4 \| Meta Repo Management Tool
---

# NAME

**marmot init** - Initialize a meta repo

# SYNOPSIS

**marmot init** \[**\--help**\]\
**marmot init**

# DESCRIPTION

Initialize a blank Meta Repo in the configured directory, if none is
already present.

# OPTIONS

**\--help**  
Show help

# ENVIRONMENT VARIABLES

**MARMOT_META_REPO**  
Path in which to create the Meta Repo (default: \$HOME/meta)

# FILES

*\$MARMOT_META_REPO/.marmot/meta-repo.json*  
Blank metadata with no registered repositories or categories

# EXIT STATUS

0  
Success

1+  
Invalid command, command failure, or meta repo already exists

# SEE ALSO

[*marmot(1)*](./marmot.1.md), [*marmot(7)*](./marmot.7.md)
