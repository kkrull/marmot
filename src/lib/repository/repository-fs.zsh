
## Repositories

function _fs_normalize_repo_path() {
  local some_repo_path="$1"
  local absolute_path="${some_repo_path:A}"
  echo "${absolute_path%%/.git}"
}
