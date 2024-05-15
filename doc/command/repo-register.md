# `marmot repo register`

`marmot repo register` - Register repositories to manage

## SYNOPSIS

```sh
marmot repo register --help
marmot repo register <Git repository> ...
```

## DESCRIPTION

This command registers 1 or more repositories with Marmot, so it can manage them.

## OPTIONS

--help        Show help

## EXAMPLE

Register all the things!

```sh
find -s ~/git -type d -name .git \
  | sed 's/[/][.]git$//g' \
  | xargs -I {} marmot repo register {}
```
