#!/bin/zsh

emulate -LR zsh

self_invocation="marmot repo list"

working_dirname="${PWD:A}"
meta_repo_data="$working_dirname/.marmot"
meta_repo_config="$meta_repo_data/meta-repo.json"

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
  # Treat lack of JSON fields as empty rather than as an error
  # https://github.com/jqlang/jq/issues/354#issuecomment-43147898
  jq < "$meta_repo_config" \
    -r '.meta_repo.repositories[]?.name'
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - List repositories

SYNOPSIS
${self_invocation}
${self_invocation} [--help]

DESCRIPTION
This command lists repositories that exist on the specified host.

OPTIONS
--help        Show help
EOF
}

main "$@"