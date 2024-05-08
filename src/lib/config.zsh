# Marmot configuration

function _config_metadata_init() {
  local directory
  directory="$1"

  local meta_repo_file
  meta_repo_file="$directory/meta-repo.json"

  local template
  template=$(cat <<'EOF'
{
  meta_repo: {
    categories: [],
    repositories: []
  },
  version: $version
}
EOF
  )

  jq > "$meta_repo_file" \
    --arg version 0.3.2 \
    --null-input \
    "$template"
}

## .categories

function _config_add_categories() {
  local config_file
  config_file="$1"
  shift 1

  local category_name
  category_name="$1"
  shift 1

  local categories
  categories=("$(__config_category_name_to_json "$category_name")")

  __config_subcategory_names_to_json "$category_name" "$@"
  categories+=("${reply[@]}")

  local categories_as_json
  categories_as_json="$(_json_to_array "${categories[@]}")"
  _json_update "$config_file" ".meta_repo.categories += ${categories_as_json}"
}

function _config_category_fullnames() {
  local config_file
  config_file="$1"

  jq < "$config_file" \
    -r '.meta_repo.categories[]?.full_name'
}

function __config_category_names_to_json() {
  local categories category_json

  categories=()
  for name in "$@"
  do
    category_json="$(__config_category_name_to_json "$name")"
    categories+=("$category_json")
  done

  _json_to_array "${categories[@]}"
}

function __config_category_name_to_json() {
  local name parent_name
  name="$1"
  parent_name="$2"

  if [[ -n "$parent_name" ]]
  then
    jo -- \
      "full_name=$parent_name/$name" \
      "name=$name" \
      "parent_name=$parent_name"
  else
    jo -- \
      "full_name=$name" \
      "name=$name" \
      -s 'parent_name='
  fi
}

function __config_subcategory_names_to_json() {
  local parent_category_name
  parent_category_name="$1"
  shift 1

  local subcategories
  subcategories=()
  for name in "$@"
  do
    subcategory_json="$(__config_category_name_to_json "$name" "$parent_category_name")"
    subcategories+=("$subcategory_json")
  done

  reply=("${subcategories[@]}")
}

## .repositories

function _config_add_repositories() {
  local config_file
  config_file="$1"
  shift 1

  local repositories_as_json
  repositories_as_json=$(__config_repository_paths_to_json "$@")
  _json_update "$config_file" ".meta_repo.repositories += ${repositories_as_json}"
}

function _config_repository_paths() {
  local config_file
  config_file="$1"

  # Treat lack of JSON fields as empty rather than as an error
  # https://github.com/jqlang/jq/issues/354#issuecomment-43147898
  jq < "$config_file" \
    -r '.meta_repo.repositories[]?.path'
}

# __ prefix indicates private access - e.g. implementation details not meant to cross the interface

function __config_repository_paths_to_json() {
  local repositories repository_json

  repositories=()
  for repository_path in "$@"
  do
    repository_json=$(__config_repository_path_to_json "$repository_path")
    repositories+=("$repository_json")
  done

  _json_to_array "${repositories[@]}"
}

function __config_repository_path_to_json() {
  local repo_path
  repo_path="$1"
  jo -- "path=$repo_path"
}
