#!/bin/zsh

emulate -LR zsh

self_invocation="marmot category add"

working_dirname="${PWD:A}"
meta_repo_data="$working_dirname/.marmot"
meta_repo_config="$meta_repo_data/meta-repo.json"

function main() {
  if [[ $# == 0 ]]
  then
    list_categories
    exit 0
  fi

  zparseopts -D -E \
    -help=help_option
  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  else
    echo "Unknown option: $1"
    exit 1
  fi
}

function list_categories() {
  jq < "$meta_repo_config" \
    -r '.meta_repo.categories[].name'
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - Add a category that will be used to slice repositories into clusters

SYNOPSIS
${self_invocation} <name> [...value]
  [--help]

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

main "$@"
