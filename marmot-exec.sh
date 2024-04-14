#!/usr/bin/env zsh

set -e


printf "$HOME/ang/acb-lib: "
pushd . > /dev/null
cd "$HOME/ang/acb-lib"
node --version
popd > /dev/null

printf "$HOME/ang/angServer: "
(cd "$HOME/ang/angServer" && node --version)
