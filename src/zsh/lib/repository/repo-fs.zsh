
# Repository filesystem

function _repofs_normalize_path() {
  local some_repo_path="$1"
  local absolute_path="${some_repo_path:A}"
  echo "${absolute_path%%/.git}"
}
