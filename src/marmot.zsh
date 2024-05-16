#!/usr/bin/env zsh

# https://stackoverflow.com/a/56311706/112682
emulate -LR zsh
set -e

## Shared environment

export _MARMOT_HOME="${0:A:h}"
export _MARMOT_INVOCATION="${0:t}"

## Local environment

source "$_MARMOT_HOME/lib/fs.zsh"

## Command

function main() {
  if [[ $# == 0 ]]
  then
    print_usage
    exit 0
  fi

  # Distinguish `marmot --help` (top-level usage) from `marmot <command> --help` (command usage)
  zparseopts -E \
    -help=help_option \
    -version=version_option
  if [[ $# == 1 && -n "$help_option" ]]
  then
    print_usage
    exit 0
  elif [[ $# == 1 && -n "$version_option" ]]
  then
    print_version
    exit 0
  fi

  command="$1"
  case "$command" in
  'category')
    shift 1
    exec "$_MARMOT_HOME/cmd/category/category.zsh" "$@"
    ;;

  'exec')
    shift 1
    exec "$_MARMOT_HOME/cmd/exec/exec.zsh" "$@"
    ;;

  'init')
    shift 1
    exec "$_MARMOT_HOME/cmd/init/init.zsh" "$@"
    ;;

  'meta')
    shift 1
    exec "$_MARMOT_HOME/cmd/meta/meta.zsh" "$@"
    ;;

  'repo')
    shift 1
    exec "$_MARMOT_HOME/cmd/repo/repo.zsh" "$@"
    ;;

  *)
    echo "Unknown command: $command"
    exit 1
    ;;
  esac
}

function print_usage() {
  cat >&2 <<-EOF
USAGE
$_MARMOT_INVOCATION [--help] [--version]
$_MARMOT_INVOCATION command [args ...]

OPTIONS
--help        Show help
--version     Show version

COMMANDS
category      Work with categories
exec          Execute a command in multiple repositories
init          Make a new meta repo in the default directory
meta          Information about the meta repo (not the data it manages)
repo          Work with repositories

See \`man ${_MARMOT_INVOCATION// /-}\` for details.
EOF
}

function print_version() {
  _fs_marmot_version
}

## Main

# Make sure the script exits, even if main doesn't
# https://unix.stackexchange.com/a/449508/37734
main "$@"; exit
