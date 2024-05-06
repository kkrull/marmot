#!/bin/zsh

emulate -LR zsh

source "$_MARMOT_HOME/lib/config-file.zsh"
source "$_MARMOT_HOME/lib/json.zsh"
source "$_MARMOT_HOME/lib/paths.zsh"

## Command

self_invocation="marmot repo register"

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
    register_local_repositories "$(meta_repo_config_file)" "$@"
    exit 0
  fi
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

function register_local_repositories() {
  local config_file
  config_file="$1"
  shift 1

  local repositories
  repositories=$(to_marmot_repositories "$@")
  jq_update "$config_file" ".meta_repo_next.repositories += ${repositories}"
}

## Main

main "$@"
