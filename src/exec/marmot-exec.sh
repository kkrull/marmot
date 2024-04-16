#!/bin/zsh -i

# TODO KDK: Consider adding an option for whether to exit on the first failure, or keep going.
# Or ask the user if/when the first failure happens, since you probably don't know in advance.
#set -e

command_path="${0:A:r}"
self_command="marmot exec"
self_dir="${0:A:h}"

function usage() {
  echo "Usage: $0 [--direnv] --project-file <file> <shell command> [args...]"
  echo "Example: $0 --project-file my-project.conf node --version"

  cat >&2 <<-EOF
Execute a command on multiple repositories.
Usage: ${self_command}
  [options...] --project-file <file>
  <shell command> [args...]

OPTIONS
--direnv        Open a console with the mysql client.
--help          Show help.
--print         Print the name of each repository on its own line above any
                command output.
--project-file  The project to operate on.  See PROJECT CONFIGURATION.

PROJECT CONFIGURATION
A project configuration file is a newline-delimited text file of absolute
paths to 1 or more Git repositories.

EXAMPLES
List version of Node.js used in repositories that use direnv+nvm:
\$ ${self_command} --direnv --project node-projects.conf node --version

Grep for matching source code in all repositories:
\$ ${self_command} --project project.conf git --no-pager grep someFunction

TIPS
direnv    Add --direnv option to marmot exec commands to suppress distracting
          direnv output.
git       Add --no-pager to git commands that would normally pipe to less
          (and pause for input).
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
