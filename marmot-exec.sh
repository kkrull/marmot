#!/usr/bin/env zsh -i

set -e

if [[ "$1" == "--direnv" ]]
then
  export DIRENV_LOG_FORMAT=''
  shift 1
fi

if (( $# < 1 ))
then
  echo "Usage: $0 [--direnv] <command> [args...]"
  echo "Example: $0 --direnv node --version"
  exit 1
fi

project_repository_paths=("$HOME/ang/acb-lib" "$HOME/ang/angServer")

for repository_path in $project_repository_paths; do
  printf "$repository_path: "
  (cd "$repository_path" && $@)
done
