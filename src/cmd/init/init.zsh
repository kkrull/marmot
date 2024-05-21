#!/bin/zsh

emulate -LR zsh

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"
source "$_MARMOT_HOME/lib/json.zsh"

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

  _config_init "$directory"
  echo "Initialized meta repository at $directory"
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

## Main

main "$@"; exit
