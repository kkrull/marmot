#!/bin/zsh

emulate -LR zsh

source "$_MARMOT_HOME/lib/fs.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} home"

## Command

function main() {
  if [[ $# == 0 ]]
  then
    show_home
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

function show_home() {
  _fs_metarepo_home
}

function print_usage() {
  cat >&2 <<-EOF
$_MARMOT_INVOCATION - Show path to Meta Repo

SYNOPSIS
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION

DESCRIPTION
This command shows the base directory of the Meta Repo.

OPTIONS
--help        Show help
EOF
}

## Main

main "$@"
