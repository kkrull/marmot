#!/usr/bin/env zsh

# https://stackoverflow.com/a/56311706/112682
emulate -LR zsh

set -e

base_path="${0:A:h}"
link_path='/usr/local/bin/marmot'
self_name="${0:t}"

function main() {
  # https://stackoverflow.com/questions/59981648/how-to-create-scripts-in-zsh-that-can-accept-out-of-order-arguments
  # https://zsh.sourceforge.io/Doc/Release/Zsh-Modules.html#The-zsh_002fzutil-Module
  zparseopts -D -E \
    -help=help_option

  command="$1"

  if [[ -n "$help_option" || $# == 0 ]]
  then
    print_usage
    exit 0
  fi

  case "$command" in
  'exec')
    shift 1
    exec "${base_path}/exec/marmot-exec.sh" "$@"
    ;;

  'link')
    ln -s "${0:P}" "$link_path" && echo "Added symlink: $link_path"
    ;;

  'unlink')
    rm -f "$link_path" && echo "Removed symlink: $link_path"
    ;;

  *)
    echo "Unknown command: $command"
    exit 1
    ;;
  esac
}

function print_usage() {
  cat >&2 <<-EOF
${self_name} - Meta Repo Management Tool

SYNOPSIS
${self_name} command [options...]

OPTIONS
--help    Show help

COMMANDS
exec      Execute a command on a project's repositories

INSTALLATION
link      Add symlink so you can use this on your path
unlink    Remove symlink for this script
EOF
}

# https://unix.stackexchange.com/a/449508/37734
main "$@"; exit
