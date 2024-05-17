#!/bin/zsh

emulate -LR zsh
set -euo pipefail

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} prune"

## Command

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ $# -gt 0 ]]
  then
    echo "Unknown option: $1"
    exit 1
  fi

  if [[ -n "$help_option" ]]
  then
    print_usage
  else
    _config_repository_paths "$(_fs_metadata_file)"
  fi
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
