#!/bin/zsh

emulate -LR zsh

self_invocation="marmot category add"
working_dirname="${PWD:A}"

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  elif [[ $# == 0 ]]
  then
    echo "$self_invocation: Missing category name"
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
  category_value_path="$working_dirname/$category_value_name"

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
${self_invocation} - Add repositories to a category

SYNOPSIS
${self_invocation} <category> <repository> ...repository
${self_invocation} <category>/<value> <repository> ...repository
  [--help]

DESCRIPTION
This command adds 1 or more repositories to a category, or to a particular
value of a category.

OPTIONS
--help        Show help

EXAMPLES
• Add a repository to the "user" category:
    \$ ${self_invocation} user ~/git/dotfiles
• Add some repositories to the "skunkworks" project (something big is about to happen):
    \$ ${self_invocation} project/skunkworks ~/git/robot-masters ~/git/evil-castle
EOF
}

main "$@"; exit
