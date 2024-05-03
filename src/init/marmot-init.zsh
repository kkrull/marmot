#!/bin/zsh

emulate -LR zsh

self_invocation="marmot init"
working_dirname="${PWD:A}"

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  fi

  create_meta_repo
}

function create_meta_repo() {
  local meta_repo_data="$working_dirname/.marmot"
  mkdir -p "$meta_repo_data"

  local meta_repo_file="$meta_repo_data/meta-repo.json"
  echo "Hello World!" > "$meta_repo_file"

  echo "Initialized meta repository at $working_dirname"
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - Initialize a meta repository in the current directory

SYNOPSIS
${self_invocation} [--help]

OPTIONS
--help          Show help.
EOF
}

main "$@"; exit
