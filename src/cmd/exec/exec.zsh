#!/bin/zsh -i

emulate -LR zsh
#set -e

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} exec"

## Command

# TODO KDK: Consider adding an option for whether to exit on the first failure, or keep going.
# Or ask the user if/when the first failure happens, since you probably don't know in advance.

function main() {
  # Parse GNU-style long options
  # https://stackoverflow.com/questions/59981648/how-to-create-scripts-in-zsh-that-can-accept-out-of-order-arguments
  # https://zsh.sourceforge.io/Doc/Release/Zsh-Modules.html#The-zsh_002fzutil-Module

  # TODO KDK: Change it to `--project-file=`, to avoid ambiguity with the command to run
  zparseopts -D -E \
    -direnv=direnv_option \
    -help=help_option \
    -print=print_option \
    -project-file:=project_file_option

  if [[ -n "$direnv_option" ]]
  then
    export DIRENV_LOG_FORMAT=''
  fi

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  fi

  if [[ -n "$print_option" ]]
  then
    print_style="heading"
  fi

  if [[ $# == 0 ]]
  then
    echo "$_MARMOT_INVOCATION: Missing command"
    exit 1
  elif [[ -z "$project_file_option" ]]
  then
    echo "$_MARMOT_INVOCATION: Missing --project-file"
    exit 1
  fi

  # shellcheck disable=SC2154
  project_file="${project_file_option[2]}"

  # shellcheck disable=SC2086,SC2296
  project_repository_paths=("${(@f)"$(<${project_file})"}")

  for repository_path in "${project_repository_paths[@]}"
  do
    if [[ "$print_style" == "heading" ]]
    then
      printf "\n%s:\n" "$repository_path"
    else
      printf "%s: " "$repository_path"
    fi

    (cd "$repository_path" && "$@")
  done
}

function print_usage() {
  cat >&2 <<-EOF
$_MARMOT_INVOCATION - Execute a command repeatedly

SYNOPSIS
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION
  [--direnv] [--print]
  --project-file <file>
  <shell command> [args...]

DESCRIPTION
This repeats a given shell command on all repositories that are part of a project.

OPTIONS
--direnv        Suppress distracting direnv output when changing directories
--help          Show help
--print         Print the name of each repository on its own line above any command output
--project-file  The project to operate on (see CONFIGURATION)

TIPS
git:
• Add --no-pager to git commands that would normally pipe to less (and pause for input)

EXAMPLES
• List version of Node.js used in repositories that use direnv+nvm:
    \$ $_MARMOT_INVOCATION --direnv --project node-projects.conf node --version

• Grep for matching source code in all repositories:
    \$ $_MARMOT_INVOCATION --project project.conf git --no-pager grep someFunction

CONFIGURATION
A project configuration file is a newline-delimited text file containing absolute paths to 1 or more
Git repositories.
EOF
}

## Main

main "$@"; exit
