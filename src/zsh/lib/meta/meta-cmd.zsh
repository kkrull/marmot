
# Metadata commands

function _metacmd_init() {
  local directory
  directory="$(_fs_metadata_dir)"

  if [[ -d "$directory" ]]
  then
    printf "Meta repo already exists: %s" "$directory"
    exit 1
  fi

  mkdir -p "$directory"
  _metafs_init "$directory"
  echo "Initialized meta repository at $directory"
}

function _metacmd_show_home() {
  _fs_metarepo_home
}
