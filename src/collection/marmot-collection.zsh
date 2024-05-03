#!/bin/zsh

emulate -LR zsh

self_invocation="marmot collection"

working_dirname="${PWD:A}"
meta_repo_data="$working_dirname/.marmot"
meta_repo_config="$meta_repo_data/meta-repo.json"

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  fi

  list_collections
}

function list_collections() {
  jq < "$meta_repo_config" \
    -r '.meta_repo.collection_types[].name'
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - Work with collections

SYNOPSIS
${self_invocation} [--help]

SUB-COMMANDS
list      List collections

OPTIONS
--help          Show help.
EOF
}

main "$@"; exit
