#!/usr/bin/env zsh

# https://stackoverflow.com/a/56311706/112682
emulate -LR zsh

set -e

command_path="${0:A:h}"
link_path='/usr/local/bin/marmot'
self_command="${0:t}"

function usage() {
  cat >&2 <<-EOF
${self_command} - Meta Repo Management Tool

SYNOPSIS
${self_command} command [options...]

OPTIONS
--help    Show help.

COMMANDS
exec      Execute a command on a project's repositories

INSTALLATION
link      Add symlink so you can use this on your path
unlink    Remove symlink for this script
EOF
}

# https://stackoverflow.com/questions/59981648/how-to-create-scripts-in-zsh-that-can-accept-out-of-order-arguments
# https://zsh.sourceforge.io/Doc/Release/Zsh-Modules.html#The-zsh_002fzutil-Module
zparseopts -D -E \
  -help=help_option

command="$1"

if [[ -n "$help_option" || -z "$*" ]]
then
  usage
  exit 0
fi

case "$command" in
'exec')
  set -x
  shift 1
  exec "${command_path}/exec/marmot-exec.sh" "$@"
  ;;

'link')
  set -x
  ln -s "${0:P}" "$link_path"
  ;;

'unlink')
  set -x
  rm -f "$link_path"
  ;;

*)
  echo "Unknown command: $command"
  exit 1
  ;;
esac
