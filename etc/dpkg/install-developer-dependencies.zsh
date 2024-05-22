#!/usr/bin/env zsh

emulate -LR zsh
set -euo pipefail

apt install -y \
  fswatch \
  pandoc \
  pre-commit
