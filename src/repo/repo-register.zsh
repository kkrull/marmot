#!/bin/zsh

emulate -LR zsh

## Command

self_invocation="marmot repo register"

function main() {
  zparseopts -D -E \
    -help=help_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  elif [[ $# == 0 ]]
  then
    echo "$self_invocation: Missing repository path"
    exit 1
  else
    register_local_repositories "$(meta_repo_config_file)" "$@"
    exit 0
  fi
}

function print_usage() {
  cat >&2 <<-EOF
${self_invocation} - Register repositories to manage

SYNOPSIS
${self_invocation} <Git repository> ...
${self_invocation} [--help]

DESCRIPTION
This command registers 1 or more repositories with Marmot, so it can manage them.

OPTIONS
--help        Show help
EOF
}

function register_local_repositories() {
  local config_file
  config_file="$1"
  shift 1

  local new_repositories
  new_repositories=$(to_marmot_repositories "$@")

  echo "[register_local_repositories] config_file=$config_file new_repositories=$new_repositories"
  jq < "$config_file" \
    ".meta_repo_next.repositories += ${new_repositories}"
}

## JSON

function to_json_array() {
  if [[ $# == 0 ]]
  then
    echo "[]"
  elif [[ $# == 1 ]]
  then
    echo "[$1]"
  else
    printf "[%s" "$1"
    for element in "${@:2}"
    do
      printf ", %s" "$element"
    done

    printf ']'
  fi
}

## Marmot configuration

function to_marmot_repositories() {
  local new_repositories new_repository

  new_repositories=()
  for repository_path in "$@"
  do
    new_repository=$(to_marmot_repository "$repository_path")
    new_repositories+=("$new_repository")
  done

  to_json_array "${new_repositories[@]}"
}

function to_marmot_repository() {
  local repo_path
  repo_path="$1"
  echo "{ \"path\": \"$repo_path\" }"
}

## Marmot paths

function meta_repo_config_file() {
  local meta_home

  meta_home="${PWD:A}"
  echo "$meta_home/.marmot/meta-repo.json"
}

## Main

main "$@"
