#!/bin/zsh

emulate -LR zsh
set -euo pipefail
while IFS= read -d $'\0' -r f; do
  source "$f"
done < <(find "$_MARMOT_HOME/lib" -type f -iname '*.zsh' -print0 | sort -nz)

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} list"

## Command

function main() {
  zparseopts -D -E \
    -category:=category_option \
    -help=help_option

  if [[ $# -gt 0 ]]
  then
    echo "Unknown option: $1"
    exit 1
  fi

  if [[ -n "$help_option" ]]
  then
    print_usage
  elif [[ -n "$category_option" ]]
  then
    _repocmd_list_category "${category_option[2]}"
  else
    _repocmd_list_all
  fi
}

function print_usage() {
  cat >&2 <<-EOF
USAGE
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION [--category <category|category/sub-category>]

OPTIONS
--category    List repositories that have been added to the given category
              or sub-category.
--help        Show help

See \`man ${_MARMOT_INVOCATION// /-}\` for details.
EOF
}

## Main

main "$@"
