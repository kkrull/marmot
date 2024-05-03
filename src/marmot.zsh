#!/usr/bin/env zsh

# https://stackoverflow.com/a/56311706/112682
emulate -LR zsh

set -e

self="${0:P}"
self_basename="${0:t}"
self_dirname="${0:A:h}"
link_path='/usr/local/bin/marmot'

function main() {
  if [[ $# == 0 ]]
  then
    print_usage
    exit 0
  fi

  # Distinguish `marmot --help` (top-level usage) from `marmot <command> --help` (command usage)
  zparseopts -E \
    -help=help_option
  if [[ $# == 1 && -n "$help_option" ]]
  then
    print_usage
    exit 0
  fi

  command="$1"
  case "$command" in
  'collection')
    shift 1
    exec "${self_dirname}/collection/marmot-collection.zsh" "$@"
    ;;

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
collection  Work with collections
exec        Execute a command on a project's repositories
init        Make a new meta repo in the current directory

INSTALLATION
link      Add symlink so you can use this on your path
unlink    Remove symlink for this script
EOF
}

# Make sure the script exits, even if main doesn't
# https://unix.stackexchange.com/a/449508/37734
main "$@"; exit
