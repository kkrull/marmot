# Marmot configuration

function _config_init() {
  local directory
  directory="$1"

  _json_jq_create "$directory/meta-repo.json" \
    --arg version "$(_fs_marmot_version)" \
    --null-input \
    --sort-keys <<-'EOF'
{
  meta_repo: {
    categories: [],
    repositories: [],
    updated: now | todate
  },
  version: $version
}
EOF
}

## .categories

function _config_add_categories() {
  local config_file category_name
  config_file="$1"
  category_name="$2"

  local categories
  categories=("$(__config_category_from_name "$category_name")")
  for subcategory_name in "${@:3}"
  do
    categories+=("$(__config_category_from_name "$subcategory_name" "$category_name")")
  done

  _json_jq_update "$config_file" \
    --argjson categories "$(jo -a "${categories[@]}")" \
    --sort-keys <<-'EOF'
    . | .meta_repo.categories |= (. + $categories | unique_by(.full_name))
      | .meta_repo.updated |= (now | todate)
EOF
}

function _config_add_repositories_to_category() {
  local config_file category_full_name
  config_file="$1"
  category_full_name="$2"

  declare -a repository_paths=()
  for some_repo_path in "${@:3}"
  do
    repository_paths+=("$(__config_normalize_path "$some_repo_path")")
  done

  # Complex assignment to update one element in the array without deleting the others
  # https://jqlang.github.io/jq/manual/#complex-assignments
  _json_jq_update "$config_file" \
    --arg category_full_name "$category_full_name" \
    --argjson repository_paths "$(jo -a "${repository_paths[@]}")" \
    --sort-keys <<-'EOF'
    . | (.meta_repo.categories[]
          | select(.full_name == $category_full_name)
          | .repository_paths)
        |= (. + $repository_paths | unique)
      | .meta_repo.updated |= (now | todate)
EOF
}

function _config_category_fullnames() {
  local config_file
  config_file="$1"

  jq -r \
    '.meta_repo.categories[]?.full_name' \
    "$config_file"
}

# private

function __config_category_from_name() {
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

## .repositories

function _config_add_repositories() {
  local config_file="$1"
  declare -a repositories=()

  local repo_path repository
  for some_path in "${@:2}"
  do
    repo_path="$(__config_normalize_path "$some_path")"
    repository="$(__config_repository_from_path "$repo_path")"
    repositories+=("$repository")
  done

  _json_jq_update "$config_file" \
    --argjson repositories "$(jo -a "${repositories[@]}")" \
    --sort-keys <<-'EOF'
    .
      | .meta_repo.repositories |= (. + $repositories | unique_by(.path))
      | .meta_repo.updated |= (now | todate)
EOF
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
  done <<-EOF
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
  filter=$(cat <<-'EOF'
    .meta_repo.categories[]
      | select(.full_name == $full_name)
      | .repository_paths[]?
EOF
  )

  jq \
    --arg full_name "$category_or_subcategory" \
    -r \
    "$filter" "$config_file"
}

function _config_remove_repositories() {
  declare config_file="$1"
  declare remove_paths=("${@:2}")

  _json_jq_update "$config_file" \
    --argjson repository_paths "$(jo -a "${remove_paths[@]}")" \
    --sort-keys <<-'EOF'
    . | .meta_repo.categories[].repository_paths? -= $repository_paths
      | .meta_repo.repositories[]
        |= del(select(
                .path
                | in($repository_paths
                      | map(. as $elem | { key: $elem, value: 1 })
                      | from_entries)))
        | del(..|nulls)
      | .meta_repo.updated |= (now | todate)
EOF
}

# private

function __config_normalize_path() {
  local some_path="$1"
  local absolute_path="${some_path:A}"
  echo "${absolute_path%%/.git}"
}

function __config_repository_from_path() {
  local repo_path
  repo_path="$1"
  jo -- "path=$repo_path"
}
