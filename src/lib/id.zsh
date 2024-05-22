
## Category IDs

function _id_category_name() {
  local category_or_subcategory
  category_or_subcategory="$1"
  IFS='/' read -r category_name <<< "$category_or_subcategory"
  echo "$category_name"
}

function _id_subcategory_name() {
  local category_or_subcategory
  category_or_subcategory="$1"
  IFS='/' read -r category_name subcategory_name <<< "$category_or_subcategory"
  echo "$subcategory_name"
}
