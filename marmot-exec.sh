#!/usr/bin/env zsh -i

set -e

project_repository_paths=("$HOME/ang/acb-lib" "$HOME/ang/angServer")

for repository_path in $project_repository_paths; do
  printf "$repository_path: "
  (cd "$repository_path" && $@)
done
