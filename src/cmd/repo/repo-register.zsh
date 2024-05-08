#!/bin/zsh

emulate -LR zsh

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"
source "$_MARMOT_HOME/lib/json.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} register"

## Command

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  elif [[ $# == 0 ]]
  then
    echo "$_MARMOT_INVOCATION: Missing repository path"
    exit 1
  else
    register_local_repositories "$(_fs_metadata_file)" "$@"
    exit 0
  fi
}

function print_usage() {
  cat >&2 <<-EOF
$_MARMOT_INVOCATION - Register repositories to manage

SYNOPSIS
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION <Git repository> ...

DESCRIPTION
This command registers 1 or more repositories with Marmot, so it can manage them.

OPTIONS
--help        Show help

EXAMPLES
• Register all the things!
  \$ find -s ~/git -type d -name .git \\
    | sed 's/[/][.]git$//g' \\
    | xargs -I {} $_MARMOT_INVOCATION {}
EOF
}

function register_local_repositories() {
  local config_file
  config_file="$1"
  shift 1

  _config_add_repositories "$config_file" "$@"
}

## Main

main "$@"
