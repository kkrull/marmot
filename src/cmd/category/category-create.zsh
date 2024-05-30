#!/bin/zsh

emulate -LR zsh
set -euo pipefail
while IFS= read -d $'\0' -r f; do
  source "$f"
done < <(find -s "$_MARMOT_HOME/lib" -type f -iname '*.zsh' -print0)

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
  elif [[ $# == 0 || -z "$1" ]]
  then
    echo "$_MARMOT_INVOCATION: Missing category name"
    exit 1
  else
    # Remove array elements matching '' to avoid entering a world of pain
    # https://zsh.sourceforge.io/Doc/Release/Expansion.html#Parameter-Expansion
    make_category_directories "${@:#}"
    _config_add_categories "$(_fs_metadata_file)" "${@:#}"
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
