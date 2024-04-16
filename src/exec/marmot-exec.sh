#!/bin/zsh -i

# TODO KDK: Consider adding an option for whether to exit on the first failure, or keep going.
# Or ask the user if/when the first failure happens, since you probably don't know in advance.
#set -e

self_command="marmot exec"

function usage() {
  cat >&2 <<-EOF
${self_command} - Execute a command on multiple repositories

SYNOPSIS
${self_command}
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
    \$ ${self_command} --direnv --project node-projects.conf node --version

• Grep for matching source code in all repositories:
    \$ ${self_command} --project project.conf git --no-pager grep someFunction

CONFIGURATION
A project configuration file is a newline-delimited text file containing absolute paths to 1 or more
Git repositories.
EOF
}

if [[ "$1" == "--direnv" ]]
then
  export DIRENV_LOG_FORMAT=''
  shift 1
fi

if [[ "$1" == "--help" ]]
then
  usage
  exit 0
fi

if [[ "$1" == "--print" ]]
then
  print_style="heading"
  shift 1
fi

if [[ "$1" == "--project-file" ]]
then
  project_file="$2"
  shift 2

  # shellcheck disable=SC2086,SC2296
  project_repository_paths=("${(@f)"$(<${project_file})"}")
else
  echo "Missing: --project-file <file>"
  exit 1
fi

if (( $# < 1 ))
then
  usage
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
