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
    --arg version 0.3.1 \
    --null-input \
    "$template"
}

## .categories

function _config_category_names() {
  local config_file
  config_file="$1"

  jq < "$config_file" \
    -r '.meta_repo.categories[]?.name'
}

## .repositories

function _config_add_repositories() {
  local config_file repositories
  config_file="$1"
  shift 1

  repositories="$*"
  _json_update "$config_file" ".meta_repo.repositories += ${repositories}"
}

function _config_repository_paths() {
  local config_file
  config_file="$1"

  # Treat lack of JSON fields as empty rather than as an error
  # https://github.com/jqlang/jq/issues/354#issuecomment-43147898
  jq < "$config_file" \
    -r '.meta_repo.repositories[]?.path'
}

function _config_paths_to_repositories() {
  local repositories repository

  repositories=()
  for repository_path in "$@"
  do
    repository=$(_config_path_to_repository "$repository_path")
    repositories+=("$repository")
  done

  _json_to_array "${repositories[@]}"
}

function _config_path_to_repository() {
  local repo_path
  repo_path="$1"
  echo "{ \"path\": \"$repo_path\" }"
}
