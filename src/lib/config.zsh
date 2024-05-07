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
  local config_file repositories_as_json #TODO KDK: Broken abstraction
  config_file="$1"
  shift 1

  repositories_as_json="$*"
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

function __config_paths_to_repositories_json() {
  local repositories repository_json

  repositories=()
  for repository_path in "$@"
  do
    repository_json=$(__config_path_to_repository_json "$repository_path")
    repositories+=("$repository_json")
  done

  _json_to_array "${repositories[@]}"
}

function __config_path_to_repository_json() {
  local repo_path
  repo_path="$1"
  echo "{ \"path\": \"$repo_path\" }"
}
