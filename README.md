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
- [Environment Variables](./doc/environment-variables.md)
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

See [Commands](./doc/commands.md) and [Environment Variables](./doc/environment-variables.md) for
details.

## Versions

- 0.5: `marmot exec` operates on repositories in a matching category now, instead of a project list.
  - [ ] Update `exec` to work on those matching repositories.
  - [ ] Update [documentation](./doc/commands.md)
- 0.4.1: Add `--category` criteria to `marmot repo list`.
- 0.4: `marmot` can be called from anywhere, not just the meta repo.
- 0.3.2: `marmot category create` adds the category to local metadata.
- 0.3.1: Standardize use of code in `src/lib/`.
- 0.3: Add `marmot repo`.
- 0.2: Add `marmot category` and `marmot init`.
- 0.1: Add `marmot exec`.

### Future work

- Auto-complete for zsh.
- Error handling: `set -euo pipefail` more consistently. See
  <https://www.mulle-kybernetik.com/modern-bash-scripting/state-euxo-pipefail.html>.
- `exec`: Consider adding an option for whether to exit on the first failure, or keep going.
  Or ask the user if/when the first failure happens, since you probably don't know in advance.
- `host`: Add `marmot host import <host: bitbucket.org|github.com>` to register remote
  repositories and `marmot host clone` to clone them.
- `repo`: Add `marmot repo move <host> repository...`.
