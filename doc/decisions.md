# Decision Log

A brief description of some of the major decisions along the way.

## 00: Starting conditions

With a growing network of Git repositories that are hosted in a variety of locations, I often find
myself wanting to run `git grep` on several repositories at once.

Could there be a way to arbitrarily group, query, and operate upon several repositories at once?
For example, a `language` category might have sub-categories like `java` and `typescript`, or a
`project` category might have sub-categories for each project you have worked on.

Could there be a way to build my own logical structure of somebody else's code, without having to
talk them into condensing their code into a multi-repo?  Could that help manage disjointed code?

Building a meta repo could help, but I would need a tool to maintain it.  Let's call it the Meta
Repo Management Tool, or "marmot" for short.

### Decisions

- Create a tool that tags Git repositories with categories and sub-categories and lets me run.
  shell commands for a category of repositories as if they are all part of a single unit.
- Store meta data about categories externally, instead of in the Git repositories themselves.

## 01: Target Z Shell

Implement marmot in *nix tools that are widely-available on the platforms I use - e.g. MacOS, Linux,
and Windows Subsystem for Linux.  Writing it in Z Shell can make it easier to try new ideas, while
avoiding the need to port or re-build for other platforms.  Breaking the scripts up into commands
and sub-commands (ala Git) can help keep the size of each script manageable.

### Decisions

- marmot is a program with a command line interface, using a Git-like command/sub-command style.
- Write marmot with Z Shell scripts.
- Delegate to tools from commonly-available packages in MacOS and Linux.

## 02: Store metadata in JSON files

I don't know exactly what kind of meta data marmot will need to store about Git repositories, aside
from categories and paths to repositories.  Storing this data in JSON files can make it possible to
extend with new fields, fix by hand when necessary, and query with tools like `jq`.

Categorizations are other meta data may grow over time, as I learn more about the Git repositories I
am using.  Maybe it might even grow into some sort of neural network of Git repositories?  Storing
the Meta Repo's contents in a separate Git repository would make it possible to track changes, roll
back, or even share with teammates.

### Decisions

- Store meta data in JSON files.
- Use tools like `jq` and `jo` to query and construct JSON data from marmot.
- Store meta data in its own Git repository.

## 03: Directory Structure in the Meta Repo

Sometimes I need to search in several Git repositories that use the same programming language. Other
times, I do full-stack development in all the languages used for a product.  Each repository can be
in more than 1 category, so there is no, single directory structure that works in every case.
Creating a directory structure of categories that link back to the Git repositories can help manage
this complexity.  Version managers for Node.js and Ruby come to mind, as sources of inspiration.

Git repositories still have to be cloned somewhere, though.  The host/repository structure of Golang
(e.g. `go get ...`) comes to mind, as a way to avoid name collisions.

### Decisions

- Clone Git repositories in a common location, separated by host - e.g. `$HOME/git/:host/:name`.
- Create directories in the Meta Repo for each (sub-)category - e.g.
  `$HOME/meta/:category[/:sub-category]`.
- Create symbolic links in (sub-)category directories that link to the underlying Git repositories.

## 04: Use Semantic Versioning

Use a semantic versioning system with fairly objective criteria, to avoid prolonged deliberation
over what changes merit what kind of version bump.

### Decisions

- Major version: Increment from 0.x to 1.x when there are enough features to be useful.
- Minor version: Increment when adding a new feature (e.g. a command or sub-command).
- Patch version: Increment when refactoring to prepare for another feature.

## 05: Apply Single Responsibility Principle to scripts

Scripts are getting more complex, leading to duplication of concepts and algorithms.  Applying the
Single Responsibility Principle (SRP) can help manage complexity and avoid unnecessary duplication.
This may drive adoption of other SOLID principles, as well.

### Decisions

#### Function structure

- Write shared code as functions, using a functional style (e.g. command-query separation).
- Commands: Make separate functions for separate side-effects.
- Queries:
  - Pass data in as parameters, using quotes for any variables that may contain whitespace.
  - Pass arrays as `"${my_array[@]}"` so the whole array is passed instead of just the first word.
  - Try returning data via `echo` or `printf`, at first.  This incurs a performance penalty of the
    call site having to fork a sub-shell, but this is not expected to be a concern in practice.
  - If queries must be invoked without starting a sub-shell, environment variables `REPLY` and
    `reply` may be used to return conventional data and arrays, respectively.

Source: <https://unix.stackexchange.com/a/365417/37734>

#### Location of shared code

- Put shared code in `src/lib/`.
- Gather together shared functions that operate on the same bounded context (e.g. the same data).
  Explore a convention of making that bounded context the first parameter in each function.
- Name files according to their bounded context.

#### Using shared code

- Use `_MARMOT_HOME` set in the top level `marmot.zsh` script to locate shared code scripts.
- Source code from (sub-)command scripts (e.g. the script used to start the process), ala Rails.
  - Some code in `lib/` may depend upon other code in `lib/`, but it is up to the top-level script
    to `source` dependencies and transitive dependencies.
  - This is approach is intended to avoid any complexities in the same code being sourced twice.  I
    have no idea what could happen then, and I'd rather not have to find out.
