# Tools

Tools used by this project.

## EditorConfig

- Files:
  - Configuration: `.editorconfig`
- Interactions:
  - `pre-commit`: checks files with this check.

## GitHub Actions

- Documentation:
  - [Jobs](./cicd-jobs.md)
- Files:
  - Job definition: `.github/workflows/`
- Interactions:
  - `pre-commit`: GitHub Actions runs `pre-commit` to do its checks.

## `make`

- Files:
  - Task definition: `Makefile`

## `markdownlint`

- Files:
  - Configuration: `.markdownlint.json`
- Interactions:
  - `pre-commit`: checks files with this check.

## `pre-commit`

- Files:
  - Configuration: `.pre-commit.yaml`
- Usage:
  - Install Git hooks: `pre-commit install`
  - Run checks: `pre-commit run --all`

## `shellcheck`

- Files:
  - Configuration: `.shellheckrc`
- Interactions:
  - `pre-commit`: checks files with this check.
