# Getting Started

Please make sure you have completed the [installation instructions](../README.md#installation),
first.

Now prepare to use your imagination.

For the purpose of these examples, imagine a web application with a classic, 3-tiered architecture:
it has a front-end, a back-end, and a database.  Your organization has chosen to store those 3 kinds
of sources in separate Git repositories.

Let us also imagine that there is no talking your team into maybe–just maybe–putting all of those
sources in one place (e.g. a multi- or mono-repo).  Wouldn't it make sense to put sources together
that have the same reason to change?  Wouldn't that help developers avoid bugs at interface
boundaries and DevOps-related folks avoid deploying things in the wrong order?

If you're using Marmot, you're not on a team that feels comfortable trying that right now.

## Tell Marmot about this code

Of course you will need to start by cloning all of those repositories:

```sh
# Assume the junk-drawer model; all code is in $HOME/git
git clone ssh://git@mycompany.com/app-ui $HOME/git/app-ui
git clone ssh://git@mycompany.com/app-server $HOME/git/app-server
git clone ssh://git@mycompany.com/app-database $HOME/git/app-database
```

If this is your first time using Marmot on this machine, please remember to:

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
$HOME/git/app-ui:
api-service.js: fetch('path/to/controller', /* response handler */)

$HOME/git/app-server:
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
