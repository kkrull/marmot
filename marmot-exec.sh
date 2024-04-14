#!/usr/bin/env zsh -i

set -e

if [[ "$1" == "--direnv" ]]
then
  export DIRENV_LOG_FORMAT=''
  shift 1
fi

if [[ "$1" == "--project-file" ]]
then
  set -x
  profile_file="$2"
  shift 2

  typeset -a project_repository_paths=("${(@f)"$(<conf/acb-project.conf)"}")
  set +x
else
  echo "Missing: --project-file <file>"
  exit 1
fi

if (( $# < 1 ))
then
  echo "Usage: $0 [--direnv] --project-file <file> <command> [args...]"
  echo "Example: $0 --project-file my-project.conf node --version"
  exit 1
fi

# project_repository_paths=("$HOME/ang/acb-lib" "$HOME/ang/angServer")

for repository_path in $project_repository_paths; do
  printf "$repository_path: "
  (cd "$repository_path" && $@)
done
