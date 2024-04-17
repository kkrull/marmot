#!/bin/zsh -i

emulate -LR zsh

# TODO KDK: Consider adding an option for whether to exit on the first failure, or keep going.
# Or ask the user if/when the first failure happens, since you probably don't know in advance.
#set -e

self_invocation="marmot exec"

function main() {
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

  # shellcheck disable=SC2154
  project_file="${project_file_option[2]}"

  # shellcheck disable=SC2086,SC2296
  project_repository_paths=("${(@f)"$(<${project_file})"}")

  if [[ $# == 0 ]]
  then
    echo "Missing: command"
    exit 1
  fi

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
${self_invocation} - Execute a command on multiple repositories

SYNOPSIS
${self_invocation}
  [--direnv] [--help] [--print]
  --project-file <file>
  <shell command> [args...]

OPTIONS
--direnv        Suppress distracting direnv output when changing directories.
--help          Show help.
--print         Print the name of each repository on its own line above any command output.
--project-file  The project to operate on.  See CONFIGURATION.

TIPS
git:
• Add --no-pager to git commands that would normally pipe to less (and pause for input).

EXAMPLES
• List version of Node.js used in repositories that use direnv+nvm:
    \$ ${self_invocation} --direnv --project node-projects.conf node --version

• Grep for matching source code in all repositories:
    \$ ${self_invocation} --project project.conf git --no-pager grep someFunction

CONFIGURATION
A project configuration file is a newline-delimited text file containing absolute paths to 1 or more
Git repositories.
EOF
}

main "$@"; exit
