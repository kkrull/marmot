#!/bin/zsh

emulate -LR zsh
set -euo pipefail

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} prune"

## Command

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ $# -gt 0 ]]
  then
    echo "Unknown option: $1"
    exit 1
  fi

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  fi

  prune_repositories "$(_fs_metadata_file)"
}

function print_usage() {
  cat >&2 <<-EOF
USAGE
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION

OPTIONS
--help        Show help

See \`man ${_MARMOT_INVOCATION// /-}\` for details.
EOF
}

function prune_repositories() {
  local config_file
  config_file="$1"

  local all_paths stale_paths
  stale_paths=()
  _config_repository_paths_reply "$config_file" 'all_paths'
  for repo_path in "${all_paths[@]}"
  do
    echo "? $repo_path"
    [[ -d "$repo_path" ]] && continue
    stale_paths+=("$repo_path")
  done

  for repo_path in "${stale_paths[@]}"
  do
    echo "- $repo_path"
  done
}

## Main

main "$@"
