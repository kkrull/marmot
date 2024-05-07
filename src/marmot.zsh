#!/usr/bin/env zsh

# https://stackoverflow.com/a/56311706/112682
emulate -LR zsh
set -e

## Local environment

link_path='/usr/local/bin/marmot'
self="${0:P}"

## Shared environment

export _MARMOT_HOME="${0:A:h}"
export _MARMOT_INVOCATION="${0:t}"

## Command

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
  # Installation
  'link')
    ln -s "$self" "$link_path"
    echo "Added symlink: $link_path"
    ;;

  'unlink')
    rm -f "$link_path"
    echo "Removed symlink: $link_path"
    ;;

  # Commands
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
$_MARMOT_INVOCATION - Meta Repo Management Tool

SYNOPSIS
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION command [options...]

DESCRIPTION
Marmot creates and maintains a Meta Repository (e.g. "meta repo"), which can
be used to group several Git repositories by 1 or more arbitrary categories.

Marmot creates a directory structure in the meta repo's file system to mirror
the way that repositories have been categorized, so that there is a
\`/:category/:value\` directory for each known value of each category.  Each
directory contains symbolic links back to the Git repositories that share
the same value for the same category.

Users run commands from one of these directories in order to restrict
commands to the Git repositories that have that categorization in common.
In this fashion, users can do things like search closely-related
code with \`git grep\` or open an editor for those Git repositories, without
clutter and noise from irrelevant sources in unrelated repositories.

OPTIONS
--help        Show help

COMMANDS
category      Work with categories
exec          Execute a command on a project's repositories
init          Make a new meta repo in the current directory
repo          Work with repositories

INSTALLATION
link          Add symlink so you can use this on your path
unlink        Remove symlink for this script
EOF
}

## Main

# Make sure the script exits, even if main doesn't
# https://unix.stackexchange.com/a/449508/37734
main "$@"; exit
