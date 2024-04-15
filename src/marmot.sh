#!/usr/bin/env zsh

set -e

self_name="${0:A}"
self_dir="$(dirname "$self_name")"

case "$1" in
'exec')
  shift 1
  exec "${self_dir}/exec/marmot-exec.sh" "$@"
  ;;

'link')
  set -x
  ln -s "$self_name" /usr/local/bin/marmot
  ;;

'unlink')
  set -x
  rm -f /usr/local/bin/marmoot
  ;;

*)
  cat >&2 <<-EOF
Meta Repo Management Tool
Usage: $0 command [options...]

exec      Execute a command on a project's repositories
link      Add symlink for this script to /usr/local/bin
unlink    Remove symlink for this script
EOF

  exit 1
  ;;
esac
