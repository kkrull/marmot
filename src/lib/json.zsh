## JSON

function to_json_array() {
  if [[ $# == 0 ]]
  then
    echo "[]"
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
