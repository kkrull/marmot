% MARMOT(7) Version 0.6 | Meta Repo Management Tool
% Kyle Krull
% May 2024

# NAME

**marmot** - Meta Repo Management Tool

# DESCRIPTION

**marmot** creates and maintains a Meta Repository (e.g. a "Meta Repo") of several Git repositories.

A Meta Repo, for the purposes of this program, is a set of references to other Git repositories that
are grouped into 1 or more categories.  **marmot** helps you work with all the repositories in a
category, as if they are a single unit.

## What it Does

**marmot** creates a directory structure in the meta repo's file system to mirror the way that
repositories have been categorized, so that there is a `/:category/:sub-category` directory for each
(sub-)category.  Each directory contains symbolic links back to the Git repositories that are
grouped into the same (sub-)category.

Users run commands from one of these directories in order to restrict commands to the Git
repositories that have that category in common.  This can be done directly by passing a category's
path to a shell command or indirectly by passing a shell command and the name of the category to
**marmot exec**.  Either way causes a command to run within the scope of a category, instead of
getting distracted by irrelevant sources in unrelated repositories.

# GETTING STARTED

## Initialize a Meta Repo

Before using **marmot**, you will need to:

```sh
$ marmot init
Initialized meta repository at ~/meta/.marmot
```

Now you have a Meta Repo, but it will not be very useful until you register some Git repositories
and categorize them in some way.

## Register Repositories

Let us assume that you have cloned some Git repositories and put them in `$HOME/git`.  Note however
that they can be on any reachable path.

```sh
marmot repo register ~/git/app-ui ~/git/app-server ~/git/app-database
```

Registering a Git repository simply means that **marmot** stores a reference (e.g. a path) to it.
You can now execute shell commands on every registered repository with **marmot exec**.

This is useful for a small number of repositories, but categorizing them allows for finer-grained
control.

## Categorize Repositories

Categorizing Git repositories in **marmot** allows you to work on subsets of all the Git
repositories you have on your machine.  Each Git repository can be added to 0, 1, or multiple
categories.

A category is simply a 1- or 2-tiered tag; `:category` or `:category/:sub-category`, respectively.
The categories you create and the way you use them are up to you.

Creating a new category creates a directory structure in the Meta Repo:

```sh
$ marmot category create project frontend services
+ ~/meta/project (category)
+ ~/meta/project/frontend (sub-category)
+ ~/meta/project/services (sub-category)
```

Adding a repository to a category creates a symlink to it, in the Meta Repo:

```sh
# Group together front-end components for the same project
$ marmot category add project/frontend ~/git/webclient ~/git/webconfig
+ project/frontend/webclient (link)
+ project/frontend/webconfig (link)
```

```sh
# Group together back-end components for the same project
$ marmot category add project/services ~/git/fooservice ~/git/barservice
+ project/services/fooservice (link)
+ project/services/barservice (link)
```

In the above example, both microservices are now grouped together in `~/meta/project/services`:

```sh
$ ls -l ~/meta/project/services
lrwxr-xr-x somebody  users  ... barservice -> ~/git/barservice
lrwxr-xr-x somebody  users  ... fooservice -> ~/git/fooservice
```

# EXAMPLE

Note that examples list directories using `~/` instead of the path to an imaginary home directory,
for brevity.

## Use Categories for Editing

Sometimes it can be helpful to have all the code that is strongly-related by a common reason to
change open in a single editor.  Imagine how opening repositories for a front-end app and its
feature flags in a single editor might help you spot typos and stale flags:

```sh
code ~/meta/project/frontend
```

## Use Categories for Shell Commands

Sometimes it can be helpful to run the a shell command on several–but not all–repositories at once.
Imagine how categorizing repositories by platform might make it easier to check if they are all
using the same version of the platform:

```sh
# direnv switches node to use the version listed in each directory
$ marmot exec --category platform/node cat .node-version
~/meta/platform/node/fooclient: v16.0
~/meta/platform/node/barclient: v18.0
```

## Search Within a Category

Sometimes it can be helpful to search for the same text in several places.

```sh
# Are all API clients using the same version?
$ marmot exec --category api/client git grep 'apiVersion'
~/meta/api/client/fooclient/api.js: apiVersion: "1.0"
~/meta/api/client/barclient/api.js: apiVersion: "2.0"
```

```sh
# Are both sides of an API using the same names for things?
$ marmot exec --category webapp git grep -e 'some[_]?field'
~/meta/webapp/fooserver/controller.js: some_field: "42"
~/meta/webapp/barclient/api.js: const answer = response.someField

```

# SEE ALSO

[*marmot(1)*](./marmot.1.md), [*marmot-category(1)*](./marmot-category.1.md),
[*marmot-exec(1)*](./marmot-exec.1.md), [*marmot-init(1)*](./marmot-init.1.md),
[*marmot-repo(1)*](./marmot-repo.1.md)
