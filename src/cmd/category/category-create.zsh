#!/bin/zsh

emulate -LR zsh

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"
source "$_MARMOT_HOME/lib/json.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} create"

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
  else
    make_category_directories "$@"
    _config_add_categories "$(_fs_metadata_file)" "$@"
  fi
}

function make_category_directories() {
  local category_name
  category_name="$1"
  shift 1

  local category_path
  category_path="$(_fs_make_category_path "$category_name")"
  echo "+ $category_path (category)"

  local subcategory_path
  for subcategory_name in "$@"
  do
    subcategory_path="$(_fs_make_subcategory_path "$category_name" "$subcategory_name")"
    echo "+ $subcategory_path (sub-category)"
  done
}

function print_usage() {
  cat >&2 <<-EOF
USAGE
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION <category> [sub-category ...]

OPTIONS
--help        Show help

See \`man ${_MARMOT_INVOCATION// /-}\` for details.
EOF
}

## Main

main "$@"; exit
