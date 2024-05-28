% MARMOT-REPO-REGISTER(1) Version 0.6.1 | Meta Repo Management Tool
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

**marmot repo register** resolves symlinks and relative paths in *repository-path*, thereby
registering an absolute path to each repository.  *repository-path* may also end in `.git` or
`.git/` to ease the process of finding and registering lots of Git repositories at once.

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

# EXAMPLE

Register all the things!

```sh
find ~/git -type d -name .git \
  -exec marmot repo register {} +
```

# SEE ALSO

[*marmot(1)*](./marmot.1.md), [*marmot-repo(1)*](./marmot-repo.1.md)

[*marmot(7)*](./marmot.7.md)
