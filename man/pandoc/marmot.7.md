% MARMOT(7) Version 0.5 | Meta Repo Management Tool
% Kyle Krull
% May 2024

# WHAT IT DOES

Marmot creates and maintains a Meta Repository (e.g. a "meta repo"), which can
be used to group several Git repositories by 1 or more arbitrary categories.

Marmot creates a directory structure in the meta repo's file system to mirror
the way that repositories have been categorized, so that there is a
`/:category/:sub-category` directory for each (sub-)category.  Each directory
contains symbolic links back to the Git repositories that are grouped into
the same (sub-)category.

Users run commands from one of these directories in order to restrict
commands to the Git repositories that have that categorization in common.
In this fashion, users can do things like search closely-related
code with `git grep` or open an editor for those Git repositories, without
clutter and noise from irrelevant sources in unrelated repositories.
