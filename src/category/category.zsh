#!/bin/zsh

emulate -LR zsh

## Command

self_invocation="marmot category"

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
    exec "$_MARMOT_HOME/category/category-add.zsh" "$@"
    ;;

  'create')
    shift 1
    exec "$_MARMOT_HOME/category/category-create.zsh" "$@"
    ;;

  'list')
    shift 1
    exec "$_MARMOT_HOME/category/category-list.zsh" "$@"
    ;;

  *)
    echo "Unknown sub-command: $sub_command"
    exit 1
    ;;
  esac
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - Work with categories

SYNOPSIS
${self_invocation} [--help]

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
