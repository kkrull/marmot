#!/bin/zsh

emulate -LR zsh
set -euo pipefail

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} meta"

## Command

function main() {
  if [[ $# == 0 ]]
  then
    print_usage
    exit 0
  fi

  zparseopts -E \
    -help=help_option
  if [[ $# == 1 && -n "$help_option" ]]
  then
    print_usage
    exit 0
  fi

  sub_command="$1"
  case "$sub_command" in
  'home')
    shift 1
    exec "$_MARMOT_HOME/cmd/meta/meta-home.zsh" "$@"
    ;;

  *)
    echo "Unknown sub-command: $sub_command"
    exit 1
    ;;
  esac
}

function print_usage() {
  cat >&2 <<-EOF
USAGE
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION sub-command [args ...]

SUB-COMMANDS
home          Show the base directory of the Meta Repo

OPTIONS
--help        Show help

See \`man ${_MARMOT_INVOCATION// /-}\` for details.
EOF
}

## Main

main "$@"; exit
