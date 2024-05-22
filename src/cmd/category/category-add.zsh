#!/bin/zsh

emulate -LR zsh
set -euo pipefail

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"
source "$_MARMOT_HOME/lib/id.zsh"
source "$_MARMOT_HOME/lib/json.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} add"

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

  local category_or_subcategory category_name subcategory_name
  category_or_subcategory="$1"
  category_name="$(_id_category_name "$category_or_subcategory")"
  subcategory_name="$(_id_subcategory_name "$category_or_subcategory")"

  # Given (sub-)category may be new; create if so
  _config_add_categories "$(_fs_metadata_file)" "$category_name" "$subcategory_name"
  _fs_make_category_path "$category_name"
  _fs_make_subcategory_path "$category_name" "$subcategory_name"

  # Add these repositories to the (sub-)category
  _config_add_repositories_to_category \
    "$(_fs_metadata_file)" \
    "$category_or_subcategory" \
    "${@:2}"
  link_to_category "$category_or_subcategory" "${@:2}"
}

function link_to_category() {
  local category_name
  category_name="$1"

  local link_path
  local repository_path
  for repository_path in "${@:2}"
  do
    link_path="$(_fs_add_repository_link "$category_name" "$repository_path")"
    echo "+ ${link_path} (link)"
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
