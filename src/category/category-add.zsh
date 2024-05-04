#!/bin/zsh

emulate -LR zsh

self_invocation="marmot category add"

#working_dirname="${PWD:A}"
#meta_repo_data="$working_dirname/.marmot"
#meta_repo_config="$meta_repo_data/meta-repo.json"

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

  echo "$name"
  for value in "$@"
  do
    echo "+$value"
  done
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - Add a category that will be used to slice repositories into clusters

SYNOPSIS
${self_invocation} <name> [...value]
  [--help]

DESCRIPTION
This command adds a category to the meta repo in the current working directory.

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
