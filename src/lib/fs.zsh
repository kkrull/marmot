
# Marmot filesystem

## Configuration

function _fs_metadata_dir() {
  echo "$(_fs_metarepo_home)/.marmot"
}

function _fs_metarepo_home() {
  echo "${MARMOT_META_REPO-$HOME/meta}"
}

function _fs_marmot_version() {
  cat "$_MARMOT_HOME/version"
}
