#!/bin/zsh

emulate -LR zsh

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"
source "$_MARMOT_HOME/lib/json.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} list"

## Command

function main() {
  if [[ $# == 0 ]]
  then
    list_local_repositories "$(_fs_metadata_file)"
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
  local config_file
  config_file="$1"
  shift 1

  _config_repository_paths "$config_file"
}

function print_usage() {
  cat >&2 <<-EOF
$_MARMOT_INVOCATION - List repositories

SYNOPSIS
$_MARMOT_INVOCATION [--help]
$_MARMOT_INVOCATION

DESCRIPTION
This command lists repositories that are managed by Marmot.

OPTIONS
--help        Show help
EOF
}

## Main

main "$@"
