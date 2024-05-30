
# Repository metadata

function _repomd_add_local_paths() {
  local config_file="$1" ; shift 1

  declare -a repositories=()
  local repo_path repository
  for some_path in "${@:#}"
  do
    repo_path="$(_repofs_normalize_path "$some_path")"
    repository="$(__repomd_repository_from_local_path "$repo_path")"
    repositories+=("$repository")
  done

  _jq_update "$config_file" \
    --argjson repositories "$(jo -a "${repositories[@]}")" \
    --sort-keys <<-'EOF'
    .
      | .meta_repo.repositories |= (. + $repositories | unique_by(.path))
      | .meta_repo.updated |= (now | todate)
EOF
}

function _repomd_delete_local_paths() {
  declare config_file="$1" ; shift 1
  declare -a remove_paths=("${@:#}")

  _jq_update "$config_file" \
    --argjson repository_paths "$(jo -a "${remove_paths[@]}")" \
    --sort-keys <<-'EOF'
    . | .meta_repo.categories[].repository_paths? -= $repository_paths
      | .meta_repo.repositories[]
        |= del(select(
                .path
                | in($repository_paths
                      | map(. as $elem | { key: $elem, value: 1 })
                      | from_entries)))
        | del(..|nulls)
      | .meta_repo.updated |= (now | todate)
EOF
}

function _repomd_local_paths() {
  local config_file="$1"

  # Treat lack of JSON fields as empty rather than as an error
  # https://github.com/jqlang/jq/issues/354#issuecomment-43147898
  jq -r \
    '.meta_repo.repositories[]?.path' \
    "$config_file"
}

function _repomd_local_paths_for_category() {
  local config_file="$1" category_or_subcategory="$2"
  local filter
  filter=$(cat <<-'EOF'
    .meta_repo.categories[]
      | select(.full_name == $full_name)
      | .repository_paths[]?
EOF
  )

  jq \
    --arg full_name "$category_or_subcategory" \
    -r \
    "$filter" "$config_file"
}

function _repomd_local_paths_reply() {
  local config_file="$1"
  local filter

  filter=$(cat <<-EOF
    [ .meta_repo.categories[].repository_paths[] ]
      + [ .meta_repo.repositories[].path ]
      | unique
      | .[]
EOF
)

  reply=()
  while read -r line
  do
    reply+=("$line")
  done < <(jq --raw-output "$filter" "$config_file")
}

## private

function __repomd_repository_from_local_path() {
  local repo_path="$1"
  jo -- "path=$repo_path"
}
