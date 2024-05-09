#!/bin/zsh -i

emulate -LR zsh
#set -e

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} exec"

## Command

function main() {
  zparseopts -D -E \
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

  if [[ -n "$direnv_option" ]]
  then
    export DIRENV_LOG_FORMAT=''
  fi

  if [[ -n "$print_option" ]]
  then
    print_style="heading"
  fi

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
