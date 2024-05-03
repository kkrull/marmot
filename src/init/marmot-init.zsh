#!/bin/zsh

emulate -LR zsh

self_invocation="marmot init"

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  fi

  echo "Hello world!"
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
