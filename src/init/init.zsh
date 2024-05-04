#!/bin/zsh

emulate -LR zsh

self_invocation="marmot init"

working_dirname="${PWD:A}"
meta_repo_data="$working_dirname/.marmot"

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  elif [[ -d "$meta_repo_data" ]]
  then
    printf "Meta repo already exists: %s" "$meta_repo_data"
    exit 1
  else
    create_meta_repo "$meta_repo_data"
  fi
}

function create_meta_repo() {
  local directory="$1"
  mkdir -p "$directory"

  local template
  template=$(cat <<'EOF'
{
  meta_repo: {
    categories: []
  },
  version: $version
}
EOF
  )

  local meta_repo_file="$directory/meta-repo.json"
  jq > "$meta_repo_file" \
    --null-input \
    --arg version 0.0.1 \
    "$template"

  echo "Initialized meta repository at $working_dirname"
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - Initialize a meta repository in the current directory

SYNOPSIS
${self_invocation} [--help]

OPTIONS
--help        Show help
EOF
}

main "$@"; exit