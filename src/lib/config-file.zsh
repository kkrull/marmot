# Marmot configuration

## .repositories

function repository_paths() {
  local config_file
  config_file="$1"

  # Treat lack of JSON fields as empty rather than as an error
  # https://github.com/jqlang/jq/issues/354#issuecomment-43147898
  jq < "$config_file" \
    -r '.meta_repo.repositories[]?.path'
}

function to_marmot_repositories() {
  local repositories repository

  repositories=()
  for repository_path in "$@"
  do
    repository=$(to_marmot_repository "$repository_path")
    repositories+=("$repository")
  done

  to_json_array "${repositories[@]}"
}

function to_marmot_repository() {
  local repo_path
  repo_path="$1"
  echo "{ \"path\": \"$repo_path\" }"
}
