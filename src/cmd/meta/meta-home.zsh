#!/bin/zsh

emulate -LR zsh
set -euo pipefail

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} home"

## Local environment

source "$_MARMOT_HOME/lib/fs.zsh"

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
USAGE
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION

OPTIONS
--help        Show help

See \`man ${_MARMOT_INVOCATION// /-}\` for details.
EOF
}

## Main

main "$@"
