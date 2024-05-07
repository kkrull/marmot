#!/bin/zsh

emulate -LR zsh

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"

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
  elif [[ -d "$(_fs_metadata_dir)" ]]
  then
    printf "Meta repo already exists: %s" "$(_fs_metadata_dir)"
    exit 1
  else
    create_meta_repo "$(_fs_metadata_dir)"
  fi
}

function create_meta_repo() {
  local directory="$1"
  mkdir -p "$directory"

  _config_metadata_init "$directory"
  echo "Initialized meta repository at $directory"
}

function print_usage() {
  cat >&2 <<-EOF
$_MARMOT_INVOCATION - Initialize a meta repo

SYNOPSIS
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION

DESCRIPTION
This command initializes a blank meta repo in the current working directory.

OPTIONS
--help        Show help
EOF
}

## Main

main "$@"; exit
