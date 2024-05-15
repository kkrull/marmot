% MARMOT-REPO-REGISTER(1) Version 0.5 | Meta Repo Management Tool
% Kyle Krull
% May 2024

# NAME

**marmot repo register** - Register repositories to manage

# SYNOPSIS

| **marmot repo register** [**\-\-help**]
| **marmot repo register** *repository-path* [...]

# DESCRIPTION

**marmot repo register** registers the each given *repository-path*, so **marmot** can start to
categorize and operate upon them.

# OPTIONS

**-\-help**

: Show help

# ENVIRONMENT VARIABLES

See *marmot-repo(1)*.

# FILES

See *marmot-repo(1)*.

# EXIT STATUS

0

: Success

1+

: Invalid command or command failure

# EXAMPLE

Register all the things!

```sh
find -s ~/git -type d -name .git \
  | sed 's/[/][.]git$//g' \
  | xargs -I {} marmot repo register {}
```

# SEE ALSO

*marmot(1)*, *marmot-repo(1)*

*marmot(7)*
