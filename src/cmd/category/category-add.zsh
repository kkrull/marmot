#!/bin/zsh

emulate -LR zsh
set -e

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"
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

  local category_or_subcategory
  category_or_subcategory="$1"

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
$_MARMOT_INVOCATION - Add repositories to a category

SYNOPSIS
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION <category> <repository> [repository...]
$_MARMOT_INVOCATION <category>/<sub-category> <repository> [repository...]

DESCRIPTION
This command adds 1 or more repositories to a (sub-)category.

OPTIONS
--help        Show help

EXAMPLES
• Add a repository to the "user" category:
  \$ $_MARMOT_INVOCATION user ~/git/dotfiles
• Add some repositories to the "skunkworks" project (lookout Dr. Light):
  \$ $_MARMOT_INVOCATION project/skunkworks ~/git/robot-masters ~/git/skull-fortress
EOF
}

## Main

main "$@"; exit
