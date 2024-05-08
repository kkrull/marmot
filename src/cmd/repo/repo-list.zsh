#!/bin/zsh

emulate -LR zsh
set -e

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"
source "$_MARMOT_HOME/lib/json.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} list"

## Command

function main() {
  zparseopts -D -E \
    -category:=category_option \
    -help=help_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  elif [[ $# -gt 0 ]]
  then
    echo "Unknown option: $1"
    exit 1
  elif [[ -n "$category_option" ]]
  then
    local category_or_subcategory
    category_or_subcategory="${category_option[2]}"
    echo "[list_local_repositories] category=$category_or_subcategory"
    exit 0
  else
    list_local_repositories "$(_fs_metadata_file)"
    exit 0
  fi
}

function list_local_repositories() {
  local config_file
  config_file="$1"
  shift 1

  _config_repository_paths "$config_file"
}

function print_usage() {
  cat >&2 <<-EOF
$_MARMOT_INVOCATION - List repositories

SYNOPSIS
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION [--category <category|sub-category>]

DESCRIPTION
This command lists repositories that have been registered with Marmot.
Given options, this lists only the repositories that match the given criteria.

OPTIONS
--category    List repositories that have been added to the given category
              or sub-category.
--help        Show help

EXAMPLES
• List registered TypeScript repositories
  \$ $_MARMOT_INVOCATION --category lang/typescript
EOF
}

## Main

main "$@"
