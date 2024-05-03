#!/bin/zsh

emulate -LR zsh

#self_invocation="marmot init"

function main() {
  echo "Hello world!"
}

main "$@"; exit
