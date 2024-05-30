
# Repository commands

function _repocmd_list_all() {
  _repomd_local_paths "$(_fs_metadata_file)"
}

function _repocmd_list_category() {
  local category_id="$1"
  _repomd_local_paths_for_category "$(_fs_metadata_file)" "$category_id"
}

function _repocmd_prune_missing() {
  local config_file
  config_file="$(_fs_metadata_file)"

  declare reply stale_paths=()
  _repomd_local_paths_reply "$config_file"
  for repo_path in "${reply[@]}"
  do
    [[ ! -d "$repo_path" ]] && stale_paths+=("$repo_path")
  done

  [[ "${#stale_paths}" -eq 0 ]] && exit 0

  _repomd_delete_local_paths "$config_file" "${stale_paths[@]}"
  for removed_path in "${stale_paths[@]}"
  do
    echo "- $removed_path (repository)"
  done
}

function _repocmd_register_local_paths() {
  _repomd_add_local_paths "$(_fs_metadata_file)" "$@"
}
