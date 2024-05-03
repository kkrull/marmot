#!/usr/bin/env zsh

# https://stackoverflow.com/a/56311706/112682
emulate -LR zsh

set -e

self="${0:P}"
self_basename="${0:t}"
self_dirname="${0:A:h}"
link_path='/usr/local/bin/marmot'

function main() {
  # Parse GNU-style long options
  # https://stackoverflow.com/questions/59981648/how-to-create-scripts-in-zsh-that-can-accept-out-of-order-arguments
  # https://zsh.sourceforge.io/Doc/Release/Zsh-Modules.html#The-zsh_002fzutil-Module
  zparseopts -D -E \
    -help=help_option

  if [[ $# == 0 || -n "$help_option" ]]
  then
    print_usage
    exit 0
  fi

  command="$1"
  case "$command" in
  'exec')
    shift 1
    exec "${self_dirname}/exec/marmot-exec.zsh" "$@"
    ;;

  'init')
    exec "${self_dirname}/init/marmot-init.zsh" "$@"
    ;;

  'link')
    ln -s "$self" "$link_path"
    echo "Added symlink: $link_path"
    ;;

  'unlink')
    rm -f "$link_path"
    echo "Removed symlink: $link_path"
    ;;

  *)
    echo "Unknown command: $command"
    exit 1
    ;;
  esac
}

function print_usage() {
  cat >&2 <<-EOF
${self_basename} - Meta Repo Management Tool

SYNOPSIS
${self_basename} command [options...]

OPTIONS
--help    Show help

COMMANDS
exec      Execute a command on a project's repositories
init      Make a new meta repo in the current directory

INSTALLATION
link      Add symlink so you can use this on your path
unlink    Remove symlink for this script
EOF
}

# Make sure the script exits, even if main doesn't
# https://unix.stackexchange.com/a/449508/37734
main "$@"; exit
