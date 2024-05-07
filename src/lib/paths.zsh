# Marmot paths

function meta_repo_config_file() {
  echo "$(meta_repo_home)/.marmot/meta-repo.json"
}

function meta_repo_data() {
  echo "$(meta_repo_home)/.marmot"
}

function meta_repo_home() {
  local meta_home

  meta_home="${PWD:A}"
  echo "$meta_home"
}
