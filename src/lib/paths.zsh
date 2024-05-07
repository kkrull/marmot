# Marmot paths

function meta_repo_config_file() {
  local meta_home

  meta_home="${PWD:A}"
  echo "$meta_home/.marmot/meta-repo.json"
}
