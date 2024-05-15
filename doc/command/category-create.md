# `marmot category create`

`marmot category create` - Create a category

## SYNOPSIS

```sh
marmot category create --help
marmot category create <category> [sub-category...]
```

## DESCRIPTION

This command creates a new category and adds its directory structure to the
meta repo.

## OPTIONS

```text
--help        Show help
```

## EXAMPLE

Create a "lang" category with sub-categories "java" and "typescript":

```sh
marmot category create lang java typescript
```

Create a "platform" category with sub-categories "beam, "clr", "jvm", and "node":

```sh
marmot category create platform beam clr jvm node
```

Create a "project" category with sub-categories "dotnet-8-migration" and "skunkworks":

```sh
marmot category create project dotnet-8-migration skunkworks
```
