#!/bin/zsh

emulate -LR zsh

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} list"

## Command

function main() {
  if [[ $# == 0 ]]
  then
    _config_category_names "$(_fs_metadata_file)"
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

function print_usage() {
  cat >&2 <<-EOF
$_MARMOT_INVOCATION - List categories

SYNOPSIS
$_MARMOT_INVOCATION [--help]
$_MARMOT_INVOCATION

DESCRIPTION
This command lists the categories that are used to group repositories.

OPTIONS
--help        Show help
EOF
}

## Main

main "$@"
