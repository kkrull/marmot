
function _category_name_from_id() {
  local category_or_subcategory="$1"
  IFS='/' read -r category_name subcategory_name <<< "$category_or_subcategory"
  echo "$category_name"
}

function _category_subcategory_from_id() {
  local category_or_subcategory="$1"
  IFS='/' read -r category_name subcategory_name <<< "$category_or_subcategory"
  echo "$subcategory_name"
}
