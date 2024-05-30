#!/bin/zsh -i

emulate -LR zsh
set -euo pipefail
while IFS= read -d $'\0' -r f; do
  source "$f"
done < <(find -s "$_MARMOT_HOME/lib" -type f -iname '*.zsh' -print0)

## Shared environment

export _MARMOT_INVOCATION="${_MARMOT_INVOCATION} exec"

## Command

function main() {
  zparseopts -D -E \
    -category:=category_option \
    -direnv=direnv_option \
    -help=help_option \
    -repo-names:=repo_names_option

  if [[ -n "$help_option" ]]
  then
    print_usage
    exit 0
  elif [[ $# == 0 ]]
  then
    echo "$_MARMOT_INVOCATION: Missing command"
    exit 1
  fi

  local category_or_subcategory
  [[ -n "$category_option" ]] && category_or_subcategory="${category_option[2]}"

  [[ -n "$direnv_option" ]] && export DIRENV_LOG_FORMAT=''

  local print_next_repo_fn
  # shellcheck disable=SC2154
  print_next_repo_fn=$(print_next_repo_fn_name "${repo_names_option[@]}") \
    || { echo "$print_next_repo_fn" ; exit 1; }

  local project_repository_paths
  _selected_repositories_reply "$(_fs_metadata_file)" "$category_or_subcategory"
  project_repository_paths=("${reply[@]}")

  for repository_path in "${project_repository_paths[@]}"
  do
    $print_next_repo_fn "$repository_path"
    (cd "$repository_path" && "$@")
  done
}

function _selected_repositories_reply() {
  local config_file category_or_subcategory
  config_file="$1"
  category_or_subcategory="$2"

  if [[ -n "$category_or_subcategory" ]]
  then
    # shellcheck disable=SC2296
    reply=("${(@f)"$(_config_repository_paths_in_category "$config_file" "$category_or_subcategory")"}")
  else
    # shellcheck disable=SC2296
    reply=("${(@f)"$(_config_repository_paths "$config_file")"}")
  fi
}

## Reporting

function print_next_repo_fn_name() {
  local option_name option_value
  option_name="${1-}"
  option_value="${2-inline}"

  case "$option_value" in
  'heading')
    echo print_next_repo_heading
    ;;

  'inline')
    echo print_next_repo_inline
    ;;

  *)
    # Produce an error message instead of the name of the function
    echo "$option_name: Invalid value \"$option_value\""
    exit 1
    ;;
  esac
}

__print_next_repo_heading_first=1

# shellcheck disable=SC2317
function print_next_repo_heading() {
  case "$__print_next_repo_heading_first" in
  '1')
    # Don't start a new paragraph, the first time
    __print_next_repo_heading_first=0
    printf "%s:\n" "$@"
    ;;

  '0')
    printf "\n%s:\n" "$@"
    ;;
  esac
}

# shellcheck disable=SC2317
function print_next_repo_inline() {
  printf "%s: " "$@"
}

function print_usage() {
  cat >&2 <<-EOF
USAGE
$_MARMOT_INVOCATION --help
$_MARMOT_INVOCATION
  [--category <category|category/sub-category>]
  [--direnv] [--repo-names <inline|heading>]
  <shell command> [args ...]

OPTIONS
--direnv        Suppress direnv output when changing directories
--help          Show help
--repo-names    Print repository names \`inline\` prior to or as a \`heading\`
                above shell command output

See \`man ${_MARMOT_INVOCATION// /-}\` for details.
EOF
}

## Main

main "$@"; exit
