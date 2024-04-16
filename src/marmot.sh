#!/usr/bin/env zsh

set -e

command_path="${0:A:h}"
link_path='/usr/local/bin/marmot'
self_command="${0:t}"

case "$1" in
'exec')
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
  cat >&2 <<-EOF
Meta Repo Management Tool
Usage: ${self_command} command [options...]

COMMANDS
exec      Execute a command on a project's repositories

INSTALLATION
link      Add symlink so you can use this on your path
unlink    Remove symlink for this script
EOF

  exit 1
  ;;
esac
