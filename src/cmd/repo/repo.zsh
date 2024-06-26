#!/bin/zsh

emulate -LR zsh
set -euo pipefail
while IFS= read -d $'\0' -r f; do
  source "$f"
done < <(find "$_MARMOT_HOME/lib" -type f -iname '*.zsh' -print0 | sort -nz)

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} repo"

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
  'list')
    shift 1
    exec "$_MARMOT_HOME/cmd/repo/repo-list.zsh" "$@"
    ;;

  'prune')
    shift 1
    exec "$_MARMOT_HOME/cmd/repo/repo-prune.zsh" "$@"
    ;;

  'register')
    shift 1
    exec "$_MARMOT_HOME/cmd/repo/repo-register.zsh" "$@"
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
list          List repositories
prune         Prune references to missing repositories
register      Register repositories to manage

OPTIONS
--help        Show help

See \`man ${_MARMOT_INVOCATION// /-}\` for details.
EOF
}

## Main

main "$@"; exit
