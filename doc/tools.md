# Tools

Tools used by some part of this project's build, deployment, or development processes.

## EditorConfig

<https://editorconfig.org/>

### Files

- `.editorconfig`: Configuration

### Interactions

- [`pre-commit`](#pre-commit) includes hooks to ensure files comply with EditorConfig

## `fswatch`

<https://github.com/emcrisostomo/fswatch>

### Interactions

- [GNU Make](#gnu-make) includes tasks that use `fswatch`

## GitHub Actions

<https://docs.github.com/en/actions>

### Documentation

- [CI/CD Jobs](./cicd-jobs.md)
- Workflow Syntax:
  <https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions>

### Files

- `.github/workflows/`: Job definitions.

### Interactions

- [`pre-commit`](#pre-commit) checks are also run during some jobs

## GNU Make

<https://www.gnu.org/software/make/>

### Documentation

- GNU Make Manual: <https://www.gnu.org/software/make/manual/>
- Standard Targets:
  <https://www.gnu.org/software/make/manual/html_node/Standard-Targets.html#Standard-Targets>

### Files

- `Makefile`: targets to build everything and tasks to help set up your environment
- `man/Makefile`: targets to build manuals
- `src/Makefile`: targets to build programs

## Homebrew

<https://brew.sh/>

### Documentation

- `homebrew-bundle`: <https://github.com/Homebrew/homebrew-bundle>

### Files

- `Brewfile.developer` and `Brewfile.developer.lock.json`: packages for developers, to work on this
  project.
- `Brewfile.user` and `Brewfile.user.lock.json`: packages for end users, to run the programs built
  from these sources.

## `jo`

<https://jpmens.net/2016/03/05/a-shell-command-to-create-json-jo/>

### Documentation

- GitHub: <https://github.com/jpmens/jo>
- Installation: <https://github.com/jpmens/jo?tab=readme-ov-file#install>

## `jq`

<https://jqlang.github.io/jq/>

### Documentation

- Manual: <https://jqlang.github.io/jq/manual/v1.7/>

## `markdownlint-cli2`

<https://github.com/DavidAnson/markdownlint-cli2>

### Documentation

- Configuration JSON schema:
  <https://github.com/DavidAnson/markdownlint-cli2#markdownlint-cli2jsonc>
- `pre-commit` hook: <https://github.com/DavidAnson/markdownlint-cli2?tab=readme-ov-file#pre-commit>

### Files

- `.markdownlint-cli2.jsonc`: configuration

### Interactions

- [`pre-commit`](#pre-commit) includes hooks to ensure files comply with linting rules

## Pandoc

<https://pandoc.org/>

### Documentation

- GitHub Actions for Pandoc: <https://pandoc.org/installing.html#github-actions>
- Installation: <https://pandoc.org/installing.html>

### Interactions

- [GNU Make](#gnu-make) runs Pandoc to convert command documentation to manuals

## `pre-commit`

<https://pre-commit.com/>

### Files

- `.pre-commit.yaml`: configures repository sources and hooks to run.  May also contain
  configuration not already included in any hook-specific configuration files (e.g.
  `.editorconfig`).

### Interactions

- [EditorConfig](#editorconfig) is checked by `pre-commit`
- [GNU Make](#gnu-make) contains [tasks](./task-automation.md#pre-commit-targets) to install Git
  hooks and update `pre-commit` repositories
- [`markdownlint-cli2`](#markdownlint-cli2) is run by `pre-commit`
- [ShellCheck](#shellcheck) is run by `pre-commit`

## ShellCheck

<https://github.com/koalaman/shellcheck>

### Documentation

- Checks (open Pages dropdown): <https://github.com/koalaman/shellcheck/wiki/Checks>
- Directives: <https://github.com/koalaman/shellcheck/wiki/Directive>
- Ignoring findings: <https://www.shellcheck.net/wiki/Ignore>

### Files

- `.shellcheckrc`: configuration

### Interactions

- [`pre-commit`](#pre-commit) includes hooks to ensure files comply with ShellCheck

## Z Shell

<https://www.zsh.org/>

### Documentation

- Cheat Sheet: <https://gist.github.com/ClementNerma/1dd94cb0f1884b9c20d1ba0037bdcde2>
- Manual: <https://zsh.sourceforge.io/Doc/Release/index.html>
