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

  local category_as_json
  category_as_json="$(__config_make_category "$@")"
  _json_update "$config_file" ".meta_repo.categories += [${category_as_json}]"
}

function _config_category_names() {
  local config_file
  config_file="$1"

  jq < "$config_file" \
    -r '.meta_repo.categories[]?.name'
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
  local name
  name="$1"
  echo "{ \"name\": \"$name\" }"
}

function __config_make_category() {
  local category_name subcategory_names
  category_name="$1"
  subcategory_names=("${@:2}")

  local template
  template=$(cat <<'EOF'
{
  name: $category_name,
  categories: $subcategories
}
EOF
  )

  jq \
    --arg category_name "$category_name" \
    --argjson subcategories "$(__config_category_names_to_json "${subcategory_names[@]}")" \
    --null-input \
    "$template"
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
  echo "{ \"path\": \"$repo_path\" }"
}
