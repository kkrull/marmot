
# Category commands

function _categorycmd_rm() {
  local category_or_subcategory="$1" ; shift 1

  _categorymd_remove_repositories_as_local_paths \
    "$(_fs_metadata_file)" \
    "$category_or_subcategory" \
    "${@:#}"
  rm_link_to_category "$category_or_subcategory" "${@:#}"
}

function rm_link_to_category() {
  local category_name="$1" ; shift 1

  local link_path repository_path
  for repository_path in "${@:#}"
  do
    link_path="$(_categoryfs_rm_repository_link "$category_name" "$repository_path")"
    echo "- ${link_path} (link)"
  done
}
