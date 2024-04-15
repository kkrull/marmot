# Marmot

Marmot: the **Meta Repo Management Tool**.

Marmot helps regular people make sense of the many Git repositories around them, in a mad world that
can't stop creating them.  It helps developers create a Meta Repo, which refers to the individual
repositories that make up a large project and provides common configuration and documentation for
them all.

Marmot gives power back to individual developers who are often fighting against entropy, by helping
them organize, relate, and locate information spread out over multiple repositories.  Projects are
encouraged to publish their meta repo and share it with their teams, but individual developers can
still create their own meta repo without needing support or buy-in from anyone else.

Taking even one more step back, Marmot might even help developers locate files and information
contained in all the repositories they have ever worked in.  This is not unlike how people who
develop on the REPL sometimes store their sessions and search through them later.

## Contents

- [CI/CD Jobs](./doc/cicd-jobs.md)
- [Commands](./doc/commands.md)
- [Installation](#installation)
- [Task Automation](./doc/task-automation.md)
- [Tools](./doc/tools.md)
- [Usage](#usage)

## Future work

- Excluded list (whitespace in paths)
- git pager (turning it off automatically, or adding an option)
- comments in conf files (to disable or ignore)
- tags for repos because this is getting crazy
- `getopts` ala <https://github.com/kkrull/shell-sandbox/>

## Installation

Installs symlinks in `/usr/local/bin/`.

```sh
./src/marmot.sh link
```

## Usage

```sh
# List available commands
marmot
```

See [Commands](./doc/commands.md) for details.

## Version

0.0.1
