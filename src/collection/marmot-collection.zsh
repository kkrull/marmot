#!/bin/zsh

emulate -LR zsh

self_invocation="marmot collection"

#working_dirname="${PWD:A}"
#meta_repo_data="$working_dirname/.marmot"

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  fi

  echo "Done."
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - Work with collections

SYNOPSIS
${self_invocation} [--help]

OPTIONS
--help          Show help.
EOF
}

main "$@"; exit
