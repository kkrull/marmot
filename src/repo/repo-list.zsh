#!/bin/zsh

emulate -LR zsh

source "$_MARMOT_HOME/lib/config-file.zsh"
source "$_MARMOT_HOME/lib/json.zsh"
source "$_MARMOT_HOME/lib/paths.zsh"

## Command

self_invocation="marmot repo list"

function main() {
  if [[ $# == 0 ]]
  then
    list_local_repositories "$(meta_repo_config_file)"
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

  repository_paths "$config_file"

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

## Main

main "$@"
