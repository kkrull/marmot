
# Repository metadata

function _repomd_add_local_path() {
  local data_file="$1"
  local some_repo_path="$2"
  local ssh_url="$3"

  declare -a repositories=()
  local repo_path repository
  repo_path="$(_repofs_normalize_path "$some_repo_path")"
  repository="$(__repomd_repository_from_local_path_and_url "$repo_path" "$ssh_url")"
  repositories+=("$repository")

  _jq_update "$data_file" \
    --argjson repositories "$(jo -a "${repositories[@]}")" \
    --sort-keys <<-'EOF'
    .
      | .meta_repo.local_repositories |= (. + $repositories | unique_by(.path))
      | .meta_repo.updated |= (now | todate)
EOF
}

function _repomd_add_local_paths() {
  local data_file="$1" ; shift 1

  declare -a repositories=()
  local repo_path repository
  for some_path in "${@:#}"
  do
    repo_path="$(_repofs_normalize_path "$some_path")"
    repository="$(__repomd_repository_from_local_path "$repo_path")"
    repositories+=("$repository")
  done

  #TODO KDK: Rename to local_repositories
  _jq_update "$data_file" \
    --argjson repositories "$(jo -a "${repositories[@]}")" \
    --sort-keys <<-'EOF'
    .
      | .meta_repo.repositories |= (. + $repositories | unique_by(.path))
      | .meta_repo.updated |= (now | todate)
EOF
}

function _repomd_add_remote() {
  local data_file="$1"
  local ssh_url="$2"

  declare -a repositories=()
  local repository
  repository="$(__repomd_repository_from_url "$ssh_url")"
  repositories+=("$repository")

  #TODO KDK: Rename to remote_repositories
  _jq_update "$data_file" \
    --argjson repositories "$(jo -a "${repositories[@]}")" \
    --sort-keys <<-'EOF'
    .
      | .meta_repo.repositories |= (. + $repositories | unique_by(.path))
      | .meta_repo.updated |= (now | todate)
EOF
}

function _repomd_delete_local_paths() {
  declare data_file="$1" ; shift 1
  declare -a remove_paths=("${@:#}")

  _jq_update "$data_file" \
    --argjson repository_paths "$(jo -a "${remove_paths[@]}")" \
    --sort-keys <<-'EOF'
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
  local data_file="$1"

  # Treat lack of JSON fields as empty rather than as an error
  # https://github.com/jqlang/jq/issues/354#issuecomment-43147898
  jq -r \
    '.meta_repo.repositories[]?.path' \
    "$data_file"
}

function _repomd_local_paths_for_category() {
  local data_file="$1" category_or_subcategory="$2"
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
    "$filter" "$data_file"
}

function _repomd_local_paths_reply() {
  local data_file="$1"
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
  done < <(jq --raw-output "$filter" "$data_file")
}

## private

function __repomd_repository_from_local_path() {
  local repo_path="$1"
  jo -- "path=$repo_path"
}

function __repomd_repository_from_local_path_and_url() {
  local repo_path="$1"
  local ssh_url="$2"
  jo -- "path=$repo_path" "ssh_url=$ssh_url"
}

function __repomd_repository_from_url() {
  local ssh_url="$1"
  jo -- "ssh_url=$ssh_url"
}
