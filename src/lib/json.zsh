# Marmot JSON processing

function _json_jq_update() {
  local json_file="$1"
  shift 1

  tmp_file=$(mktemp)
  cp "$json_file" "$tmp_file"
  jq "$@" "$tmp_file" > "$json_file"
  rm -f "$tmp_file"
}
