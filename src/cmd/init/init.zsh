#!/bin/zsh

emulate -LR zsh

source "$_MARMOT_HOME/lib/paths.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} init"

## Command

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  elif [[ -d "$(meta_repo_data)" ]]
  then
    printf "Meta repo already exists: %s" "$(meta_repo_data)"
    exit 1
  else
    create_meta_repo "$(meta_repo_data)"
  fi
}

function create_meta_repo() {
  local directory="$1"
  mkdir -p "$directory"

  local template
  template=$(cat <<'EOF'
{
  meta_repo: {
    categories: []
  },
  version: $version
}
EOF
  )

  local meta_repo_file="$directory/meta-repo.json"
  jq > "$meta_repo_file" \
    --null-input \
    --arg version 0.3.1 \
    "$template"

  echo "Initialized meta repository at $(meta_repo_home)"
}

function print_usage() {
  cat >&2 <<-EOF
$_MARMOT_INVOCATION - Initialize a meta repo

SYNOPSIS
$_MARMOT_INVOCATION [--help]
$_MARMOT_INVOCATION

DESCRIPTION
This command initializes a blank meta repo in the current working directory.

OPTIONS
--help        Show help
EOF
}

## Main

main "$@"; exit
