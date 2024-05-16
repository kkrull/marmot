# Tools

Tools used by some part of this project's build, deployment, or development processes.

## EditorConfig

<https://editorconfig.org/>

### Files

- `.editorconfig`: Configuration

### Interactions

- [`pre-commit`](#pre-commit): includes hooks to ensure files comply with EditorConfig

## GitHub Actions

<https://docs.github.com/en/actions>

### Documentation

- [CI/CD Jobs](./cicd-jobs.md)
- Workflow Syntax:
  <https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions>

### Files

- `.github/workflows/`: Job definitions.

### Interactions

- [`pre-commit`](#pre-commit): GitHub Actions contains jobs that run `pre-commit` checks

## Homebrew

<https://brew.sh/>

### Documentation

- `homebrew-bundle`: <https://github.com/Homebrew/homebrew-bundle>

### Files

- `Brewfile.developer` and `Brewfile.developer.lock.json`: packages for developers, to work on this
  project.
- `Brewfile.user` and `Brewfile.user.lock.json`: packages for end users, to run the programs built
  from these sources.

## GNU Make

<https://www.gnu.org/software/make/>

### Documentation

- GNU Make Manual: <https://www.gnu.org/software/make/manual/>

### Files

- `Makefile`: targets to build everything and tasks to help set up your environment
- `man/Makefile`: targets to build manuals
- `src/Makefile`: targets to build programs

## `markdownlint-cli2`

<https://github.com/DavidAnson/markdownlint-cli2>

### Documentation

- Configuration JSON schema:
  <https://github.com/DavidAnson/markdownlint-cli2#markdownlint-cli2jsonc>
- `pre-commit` hook: <https://github.com/DavidAnson/markdownlint-cli2?tab=readme-ov-file#pre-commit>

### Files

- `.markdownlint-cli2.jsonc`: configuration

### Interactions

- [`pre-commit`](#pre-commit): includes hooks to ensure files comply with linting rules

## `pre-commit`

<https://pre-commit.com/>

### Files

- `.pre-commit.yaml`: configures repository sources and hooks to run.  May also contain
  configuration not already included in any hook-specific configuration files (e.g.
  `.editorconfig`).

### Interactions

- [EditorConfig](#editorconfig): checked by `pre-commit`
- [GNU Make](#gnu-make): contains [tasks](./task-automation.md#pre-commit-tasks) to install Git
  hooks and update `pre-commit` repositories
- [`markdownlint-cli2`](#markdownlint-cli2): run by `pre-commit`
- [ShellCheck](#shellcheck): run by `pre-commit`

## ShellCheck

<https://github.com/koalaman/shellcheck>

### Documentation

- Checks (open Pages dropdown): <https://github.com/koalaman/shellcheck/wiki/Checks>
- Directives: <https://github.com/koalaman/shellcheck/wiki/Directive>
- Ignoring findings: <https://www.shellcheck.net/wiki/Ignore>

### Files

- `.shellcheckrc`: configuration

### Interactions

- [`pre-commit`](#pre-commit): includes hooks to ensure files comply with ShellCheck
