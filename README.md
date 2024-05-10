# Marmot

Marmot: the **Meta Repo Management Tool**.

Marmot helps regular people start to make sense of the many Git repositories around them, in a mad
world that can't stop creating them.  It gives agency back to individual developers who often have
to cope with the entropy of sprawling, poorly-documented architectures that source their code from
dozens of repositories in ways that make sense to but a few.

## Why it exists

If you have a small, well-organized architecture where you have a clear idea about which
responsibilities are handled by which codebase, you may be working on a project where someone has
already separated responsibilities in a clear, maintainable manner.  Or you may have enough context
or tribal knowledge about the product to get by without it.  Marmot may not be the tool for you.

The rest of us often find ourselves learning how to take care of someone else's creation.  The drive
for features, growth of the team, and other organizational issues can cause a product's architecture
to take on a life of its own.  Developers new to a product often walk into an architecture that has
outgrown the simplicity of its earliest stages of development, leaving one to wonder:

- _"Just how is anyone on the team supposed to see the big picture?"_
- _"Is there even anyone \[left] on the team, who can?"_

If this sounds familiar to you, give Marmot a try.  It might just help.

## What it does

Once you tell Marmot which Git repositories you're working with, it gives you a way to categorize
them and then operate upon all the Git repositories in a category _as if they were a single Git
repository_.  As much as one can, at least; Marmot is not meant to replace `git submodules`.

Marmot happens to store this information–meta data about how Git repositories are related to each
other–in its own Git repository.  That's the author's idea of a Meta Repo, in a nutshell.
Developers' understanding of their environment is likely to change over time, leading to new ways of
categorizing code.  Might as well store that information in a Git repository, so you can experiment,
roll back, or even share this information with others.

## Examples

Once you can group Git repositories together and work with them as if they are a single unit, you
can do new things.  This section will try to show you some examples of what that could look like.

For the purpose of these examples, imagine a web application with a classic, 3-tiered architecture:
it has a front-end, a back-end, and a database.  Your organization has chosen to store those 3 kinds
of sources in separate Git repositories.

Let us also imagine that there is no talking your team into maybe–just maybe–putting all of those
sources–the sources that all have the same reason to change–into one multi-repo.  Wouldn't that help
developers avoid bugs at interface boundaries and DevOps-related folks avoid deploying things in the
wrong order?  If you're using Marmot, you're not on a team that feels comfortable trying that right
now.

## Tell Marmot about this code

Of course you will need to start by cloning all of those repositories:

```sh
# Assume the junk-drawer model; all code is in $HOME/git
git clone ssh://git@mycompany.com/app-ui $HOME/git/app-ui
git clone ssh://git@mycompany.com/app-server $HOME/git/app-server
git clone ssh://git@mycompany.com/app-database $HOME/git/app-database
```

If this is your first time using Marmot on this machine, then remember to:

```sh
# Creates a Meta Repo at $HOME/meta
marmot init
```

Then create a category in Marmot that relates these repositories to one another:

```sh
marmot category create project/my-app
marmot category add project/my-app \
  $HOME/git/app-ui \
  $HOME/git/app-server \
  $HOME/git/app-database
```

Marmot creates `$HOME/meta/project/my-app`, with symlinks back to the individual repositories:

```sh
$ ls -l $HOME/meta/project/my-app
app-ui -> $HOME/git/app-ui
app-server -> $HOME/git/app-server
app-database -> $HOME/git/app-database
```

Now that you have done a little bit of organization, you can now start to think of this code as a
single unit that has the same reason to change.

### Full-stack editing

The easiest way to start is to open up your favorite editor in the _category's directory_, instead
of opening an editor for each repository.  If nothing else, you can now do a project-wide find and
replace in a single window instead of in 3.

Once you have made changes to multiple repositories, you might even get away with emulating a
unified commit:

```sh
marmot exec --category <category> --repo-names heading \
  git add .
marmot exec --category <category> --repo-names heading \
  git commit -m "add new field to the application"

# commits changes to
# app-ui/widget-page - e.g. UI changes for the new entity
# app-server/widget-model - e.g. server changes for the new entity
# app-database/widget-table - e.g. database changes for the new entity
```

Or your editor may make it easy to do the same, using its own interface.

### Trace control- and data-flow

Marmot and `git grep` might help you trace control flow, such as:

- of a function call from an application to a library (look for the function name)
- of an HTTP request to its request handler (look for the request path), or
- of a stored procedure call from an application server to a database (look for the procedure name)

using something like this:

```sh
$ marmot exec --category project/my-app --repo-names heading \
  git --no-pager grep -i 'path[/]to[/]controller'
app-ui:
api-service.js: fetch('path/to/controller', /* response handler */)

app-server:
controller.js: app.get('path/to/controller', /* request handler */)
```

### Environment auditing

Have you ever wondered if different repositories are using different versions of the same thing?

```sh
$ marmot exec --category lang/typescript --repo-names heading \
  git --no-pager grep typescript 'package.json'
$HOME/new-repo:
package.json:    "typescript": "^5.2.2"

$HOME/old-repo:
package.json:    "typescript": "^4.1.6"
```

### Maybe try sharing?

Teams may also find it valuable to create a meta repo for their project and share it with each
other, but the primary audience for this tool is the individual developer.

Nothing in the underlying Git repositories has to change or even know about the existence of a meta
repo.  Individual developers can still create their own meta repo without needing support or buy-in
from anyone else.

---

## Contents

### For Users

- [Installation](#installation)
- [Getting Started](./doc/getting-started.md)
- [Commands](./doc/commands.md)
- [Environment Variables](./doc/environment-variables.md)

### For Developers

- [Architecture and Design Decisions](./doc/decisions.md)
- [CI/CD Jobs](./doc/cicd-jobs.md)
- [Task Automation](./doc/task-automation.md)
- [Tools](./doc/tools.md)

---

## Installation

Installs symlinks in `/usr/local/bin/`.

```sh
./src/marmot.zsh link
```

If you can run `marmot --help`, you have a working installation.  Head over to [Getting
Started](./doc/getting-started.md) to see if Marmot does anything that might be useful to you.

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
