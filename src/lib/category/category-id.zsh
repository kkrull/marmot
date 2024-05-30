
# Category IDs

function _categoryid_category() {
  local category_or_subcategory="$1"
  IFS='/' read -r category_name subcategory_name <<< "$category_or_subcategory"
  echo "$category_name"
}

function _categoryid_subcategory() {
  local category_or_subcategory="$1"
  IFS='/' read -r category_name subcategory_name <<< "$category_or_subcategory"
  echo "$subcategory_name"
}
