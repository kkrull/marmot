#!/bin/zsh

emulate -LR zsh

source "$_MARMOT_HOME/lib/paths.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} create"

## Command

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  elif [[ $# == 0 ]]
  then
    echo "$_MARMOT_INVOCATION: Missing category name"
    exit 1
  else
    make_category_directories "$@"
  fi
}

function make_category_directories() {
  local category_name
  category_name="$1"
  shift 1

#TODO KDK: Extract functions for the (sub-)category path
  local category_path
  category_path="$(meta_repo_home)/$category_name"
  echo "+ $category_path (category)"
  mkdir -p "$category_path"

  local subcategory_path
  for subcategory_name in "$@"
  do
    subcategory_path="$category_path/$subcategory_name"
    echo "+ $subcategory_path (sub-category)"
    mkdir -p "$subcategory_path"
  done
}

function print_usage() {
  cat >&2 <<-EOF
$_MARMOT_INVOCATION - Create a category

SYNOPSIS
$_MARMOT_INVOCATION [--help]
$_MARMOT_INVOCATION <category> [sub-category...]

DESCRIPTION
This command creates a new category and adds its directory structure to the
meta repo.

OPTIONS
--help        Show help

EXAMPLES
• Create a "lang" category with sub-categories "java" and "typescript":
    \$ $_MARMOT_INVOCATION lang java typescript
• Create a "platform" category with sub-categories "beam, "clr", "jvm", and "node":
    \$ $_MARMOT_INVOCATION platform beam clr jvm node
• Create a "project" category with sub-categories "dotnet-8-migration" and "skunkworks":
    \$ $_MARMOT_INVOCATION project dotnet-8-migration skunkworks
EOF
}

## Main

main "$@"; exit