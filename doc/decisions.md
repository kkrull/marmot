# Decision Log

A brief description of some of the major decisions along the way.

## 00: Starting conditions

With a growing network of Git repositories that are hosted in a variety of locations, I am finding
it more and more difficult to find something I remember doing years ago.  I often find myself
wanting to run `git grep` on several repositories at once.  After seeing how tagging Markdown
documents in Obsidian Notes helped me to organize information, I wondered if something similar could
be done for Git repositories.

Could I come up with a way to arbitrarily group, query, and operate upon several repositories at
once?  I'll call this a "category" for now, each of which can have 1 or more "values".  For example,
a `language` category might have values like `java` and `typescript`, or a `project` category might
have values for each project you have worked on.

Could I find a way to build my own logical structure of somebody else's code and find a way to cope
with their sprawling architecture, without having to talk them into condensing their code into
multi-repos?

Building my own meta repo could help, but I would need a tool to maintain it.  Let's call it the
Meta Repo Management Tool, or "marmot" for short.

## 01: Use widely-available *nix tools

Implement marmot in *nix tools that are widely-available on the platforms I use - e.g. MacOS, Linux,
and Windows Subsystem for Linux.  Writing it in ZShell and using readily-available packages should
make it easy to try new ideas on all these platforms, without much porting or re-building.  Plus I
have been tinkering with a lot of these ideas on the command-line already, so they won't need much
translation to put them into a script.

Since we're doing scripting, organize it into command-based scripts like Git to keep the size and
scope of responsibility of each file manageable.

## 02: Store metadata as JSON files

If marmot is going to build a neural network of information about Git repositories, it will need to
be extendible.  It would also be helpful to have something that's in plain text, in case I need to
make some sort of change to metadata that marmot doesn't support yet.  Storing JSON files in the
Meta Repo's directory should satisfy both, while also offering the ability to version control (e.g.
roll back mistakes) the metadata itself.

If the Meta Repo itself is also a Git repository, that might make it convenient to clone the Meta
Repo when I move to another machine.  Otherwise I would have to create it all over again, likely
with a different set of Git repositories, and then it's not really a Meta Repo anymore.

Try using standard tools like `jq` and `jo` for working with JSON files, so I don't have to do any
parsing myself (or resort to another language, just to get access to a JSON library).

## 03: Structure of the Meta Repo

The Meta Repo is an organized network of symlinks to the underlying Git repositories.  It's kind of
like how Node Version Manager and `rbenv` build symlinks to whichever version of `node` or `ruby`
are configured for your project.

First, put all Git repositories within reach in one place like `~/git/`.  Sub-divide that directory
by host name to avoid name collisions, much like `go get` does with Golang 1.x.  Git repositories
therefore exist at `~/git/:host/:repository/`.

Next, build a secondary structure at `~/meta/` that has sub-directories for each category and
symlinks for each repository belonging to a category.  This results in a path
`~/meta/:category-type/:category-value/:repository/`, where `:repository` is a symlink back to
wherever the Git repository is actually stored in `~/git/:host/`.

Using directories for categories should make it possible to scope command-line tools to the
repositories in the same category (e.g. `find ~/meta/... -exec ...`), while also making it possible
to open up editors on all related repositories at the same time.

Using symlinks allows for a tag-like system where each repository can be tagged or categorized 0..n
ways without duplicating data.  For example, `~/meta/lang/java/greeter-java/` and
`~/meta/kata/greeter/greeter-java/` could both point to the same repository located at
`~/git/:host/greeter-java/`.
