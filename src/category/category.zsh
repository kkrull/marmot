#!/bin/zsh

emulate -LR zsh

self_dirname="${0:A:h}"
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
    exec "${self_dirname}/category-add.zsh" "$@"
    ;;

  'list')
    shift 1
    exec "${self_dirname}/category-list.zsh" "$@"
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
add           Add a category
list          List categories

OPTIONS
--help        Show help
EOF
}

main "$@"; exit
