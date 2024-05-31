#!/bin/zsh

emulate -LR zsh
set -euo pipefail
while IFS= read -d $'\0' -r f; do
  source "$f"
done < <(find "$_MARMOT_HOME/lib" -type f -iname '*.zsh' -print0 | sort -nz)

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} category"

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
  'add')
    shift 1
    exec "$_MARMOT_HOME/cmd/category/category-add.zsh" "$@"
    ;;

  'create')
    shift 1
    exec "$_MARMOT_HOME/cmd/category/category-create.zsh" "$@"
    ;;

  'list')
    shift 1
    exec "$_MARMOT_HOME/cmd/category/category-list.zsh" "$@"
    ;;

  'rm')
    shift 1
    exec "$_MARMOT_HOME/cmd/category/category-rm.zsh" "$@"
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
add           Add repositories to a category
create        Create a new category
list          List categories
rm            Remove repositories from a category

OPTIONS
--help        Show help

See \`man ${_MARMOT_INVOCATION// /-}\` for details.
EOF
}

## Main

main "$@"; exit
