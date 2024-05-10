# `marmot repo list`

`marmot repo list` - List repositories

## SYNOPSIS

```sh
marmot repo list --help
marmot repo list [--category <category|sub-category>]
```

## DESCRIPTION

This command lists repositories that have been registered with Marmot.
Given options, this lists only the repositories that match the given criteria.

## OPTIONS

```text
--category    List repositories that have been added to the given category
              or sub-category.
--help        Show help
```

## EXAMPLES

List registered TypeScript repositories

```sh
marmot repo list --category lang/typescript
```
