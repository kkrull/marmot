
# Marmot metadata

function _config_init() {
  local directory="$1"

  _jq_create "$directory/meta-repo.json" \
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
