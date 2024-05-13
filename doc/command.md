# `marmot` command

Marmot is a Command Line Interface with multiple commands, much like Git.

- [`marmot category`](./command/category.md)
- [`marmot exec`](./command/exec.md)
- [`marmot init`](./command/init.md)
- [`marmot meta`](./command/meta.md)
- [`marmot repo`](./command/repo.md)

## Usage

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
