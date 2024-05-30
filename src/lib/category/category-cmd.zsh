
# Category commands

# ${@:#} Remove array elements matching '' to avoid entering a world of pain
# https://zsh.sourceforge.io/Doc/Release/Expansion.html#Parameter-Expansion

function _categorycmd_add_repository_paths() {
  local category_or_subcategory="$1" ; shift 1

  # Given (sub-)category may be new; create if so
  ensure_create_category "$category_or_subcategory"

  _categorymd_add_repositories_as_local_path \
    "$(_fs_metadata_file)" \
    "$category_or_subcategory" \
    "${@:#}"
  link_to_category "$category_or_subcategory" "${@:#}"
}

function _categorycmd_create() {
  make_category_directories "${@:#}"
  _categorymd_create "$(_fs_metadata_file)" "${@:#}"
}

function _categorycmd_list() {
  _categorymd_full_names "$(_fs_metadata_file)"
}

function _categorycmd_rm_repository_paths() {
  local category_or_subcategory="$1" ; shift 1

  _categorymd_remove_repositories_as_local_paths \
    "$(_fs_metadata_file)" \
    "$category_or_subcategory" \
    "${@:#}"
  rm_link_to_category "$category_or_subcategory" "${@:#}"
}

## private

function ensure_create_category() {
  local category_or_subcategory category_name subcategory_name
  category_or_subcategory="$1"
  category_name="$(_categoryid_category "$category_or_subcategory")"
  subcategory_name="$(_categoryid_subcategory "$category_or_subcategory")"

  _categorymd_create "$(_fs_metadata_file)" "$category_name" "$subcategory_name"
  _categoryfs_mkdir "$category_name" > /dev/null
  _categoryfs_mkdir_subcategory "$category_name" "$subcategory_name" > /dev/null
}

function link_to_category() {
  local category_name="$1" ; shift 1

  local link_path repository_path
  for repository_path in "${@:#}"
  do
    link_path="$(_categoryfs_add_repository_link "$category_name" "$repository_path")"
    echo "+ ${link_path} (link)"
  done
}

function make_category_directories() {
  local category_name
  category_name="$1"
  shift 1

  local category_path
  category_path="$(_categoryfs_mkdir "$category_name")"
  echo "+ $category_path (category)"

  local subcategory_path
  for subcategory_name in "$@"
  do
    subcategory_path="$(_categoryfs_mkdir_subcategory "$category_name" "$subcategory_name")"
    echo "+ $subcategory_path (sub-category)"
  done
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
