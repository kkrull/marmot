#!/bin/zsh

emulate -LR zsh

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

  *)
    echo "Unknown sub-command: $sub_command"
    exit 1
    ;;
  esac
}

function print_usage() {
  cat >&2 <<-EOF
$_MARMOT_INVOCATION - Work with categories

SYNOPSIS
$_MARMOT_INVOCATION [--help]
$_MARMOT_INVOCATION sub-command [options...]

SUB-COMMANDS
add           Add repositories to a category
create        Create a new category
list          List categories

OPTIONS
--help        Show help
EOF
}

## Main

main "$@"; exit
