#!/bin/zsh

emulate -LR zsh

source "$_MARMOT_HOME/lib/paths.zsh"

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
  else
    link_to_category "$@"
  fi
}

function link_to_category() {
  local category_value_name
  category_value_name="$1"
  shift 1

  local category_value_path
  category_value_path="$(meta_repo_home)/$category_value_name"

  local link_path
  local repository_name
  local repository_path
  for repository_path in "$@"
  do
    repository_name="${repository_path:t}"
    link_path="$category_value_name/$repository_name"
    echo "+ ${link_path} -> $repository_path"
    (cd "$category_value_path" && ln -s "$repository_path" "$repository_name")
  done
}

function print_usage() {
  cat >&2 <<-EOF
$_MARMOT_INVOCATION - Add repositories to a category

SYNOPSIS
$_MARMOT_INVOCATION [--help]
$_MARMOT_INVOCATION <category> <repository> ...repository
$_MARMOT_INVOCATION <category>/<sub-category> <repository> ...repository

DESCRIPTION
This command adds 1 or more repositories to a category, or to a particular
value of a category.

OPTIONS
--help        Show help

EXAMPLES
• Add a repository to the "user" category:
    \$ $_MARMOT_INVOCATION user ~/git/dotfiles
• Add some repositories to the "skunkworks" project (something big is about to happen):
    \$ $_MARMOT_INVOCATION project/skunkworks ~/git/robot-masters ~/git/evil-castle
EOF
}

## Main

main "$@"; exit
