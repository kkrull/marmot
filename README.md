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

## Versions

- 0.4: `marmot` can be called from anywhere, not just the meta repo.
  - [x] Add `marmot meta home` to show the path to the meta repo.
  - [ ] Use `MARMOT_HOME` to find meta repo, or fall back to `$HOME/meta`.
- 0.3.2: `marmot category create` adds the category to local metadata.
- 0.3.1: Standardize use of code in `src/lib/`.
- 0.3: Add `marmot repo`.
- 0.2: Add `marmot category` and `marmot init`.
- 0.1: Add `marmot exec`.

### Future work

- all except init: operate on a known/conventional Meta Repo directory, instead
  of the working directory.
- exec: git pager (turning it off automatically, or adding an option).
- exec: Operate on categories instead of project files.
- host: Add `marmot host import <bitbucket.org|github.com>` to register remote
  repositories and clone them.
- repo: add `marmot repo migrate --to=github.com`
