# `marmot` command

Marmot is a Command Line Interface with multiple commands, much like Git.

- [`marmot category`](./commands/category.md)
- [`marmot exec`](./commands/exec.md)
- [`marmot init`](./commands/init.md)
- [`marmot meta`](./commands/meta.md)
- [`marmot repo`](./commands/repo.md)

## Usage

`marmot` - Meta Repo Management Tool

### SYNOPSIS

```sh
marmot [--help] [--version]
marmot command [options...]
```

### DESCRIPTION

Marmot creates and maintains a Meta Repository (e.g. a "meta repo"), which can
be used to group several Git repositories by 1 or more arbitrary categories.

Marmot creates a directory structure in the meta repo's file system to mirror
the way that repositories have been categorized, so that there is a
`/:category/:sub-category` directory for each (sub-)category.  Each directory
contains symbolic links back to the Git repositories that are grouped into
the same (sub-)category.

Users run commands from one of these directories in order to restrict
commands to the Git repositories that have that categorization in common.
In this fashion, users can do things like search closely-related
code with `git grep` or open an editor for those Git repositories, without
clutter and noise from irrelevant sources in unrelated repositories.

### OPTIONS

```text
--help        Show help
--version     Show version
```

### COMMANDS

```text
category      Work with categories
exec          Execute a command in multiple repositories
init          Make a new meta repo in the default directory
meta          Information about the meta repo (not the data it manages)
repo          Work with repositories
```

### INSTALLATION

```text
link          Add symlink so you can use this on your path
unlink        Remove symlink for this script
```

### ENVIRONMENT VARIABLES

```text
MARMOT_META_REPO  Path to the Meta Repo (default: $HOME/meta)
```
