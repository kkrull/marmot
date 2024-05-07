# Marmot filesystem

## Categories

function add_repository_link_to_category() {
  local category_or_subcategory repository_path
  category_or_subcategory="$1"
  repository_path="$2"

  local link_path repository_name the_category_path
  the_category_path="$(category_path "$category_or_subcategory")"
  repository_name="${repository_path:t}"
  link_path="$category_or_subcategory/$repository_name"
  (cd "$the_category_path" && ln -s "$repository_path" "$repository_name")
  echo "$link_path"
}

function category_path() {
  local category_or_subcategory
  category_or_subcategory="$1"

  echo "$(meta_repo_home)/$category_or_subcategory"
}

function make_category_path() {
  local category_or_subcategory
  category_or_subcategory="$1"

  the_category_path="$(category_path "$category_or_subcategory")"
  mkdir -p "$the_category_path"
  echo "$the_category_path"
}

function make_subcategory_path() {
  local category subcategory
  category="$1"
  subcategory="$2"
  make_category_path "$category/$subcategory"
}

## Configuration

function meta_repo_config_file() {
  echo "$(meta_repo_home)/.marmot/meta-repo.json"
}

function meta_repo_data() {
  echo "$(meta_repo_home)/.marmot"
}

## Home

function meta_repo_home() {
  local meta_home
  meta_home="${PWD:A}"

  echo "$meta_home"
}
