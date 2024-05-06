#!/bin/zsh

emulate -LR zsh

self_invocation="marmot repo list"

#working_dirname="${PWD:A}"
#meta_repo_data="$working_dirname/.marmot"
#meta_repo_config="$meta_repo_data/meta-repo.json"

function main() {
  if [[ $# == 0 ]]
  then
    list_local_repositories
    exit 0
  fi

  zparseopts -D -E \
    -help=help_option
  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  else
    echo "Unknown option: $1"
    exit 1
  fi
}

function list_local_repositories() {
  find -s "$HOME/git" -name '.git' -print -type d \
    | sed 's/[/][.]git$//g'
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - List repositories

SYNOPSIS
${self_invocation}
${self_invocation} [--help]
${self_invocation} [--host=github] <Host URL>

DESCRIPTION
This command lists repositories that exist on the specified host.

OPTIONS
--help        Show help
--host        The type of host.  One of: github
EOF
}

main "$@"
