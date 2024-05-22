
## Category IDs

# https://zsh.sourceforge.io/Doc/Release/Expansion.html#Parameter-Expansion

function _id_category_name() {
  declare -a category_words=()
  category_words=(${=category_or_subcategory//\// })
  echo "${category_words[1]}"
}

function _id_subcategory_name() {
  declare -a category_words=()
  category_words=(${=category_or_subcategory//\// })
  echo "${category_words[2]}"
}
