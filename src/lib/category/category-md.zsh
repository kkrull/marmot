
# Category metadata

function _categorymd_add_repositories_as_local_path() {
  local config_file="$1" category_full_name="$2" ; shift 2

  # Remove array elements matching '' to avoid entering a world of pain
  # https://zsh.sourceforge.io/Doc/Release/Expansion.html#Parameter-Expansion
  declare -a repository_paths=()
  for some_repo_path in "${@:#}"
  do
    repository_paths+=("$(_repofs_normalize_path "$some_repo_path")")
  done

  # Complex assignment to update one element in the array without deleting the others
  # https://jqlang.github.io/jq/manual/#complex-assignments
  _jq_update "$config_file" \
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

function _categorymd_create() {
  local config_file="$1" category_name="$2" ; shift 2

  local categories
  categories=("$(__categorymd_category_from_name "$category_name")")
  for subcategory_name in "${@:#}"
  do
    categories+=("$(__categorymd_category_from_name "$subcategory_name" "$category_name")")
  done

  _jq_update "$config_file" \
    --argjson categories "$(jo -a "${categories[@]}")" \
    --sort-keys <<-'EOF'
    . | .meta_repo.categories |= (. + $categories | unique_by(.full_name))
      | .meta_repo.updated |= (now | todate)
EOF
}

function _categorymd_full_names() {
  local config_file="$1"

  jq -r \
    '.meta_repo.categories[]?.full_name' \
    "$config_file"
}

function _categorymd_remove_repositories_as_local_paths() {
  local config_file="$1" category_full_name="$2" ; shift 2

  declare -a repository_paths=()
  for some_repo_path in "${@:#}"
  do
    repository_paths+=("$(_repofs_normalize_path "$some_repo_path")")
  done

  _jq_update "$config_file" \
    --arg category_full_name "$category_full_name" \
    --argjson repository_paths "$(jo -a "${repository_paths[@]}")" \
    --sort-keys <<-'EOF'
    . | (.meta_repo.categories[]
          | select(.full_name == $category_full_name)
          | .repository_paths)
        |= (. - $repository_paths | unique)
      | .meta_repo.updated |= (now | todate)
EOF
}

## private

function __categorymd_category_from_name() {
  local name="$1" parent_name="${2-}"

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
