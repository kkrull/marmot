#!/usr/bin/env zsh -i

set -e

repository_path="$HOME/ang/acb-lib"
printf "$repository_path: "
(cd "$repository_path" && $@)

repository_path="$HOME/ang/angServer"
printf "$repository_path: "
(cd "$repository_path" && $@)
