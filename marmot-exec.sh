#!/usr/bin/env zsh

set -e


printf "$HOME/ang/acb-lib: "
(cd "$HOME/ang/acb-lib" && node --version)

printf "$HOME/ang/angServer: "
(cd "$HOME/ang/angServer" && node --version)
