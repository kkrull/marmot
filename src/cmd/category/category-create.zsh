#!/bin/zsh

emulate -LR zsh
set -euo pipefail
while IFS= read -d $'\0' -r f; do
  source "$f"
done < <(find -s "$_MARMOT_HOME/lib" -type f -iname '*.zsh' -print0)

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
  elif [[ $# == 0 || -z "$1" ]]
  then
    echo "$_MARMOT_INVOCATION: Missing category name"
    exit 1
  else
    _categorycmd_create "$@"
  fi
}

function print_usage() {
  cat >&2 <<-EOF
USAGE
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION <category> [sub-category ...]

OPTIONS
--help        Show help

See \`man ${_MARMOT_INVOCATION// /-}\` for details.
EOF
}

## Main

main "$@"; exit
