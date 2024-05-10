# Marmot

Marmot: the **Meta Repo Management Tool**.

Marmot helps regular people start to make sense of the many Git repositories around them, in a mad
world that can't stop creating them.  It gives agency back to individual developers who often have
to cope with the entropy of sprawling, poorly-documented architectures that source their code from
dozens of repositories in ways that make sense to but a few.

For more on this topic, see [Why does Marmot exist?](./doc/why.md).

## What it does

Once you tell Marmot which Git repositories you're working with, it gives you a way to categorize
them and then operate upon all the Git repositories in a category _as if they were a single Git
repository_.  As much as one can, at least; Marmot is not meant to replace `git submodules`.

Marmot happens to store this information–meta data about how Git repositories are related to each
other–in its own Git repository.  That's the author's idea of a Meta Repo, in a nutshell. A
developer's understanding of their environment is likely to change over time, leading to new ways of
categorizing code.  Might as well store that information in a Git repository, so you can experiment,
roll back, or even share this information with others.

Once you can group Git repositories together and work with them as if they are a single unit, you
can do new things, like finding code you remember (by name, but you don't remember _where_), tracing
control flow from one repository's sources to another, and checking for consistency (i.e. is
everyone using a reasonably current version of Node.js?).

---

## Contents

### For Users

Start with the basics:

- [Installation](#installation)
- [Getting Started](./doc/getting-started.md)

Then learn more:

- [Commands](./doc/commands.md)
- [Environment Variables](./doc/environment-variables.md)
- [Why does Marmot exist?](./doc/why.md)

### For Developers

If you are planning to contribute to Marmot in some fashion, these may be helpful:

- [Architecture and Design Decisions](./doc/decisions.md)
- [CI/CD Jobs](./doc/cicd-jobs.md)
- [Task Automation](./doc/task-automation.md)
- [Tools](./doc/tools.md)

---

## Installation

This was developed with Linux, MacOS, and Windows Subsystem for Linux (WSL) in mind.

### Clone and link

Clone this repository to a location of your choice.  Marmot doesn't have its own package to install;
just a single command that needs to be on your path somewhere.  You might try something simple like
this:

```sh
# Might require sudo
mkdir -p /usr/local/bin

# Might be unnecessary, if this is already on your path
path+=(/usr/local/bin)

# Install a symlink; might require sudo again
./src/marmot.zsh link
```

### Install dependencies

Marmot uses a few packages that are listed in `Brewfile`.  If you happen to be using Homebrew, try this:

```sh
# Installs `jo` and `jq` commands, if you do not already have them
brew bundle install
```

If you use another package manager such as `apt` (Debian, Ubuntu), there should be similarly named
packages that provide the same commands.  It doesn't matter where `jo` and `jq` come from, as long
as they are reasonably up to date and on your path.

Please also remember to install `zsh` if you do not already have it.  You shouldn't have to use
`zsh` as your main shell; it's just what Marmot uses for its own work.

### Use it

If you can run `marmot --help`, you have a working installation.  If you can re-start your terminal
and it _still_ works, you're in even better shape.

Now head over to [Getting Started](./doc/getting-started.md) to see if Marmot does anything that
might be useful to you.

## Versions

- 0.5.1:
  - [ ] Update [command documentation](./doc/commands.md).
- 0.5: `marmot exec` operates on registered repositories and accepts an optional category.
- 0.4.1: Add `--category` criteria to `marmot repo list`.
- 0.4: `marmot` can be called from anywhere, not just the meta repo.
- 0.3.2: `marmot category create` adds the category to local metadata.
- 0.3.1: Standardize use of code in `src/lib/`.
- 0.3: Add `marmot repo`.
- 0.2: Add `marmot category` and `marmot init`.
- 0.1: Add `marmot exec`.
