# Marmot JSON processing

function _json_jq_create() {
  local json_file="$1" ; shift 1

  declare -a filter_lines=()
  while read -r line
  do
    filter_lines+=("$line")
  done

  jq "$@" "${filter_lines[*]}" > "$json_file"
}

function _json_jq_update() {
  local json_file="$1" ; shift 1
  declare -a filter_lines=()
  while read -r line
  do
    filter_lines+=("$line")
  done

  local tmp_file
  tmp_file="$(mktemp)"
  cp "$json_file" "$tmp_file"

  jq "$@" "${filter_lines[*]}" "$tmp_file" > "$json_file"
  rm -f "$tmp_file"
}
