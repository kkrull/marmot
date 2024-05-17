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
  declare config_file="$1"

  declare reply stale_paths=()
  _config_repository_paths_reply "$config_file"
  for repo_path in "${reply[@]}"
  do
    [[ ! -d "$repo_path" ]] && stale_paths+=("$repo_path")
  done

  _config_remove_repositories "$config_file" "${stale_paths[@]}"
}

## Main

main "$@"
