
# Metadata filesystem

function _metafs_init() {
  local directory="$1"

  __metafs_init_metadata_file "$directory/meta-repo.json"
  __metafs_init_localdata_file "$directory/meta-repo-local.json"
}

## private

function __metafs_init_localdata_file() {
  local data_file="$1"

  _jq_create "$data_file" \
    --arg version "$(_fs_marmot_version)" \
    --null-input \
    --sort-keys <<-'EOF'
{
  meta_repo: {
    repositories: [],
    updated: now | todate
  },
  version: $version
}
EOF
}

function __metafs_init_metadata_file() {
  local data_file="$1"

  _jq_create "$data_file" \
    --arg version "$(_fs_marmot_version)" \
    --null-input \
    --sort-keys <<-'EOF'
{
  meta_repo: {
    categories: [],
    updated: now | todate
  },
  version: $version
}
EOF
}
