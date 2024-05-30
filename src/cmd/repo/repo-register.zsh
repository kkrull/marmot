#!/bin/zsh

emulate -LR zsh
set -euo pipefail
while IFS= read -d $'\0' -r f; do
  source "$f"
done < <(find -s "$_MARMOT_HOME/lib" -type f -iname '*.zsh' -print0)

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} register"

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
    echo "$_MARMOT_INVOCATION: Missing repository path"
    exit 1
  else
    register_local_repositories "$(_fs_metadata_file)" "$@"
    exit 0
  fi
}

function print_usage() {
  cat >&2 <<-EOF
USAGE
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION <repository path> ...

OPTIONS
--help        Show help

See \`man ${_MARMOT_INVOCATION// /-}\` for details.
EOF
}

function register_local_repositories() {
  local config_file
  config_file="$1"
  shift 1

  _repomd_add_local_paths "$config_file" "$@"
}

## Main

main "$@"
