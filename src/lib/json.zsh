# Marmot JSON processing

function _json_update() {
  local json_file="$1"
  shift 1

  tmp_file=$(mktemp)
  cp "$json_file" "$tmp_file"
  jq < "$tmp_file" > "$json_file" "$*"
  rm -f "$tmp_file"
}

function _json_to_array() {
  if [[ $# == 0 ]]
  then
    echo '[]'
  elif [[ $# == 1 ]]
  then
    echo "[$1]"
  else
    printf "[%s" "$1"
    for element in "${@:2}"
    do
      printf ", %s" "$element"
    done

    printf ']'
  fi
}
