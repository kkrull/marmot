# Marmot configuration

function create_meta_repo_config() {
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

function category_names() {
  local config_file
  config_file="$1"

  jq < "$config_file" \
    -r '.meta_repo.categories[]?.name'
}

## .repositories

function add_repositories() {
  local config_file repositories
  config_file="$1"
  shift 1

  repositories="$*"
  jq_update "$config_file" ".meta_repo.repositories += ${repositories}"
}

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
