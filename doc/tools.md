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

## Homebrew

- Documentation:
  - <https://github.com/Homebrew/homebrew-bundle>
- Files:
  - `Brewfile`
  - `Brewfile.lock.json`

### Usage

Check:

```sh
brew bundle check
```

Update `Brewfile`

```sh
# Dumps all packages installed on the system; requires editing
brew dump --file=./Brewfile
```

Install tools needed by this repository:

```sh
brew bundle install
```

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
