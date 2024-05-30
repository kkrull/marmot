
# Category representation on the filesystem

function _categoryfs_add_repository_link() {
  local category_or_subcategory="$1" some_repo_path="$2"
  [[ -z "$category_or_subcategory" || -z "$some_repo_path" ]] && exit 0

  local category_path full_repo_path repository_name
  category_path="$(__categoryfs_local_path "$category_or_subcategory")"
  full_repo_path="$(_fs_normalize_repo_path "$some_repo_path")"
  repository_name="${some_repo_path:t}"
  (cd "$category_path" \
    && ln -f -s "$full_repo_path" "$repository_name")

  local link_path
  link_path="$category_or_subcategory/$repository_name"
  echo "$link_path"
}

function __categoryfs_local_path() {
  local category_or_subcategory="$1"
  echo "$(_fs_metarepo_home)/$category_or_subcategory"
}

function _categoryfs_mkdir() {
  local category_or_subcategory="$1"
  [[ -z "$category_or_subcategory" ]] && exit 0

  category_path="$(__categoryfs_local_path "$category_or_subcategory")"
  mkdir -p "$category_path"
  echo "$category_path"
}

function _categoryfs_mkdir_subcategory() {
  local category="$1" subcategory="$2"
  [[ -z "$category" || -z "$subcategory" ]] && exit 0

  _categoryfs_mkdir "$category/$subcategory"
}

function _categoryfs_rm_repository_link() {
  local category_or_subcategory="$1" some_repo_path="$2"
  [[ -z "$category_or_subcategory" || -z "$some_repo_path" ]] && exit 0

  local category_path
  category_path="$(__categoryfs_local_path "$category_or_subcategory")"

  local full_repo_path repository_name
  full_repo_path="$(_fs_normalize_repo_path "$some_repo_path")"
  repository_name="${full_repo_path:t}"

  local link_path_full
  link_path_full="$category_path/$repository_name"
  rm -f "$link_path_full"

  local link_path_relative="$category_or_subcategory/$repository_name"
  echo "$link_path_relative"
}
