#!/bin/zsh -i

#set -e

if [[ "$1" == "--direnv" ]]
then
  export DIRENV_LOG_FORMAT=''
  shift 1
fi

if [[ $1 == "--print" ]]
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
  echo "Usage: $0 [--direnv] --project-file <file> <command> [args...]"
  echo "Example: $0 --project-file my-project.conf node --version"
  exit 1
fi

for repository_path in "${project_repository_paths[@]}"; do
  if [[ $print_style == "heading" ]]
  then
    printf "\n%s:\n" "$repository_path"
  else
    printf "%s: " "$repository_path"
  fi

  (cd "$repository_path" && "$@")
done
