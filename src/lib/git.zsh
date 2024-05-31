
# Git interface

function _git_ssh_url() {
  local repo_path="$1"

  # TODO KDK: Check for 0 or 2+ remotes
  local remote_name
  remote_name="$(cd "$repo_path" && git remote)"
  (cd "$repo_path" && git remote get-url "$remote_name")
}
