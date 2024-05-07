# Marmot paths

## Categories

function category_path() {
  local category_or_subcategory
  category_or_subcategory="$1"

  echo "$(meta_repo_home)/$category_or_subcategory"
}

## Configuration

function meta_repo_config_file() {
  echo "$(meta_repo_home)/.marmot/meta-repo.json"
}

function meta_repo_data() {
  echo "$(meta_repo_home)/.marmot"
}

## Home

function meta_repo_home() {
  local meta_home
  meta_home="${PWD:A}"

  echo "$meta_home"
}
