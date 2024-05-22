
## Category IDs

# https://zsh.sourceforge.io/Doc/Release/Expansion.html#Parameter-Expansion

function _string_category_name() {
  declare -a category_words=()
  category_words=(${=category_or_subcategory//\// })
  echo "${category_words[1]}"
}

function _string_subcategory_name() {
  declare -a category_words=()
  category_words=(${=category_or_subcategory//\// })
  echo "${category_words[2]}"
}
