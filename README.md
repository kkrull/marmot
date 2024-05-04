# Marmot

Marmot: the **Meta Repo Management Tool**.

Marmot helps regular people make sense of the many Git repositories around them, in a mad world that
can't stop creating them.  It helps developers create a Meta Repo, which refers to the individual
repositories that make up a large project.

Marmot gives power back to individual developers who are often fighting against entropy, by helping
them organize, relate, and locate information spread out over multiple repositories.  Projects are
encouraged to publish their meta repo and share it with their teams, but nothing in the underlying
Git repositories has to change or even know about the existence of a meta repo.  Individual
developers can still create their own meta repo without needing support or buy-in from anyone else.

## Contents

- [Architecture and Design Decisions](./doc/decisions.md)
- [CI/CD Jobs](./doc/cicd-jobs.md)
- [Commands](./doc/commands.md)
- [Installation](#installation)
- [Task Automation](./doc/task-automation.md)
- [Tools](./doc/tools.md)
- [Usage](#usage)

## Future work

- git pager (turning it off automatically, or adding an option)
- tags for repos because this is getting crazy
- add `marmot repository migrate --to=github.com`

## Installation

Installs symlinks in `/usr/local/bin/`.

```sh
./src/marmot.zsh link
```

## Usage

```sh
# List available commands
marmot
```

See [Commands](./doc/commands.md) for details.

## Version

0.0.1
