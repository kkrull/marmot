## Marmot configuration

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
