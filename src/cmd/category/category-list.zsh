#!/bin/zsh

emulate -LR zsh

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} list"

## Command

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
$_MARMOT_INVOCATION - List categories

SYNOPSIS
$_MARMOT_INVOCATION [--help]
$_MARMOT_INVOCATION

DESCRIPTION
This command lists the categories that are used to group repositories.

OPTIONS
--help        Show help
EOF
}

## Main

main "$@"
