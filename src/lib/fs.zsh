# Marmot filesystem

## Categories

function _fs_add_repository_link() {
  local category_or_subcategory repository_path
  category_or_subcategory="$1"
  repository_path="$2"

  local link_path repository_name category_path
  category_path="$(_fs_category_path "$category_or_subcategory")"
  repository_name="${repository_path:t}"
  link_path="$category_or_subcategory/$repository_name"
  (cd "$category_path" && ln -f -s "$repository_path" "$repository_name")
  echo "$link_path"
}

function _fs_category_path() {
  local category_or_subcategory
  category_or_subcategory="$1"

  echo "$(_fs_metarepo_home)/$category_or_subcategory"
}

function _fs_make_category_path() {
  local category_or_subcategory
  category_or_subcategory="$1"

  category_path="$(_fs_category_path "$category_or_subcategory")"
  mkdir -p "$category_path"
  echo "$category_path"
}

function _fs_make_subcategory_path() {
  local category subcategory
  category="$1"
  subcategory="$2"
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
  if [[ -n "$MARMOT_META_REPO" ]]
  then
    echo "$MARMOT_META_REPO"
  else
    echo "$HOME/meta"
  fi
}

function _fs_marmot_version() {
  cat "$_MARMOT_HOME/version"
}
