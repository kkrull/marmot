# Marmot filesystem

## Categories

function _fs_add_repository_link() {
  local category_or_subcategory="$1" repository_path="$2"
  [[ -z "$category_or_subcategory" || -z "$repository_path" ]] && exit 0

  local link_path repository_name category_path
  category_path="$(_fs_category_path "$category_or_subcategory")"
  repository_name="${repository_path:t}"
  link_path="$category_or_subcategory/$repository_name"
  (cd "$category_path" && ln -f -s "$repository_path" "$repository_name")
  echo "$link_path"
}

function _fs_category_path() {
  local category_or_subcategory="$1"
  echo "$(_fs_metarepo_home)/$category_or_subcategory"
}

function _fs_make_category_path() {
  local category_or_subcategory="$1"
  [[ -z "$category_or_subcategory" ]] && exit 0

  category_path="$(_fs_category_path "$category_or_subcategory")"
  mkdir -p "$category_path"
  echo "$category_path"
}

function _fs_make_subcategory_path() {
  local category="$1" subcategory="$2"
  [[ -z "$category" || -z "$subcategory" ]] && exit 0

  _fs_make_category_path "$category/$subcategory"
}

## Configuration

function _fs_metadata_dir() {
  echo "$(_fs_metarepo_home)/.marmot"
}

function _fs_metadata_file() {
  echo "$(_fs_metarepo_home)/.marmot/meta-repo.json"
}

function _fs_metarepo_home() {
  echo "${MARMOT_META_REPO-$HOME/meta}"
}

function _fs_marmot_version() {
  cat "$_MARMOT_HOME/version"
}

## Repositories

function _fs_normalize_repo_path() {
  local some_repo_path="$1"
  local absolute_path="${some_repo_path:A}"
  echo "${absolute_path%%/.git}"
}
