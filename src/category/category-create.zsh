#!/bin/zsh

emulate -LR zsh

self_invocation="marmot category create"
working_dirname="${PWD:A}"

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  elif [[ $# == 0 ]]
  then
    echo "$self_invocation: Missing category name"
    exit 1
  else
    make_category_directories "$@"
  fi
}

function make_category_directories() {
  local name
  name="$1"
  shift 1

  local category_path
  category_path="$working_dirname/$name"
  echo "+ $category_path (category)"
  mkdir -p "$category_path"

  local value_path
  for value in "$@"
  do
    value_path="$category_path/$value"
    echo "+ $value_path (category value)"
    mkdir -p "$value_path"
  done
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - Create a category

SYNOPSIS
${self_invocation} <name> [...value]
  [--help]

DESCRIPTION
This command creates a new category and adds its directory structure to the
meta repo in the current working directory.

OPTIONS
--help        Show help

EXAMPLES
• Create a "lang" category with possible values "java" and "typescript":
    \$ ${self_invocation} lang java typescript
• Create a "platform" category with possible values "beam, "clr", "jvm", and "node":
    \$ ${self_invocation} platform beam clr jvm node
• Create a "project" category with possible values "dotnet-8-migration" and "skunkworks":
    \$ ${self_invocation} project dotnet-8-migration skunkworks
EOF
}

main "$@"; exit
