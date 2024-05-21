# Marmot JSON processing

function _json_jq_update() {
  local json_file="$1"
  shift 1

  local tmp_file
  tmp_file=$(mktemp)
  cp "$json_file" "$tmp_file"
  jq "$@" "$tmp_file" > "$json_file"
  rm -f "$tmp_file"
}

function _json_jq_heredoc() {
  local json_file="$1"
  shift 1
  # local args="${@:2}"

  declare -a filter_lines=()
  while read -r line
  do
    filter_lines+=("$line")
  done

  echo "_json_jq_heredoc filter[${#filter_lines}]:"
  for line in "${filter_lines[@]}"
  do
    echo "$line"
  done

  local tmp_file
  tmp_file=$(mktemp)
  cp "$json_file" "$tmp_file"
  jq "$@" "${filter_lines[*]}" "$tmp_file" # > "$json_file"
  rm -f "$tmp_file"
}
