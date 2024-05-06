#!/bin/zsh

emulate -LR zsh

self_invocation="marmot repo register"

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
  elif [[ $# == 0 ]]
  then
    echo "$self_invocation: Missing repository path"
    exit 1
  else
    register_local_repositories "$@"
    exit 0
  fi
}

function register_local_repositories() {
  # TODO KDK: jq < ~/meta/.marmot/meta-repo.json '.meta_repo.repositories |= ["$HOME/git"]'
  jq < "$meta_repo_config" \
    '.meta_repo.repositories?[] += ["a"]'
  echo "Done."
  # jq < "$meta_repo_config" \
  #   -r '.meta_repo.repositories[]?.name'
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - Register repositories to manage

SYNOPSIS
${self_invocation} <Git repository> ...
${self_invocation} [--help]

DESCRIPTION
This command registers 1 or more repositories with Marmot, so it can manage them.

OPTIONS
--help        Show help
EOF
}

main "$@"
