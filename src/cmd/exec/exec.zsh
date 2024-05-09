#!/bin/zsh -i

emulate -LR zsh
set -euo pipefail

source "$_MARMOT_HOME/lib/config.zsh"
source "$_MARMOT_HOME/lib/fs.zsh"

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} exec"

## Command

function main() {
  zparseopts -D -E \
    -category:=category_option \
    -direnv=direnv_option \
    -help=help_option \
    -print=print_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  elif [[ $# == 0 ]]
  then
    echo "$_MARMOT_INVOCATION: Missing command"
    exit 1
  fi

  local category_or_subcategory
  [[ -n "$category_option" ]] && category_or_subcategory="${category_option[2]}"

  [[ -n "$direnv_option" ]] && export DIRENV_LOG_FORMAT=''

  local print_next_repo_fn
  print_next_repo_fn=$(print_next_repo_fn_name "$print_option")

  local project_repository_paths
  _selected_repositories_reply "$(_fs_metadata_file)" "$category_or_subcategory"
  project_repository_paths=("${reply[@]}")

  for repository_path in "${project_repository_paths[@]}"
  do
    $print_next_repo_fn "$repository_path"
    (cd "$repository_path" && "$@")
  done
}

function _selected_repositories_reply() {
  local config_file category_or_subcategory
  config_file="$1"
  category_or_subcategory="$2"

  if [[ -n "$category_or_subcategory" ]]
  then
    # shellcheck disable=SC2296
    reply=("${(@f)"$(_config_repository_paths_in_category "$config_file" "$category_or_subcategory")"}")
  else
    # shellcheck disable=SC2296
    reply=("${(@f)"$(_config_repository_paths "$config_file")"}")
  fi
}

## Reporting

function print_next_repo_fn_name() {
  local print_option
  print_option="$1"
  if [[ -n "$print_option" ]]
  then
    echo print_next_repo_heading
  else
    echo print_next_repo_inline
  fi
}

# shellcheck disable=SC2317
function print_next_repo_heading() {
  printf "\n%s:\n" "$@"
}

# shellcheck disable=SC2317
function print_next_repo_inline() {
  printf "%s: " "$@"
}

function print_usage() {
  cat >&2 <<-EOF
$_MARMOT_INVOCATION - Execute a command in multiple repositories

SYNOPSIS
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION
  [--category <category|sub-category>]
  [--direnv] [--print]
  <shell command> [args...]

DESCRIPTION
This command repeats a given shell command on all repositories matching a
(sub-)category.

$_MARMOT_INVOCATION changes directories to each repository before running the
shell command, to ensure that any path-specific environment settings are
applied.  This is helpful for directory-based tools such as
\`direnv\`, \`fnm\`, and \`rvm\`, which update the shell's path and other
parts of its environment when changing directories.  The usefulness of the
shell command may depend upon it, for example when checking if all
repositories in a project use the same version of Node.js.

OPTIONS
--direnv        Suppress direnv output when changing directories
--help          Show help
--print         Print repository names above shell command output

TIPS
git:
• Add --no-pager to git commands that pipe to less (and pause for input)

EXAMPLES
• Git: Check which branches are checked out right now:
  \$ $_MARMOT_INVOCATION --category project/too-many-microservices \\
    git branch --show-current

• Git: Grep for matching source code in all repositories:
  \$ $_MARMOT_INVOCATION --category project/robot-masters --print \\
    git --no-pager grep dungeonType

• Git: Pull all the things!
  \$ $_MARMOT_INVOCATION --print \\
    git pull --ff-only origin

• Git: Push all the things!
  \$ $_MARMOT_INVOCATION --print \\
    git push

• Node: List version of Node.js used in repositories that use direnv+nvm:
  \$ $_MARMOT_INVOCATION --category platform/node --direnv \\
    node --version
EOF
}

## Main

main "$@"; exit
