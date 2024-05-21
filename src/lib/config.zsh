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
    "version=$(_fs_marmot_version)"
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

function _config_add_repositories_to_category() {
  local config_file category_full_name repository_paths
  config_file="$1"
  category_full_name="$2"

  repository_paths=()
  for repo_path in "${@:3}"
  do
    repository_paths+=("${repo_path:A}")
  done

  # Complex assignment to update one element in the array without deleting the others
  # https://jqlang.github.io/jq/manual/#complex-assignments
  local filter
  filter=$(cat <<EOF
    (.meta_repo.categories[]
      | select(.full_name == "$category_full_name")
      | .repository_paths)
      += $(jo -a "${repository_paths[@]}")
EOF
  )

  _json_jq_update "$config_file" "$filter"
}

function _config_category_fullnames() {
  local config_file
  config_file="$1"

  jq -r \
    '.meta_repo.categories[]?.full_name' \
    "$config_file"
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
      "parent_name=$parent_name" \
      'repository_paths=[]'
  else
    jo -- \
      "full_name=$name" \
      "name=$name" \
      -s 'parent_name=' \
      'repository_paths=[]'
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
  jq -r \
    '.meta_repo.repositories[]?.path' \
    "$config_file"
}

function _config_repository_paths_reply() {
  local config_file
  config_file="$1"

  reply=()
  while read -r line
  do
    reply+=("$line")
  done <<EOF
    $(jq < "$config_file" \
      --raw-output \
      '.meta_repo.repositories[]?.path')
EOF
}

function _config_repository_paths_in_category() {
  local config_file category_or_subcategory
  config_file="$1"
  category_or_subcategory="$2"

  local filter
  filter=$(cat <<EOF
    .meta_repo.categories[]
      | select(.full_name == "$category_or_subcategory")
      | .repository_paths[]?
EOF
  )

  jq -r "$filter" "$config_file"
}

function _config_remove_repositories() {
  declare config_file="$1" remove_paths=("${@:2}")

  declare remove_paths_json
  remove_paths_json="$(jo -a "${remove_paths[@]}")"

  declare filter
  filter=$(cat <<'EOF'
. | .meta_repo.categories[].repository_paths? -= $remove_paths_json
  | .meta_repo.repositories[]
    |= del(select(
            .path
            | in($remove_paths_json
                  | map(. as $elem | { key: $elem, value: 1 })
                  | from_entries)))
    | del(..|nulls)
  | .meta_repo.updated
    |= (now | todate)
EOF
)

  _json_jq_update "$config_file" --argjson remove_paths_json "$remove_paths_json" "$filter"
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
