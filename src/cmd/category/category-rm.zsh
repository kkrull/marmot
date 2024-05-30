#!/bin/zsh

emulate -LR zsh
set -euo pipefail

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"
source "$_MARMOT_HOME/lib/category-id.zsh"
source "$_MARMOT_HOME/lib/jq.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} rm"

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
    echo "$_MARMOT_INVOCATION: Missing category name"
    exit 1
  elif [[ $# == 1 ]]
  then
    echo "$_MARMOT_INVOCATION: Missing repository"
    exit 1
  fi

  local category_or_subcategory="$1" ; shift 1

  _config_rm_repositories_from_category \
    "$(_fs_metadata_file)" \
    "$category_or_subcategory" \
    "${@:#}"
  rm_link_to_category "$category_or_subcategory" "${@:#}"
}

function rm_link_to_category() {
  local category_name="$1" ; shift 1

  local link_path repository_path
  for repository_path in "${@:#}"
  do
    link_path="$(_fs_rm_repository_link "$category_name" "$repository_path")"
    echo "- ${link_path} (link)"
  done
}

function print_usage() {
  cat >&2 <<-EOF
USAGE
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION <category> <repository> [...]
$_MARMOT_INVOCATION <category>/<sub-category> <repository> [...]

OPTIONS
--help        Show help

See \`man ${_MARMOT_INVOCATION// /-}\` for details.
EOF
}

## Main

main "$@"; exit
