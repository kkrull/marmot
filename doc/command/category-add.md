# `marmot category add`

`marmot category add` - Add repositories to a category

## SYNOPSIS

```sh
marmot category add --help
marmot category add <category> <repository> [repository...]
marmot category add <category>/<sub-category> <repository> [repository...]
```

## DESCRIPTION

This command adds 1 or more repositories to a (sub-)category.

## OPTIONS

```text
--help        Show help
```

## EXAMPLES

Add a repository to the "user" category:

```sh
marmot category add user ~/git/dotfiles
```

Add some repositories to the "skunkworks" project (lookout Dr. Light):

```sh
marmot category add project/skunkworks ~/git/robot-masters ~/git/skull-fortress
```
