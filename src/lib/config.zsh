# Marmot configuration

function _config_metadata_init() {
  local directory
  directory="$1"

  local meta_repo_file
  meta_repo_file="$directory/meta-repo.json"

  jo > "$meta_repo_file" \
    -p \
    -- \
    "meta_repo=$(jo -- 'categories=[]' 'repositories=[]')" \
    'version=0.4'
}

## .categories

function _config_add_categories() {
  local config_file category_name subcategory_names
  config_file="$1"
  category_name="$2"
  subcategory_names=("${@:3}")

  local categories
  categories=("$(__config_category_name_to_json "$category_name")")
  __config_subcategory_names "$category_name" "${subcategory_names[@]}"
  categories+=("${reply[@]}")

  _json_jq_update "$config_file" ".meta_repo.categories += $(jo -a "${categories[@]}")"
}

function _config_category_fullnames() {
  local config_file
  config_file="$1"

  jq < "$config_file" \
    -r '.meta_repo.categories[]?.full_name'
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

function __config_subcategory_names() {
  local parent_category_name subcategory_names
  parent_category_name="$1"
  subcategory_names=("${@:2}")

  local subcategories
  subcategories=()
  for name in "${subcategory_names[@]}"
  do
    subcategory_json="$(__config_category_name_to_json "$name" "$parent_category_name")"
    subcategories+=("$subcategory_json")
  done

  reply=("${subcategories[@]}")
}

## .repositories

function _config_add_repositories() {
  local config_file repository_paths
  config_file="$1"
  repository_paths=("${@:2}")

  local repositories_as_json
  repositories_as_json=$(__config_repository_paths_to_json "${repository_paths[@]}")
  _json_jq_update "$config_file" ".meta_repo.repositories += ${repositories_as_json}"
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

  jo -a "${repositories[@]}"
}

function __config_repository_path_to_json() {
  local repo_path
  repo_path="$1"
  jo -- "path=$repo_path"
}
