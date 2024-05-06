#!/bin/zsh

emulate -LR zsh

self_dirname="${0:A:h}"
self_invocation="marmot repo"

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
    exec "${self_dirname}/repo-list.zsh" "$@"
    ;;

  *)
    echo "Unknown sub-command: $sub_command"
    exit 1
    ;;
  esac
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - Work with repositories

SYNOPSIS
${self_invocation} [--help]

SUB-COMMANDS
list          List repositories

OPTIONS
--help        Show help
EOF
}

main "$@"; exit
