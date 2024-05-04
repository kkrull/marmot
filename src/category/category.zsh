#!/bin/zsh

emulate -LR zsh

self_invocation="marmot category"

working_dirname="${PWD:A}"
meta_repo_data="$working_dirname/.marmot"
meta_repo_config="$meta_repo_data/meta-repo.json"

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ $# == 0 ]]
  then
    print_usage
    exit 0
  elif [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  fi

  sub_command="$1"
  case "$sub_command" in
  'list')
    list_categories
    ;;

  *)
    echo "Unknown sub-command: $sub_command"
    exit 1
    ;;
  esac
}

function list_categories() {
  jq < "$meta_repo_config" \
    -r '.meta_repo.categories[].name'
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
