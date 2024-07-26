# Tools

Tools used by some part of this project's build, deployment, or development processes.

## Code Spell Checker

<https://cspell.org/>

### Files

- `cspell.config.yaml`: Configuration and dictionary

### Documentation

- Configuration file:
  - `enableFiletypes`:
    <https://streetsidesoftware.com/vscode-spell-checker/docs/configuration/#cspellenabledfiletypes>

### Interactions

- VS Code extension:
  <https://marketplace.visualstudio.com/items?itemName=streetsidesoftware.code-spell-checker>

---

## EditorConfig

<https://editorconfig.org/>

### Files

- `.editorconfig`: Configuration

### Interactions

- [`pre-commit`](#pre-commit) includes hooks to ensure files comply with EditorConfig
- VS Code extension: <https://marketplace.visualstudio.com/items?itemName=EditorConfig.EditorConfig>

---

## `fswatch`

<https://github.com/emcrisostomo/fswatch>

### Interactions

- [GNU Make](#gnu-make) includes tasks that use `fswatch`

---

## GitHub Actions

<https://docs.github.com/en/actions>

### Documentation

- [CI/CD Jobs](./cicd-jobs.md)
- Workflow Syntax:
  <https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions>

### Files

- `.github/workflows/`: Job definitions

### Interactions

- [`pre-commit`](#pre-commit) checks are also run during some jobs
- VS Code extension:
  <https://marketplace.visualstudio.com/items?itemName=GitHub.vscode-github-actions>

---

## GNU Make

<https://www.gnu.org/software/make/>

### Documentation

- GNU Make Manual: <https://www.gnu.org/software/make/manual/>
- Standard Targets:
  <https://www.gnu.org/software/make/manual/html_node/Standard-Targets.html#Standard-Targets>

### Files

- `Makefile`: targets to build everything and tasks to help set up your environment
- `man/Makefile`: targets to build manuals
- `src/go/Makefile`: targets to build the Go version of marmot
- `src/zsh/Makefile`: targets to build programs with scripts

### Interactions

- VS Code extension: <https://marketplace.visualstudio.com/items?itemName=ms-vscode.makefile-tools>

---

## Go

<https://go.dev/>

### Documentation

- A Tour of Go: <https://go.dev/tour/>
- Effective Go: <https://go.dev/doc/effective_go>
- Package and documentation index: <https://pkg.go.dev/>

### Interactions

- [GNU Make](#gnu-make) builds binaries from Go sources and automates other common development tasks
- VS Code extension: <https://marketplace.visualstudio.com/items?itemName=golang.Go>

---

## Homebrew

<https://brew.sh/>

### Documentation

- `homebrew-bundle`: <https://github.com/Homebrew/homebrew-bundle>

### Files

In `etc/macos/`:

- `Brewfile.developer` and `Brewfile.developer.lock.json`: packages for developers, to work on this
  project
- `Brewfile.user` and `Brewfile.user.lock.json`: packages for end users, to run the programs built
  from these sources

---

## `jo`

<https://jpmens.net/2016/03/05/a-shell-command-to-create-json-jo/>

### Documentation

- GitHub: <https://github.com/jpmens/jo>
- Installation: <https://github.com/jpmens/jo?tab=readme-ov-file#install>

---

## `jq`

<https://jqlang.github.io/jq/>

### Documentation

- Manual: <https://jqlang.github.io/jq/manual/v1.7/>

---

## Markdown

<https://www.markdownguide.org/>

### Documentation

- Syntax: <https://www.markdownguide.org/basic-syntax/>

### Interactions

- [Mermaid](#mermaid) is embedded in some Markdown documents
- VS Code extensions
  - Markdown All in One:
    <https://marketplace.visualstudio.com/items?itemName=yzhang.markdown-all-in-one>

---

## Markdownlint

<https://github.com/DavidAnson/markdownlint-cli2>

### Documentation

- Configuration JSON schema:
  <https://github.com/DavidAnson/markdownlint-cli2#markdownlint-cli2jsonc>
- Main: <https://github.com/DavidAnson/markdownlint>
- Rules: <https://github.com/DavidAnson/markdownlint/blob/main/doc/Rules.md>
- `pre-commit` hook: <https://github.com/DavidAnson/markdownlint-cli2?tab=readme-ov-file#pre-commit>

### Files

- `.markdownlint-cli2.jsonc`: configuration

### Interactions

- [`pre-commit`](#pre-commit) includes hooks to ensure files comply with linting rules
- VS Code extension:
  <https://marketplace.visualstudio.com/items?itemName=DavidAnson.vscode-markdownlint>

---

## Mermaid

<https://mermaid-js.github.io/mermaid>

### Documentation

- Customized styling directives: <https://stackoverflow.com/q/71864287/112682>
- Syntax: <https://mermaid.js.org/intro/n00b-syntaxReference.html>

### Interactions

- VS Code extensions
  - Markdown Preview Mermaid Support:
    <https://marketplace.visualstudio.com/items?itemName=bierner.markdown-mermaid>
  - Mermaid Editor:
    <https://marketplace.visualstudio.com/items?itemName=tomoyukim.vscode-mermaid-editor>

## Pandoc

<https://pandoc.org/>

### Documentation

- GitHub Actions for Pandoc: <https://pandoc.org/installing.html#github-actions>
- Installation: <https://pandoc.org/installing.html>

### Interactions

- [GNU Make](#gnu-make) runs Pandoc to convert command documentation to manuals

---

## `pre-commit`

<https://pre-commit.com/>

### Documentation

- Configuration syntax: <https://pre-commit.com/#pre-commit-configyaml---top-level>
- Installation: <https://pre-commit.com/#install>
- Supported Hooks: <https://pre-commit.com/hooks.html>

### Files

- `.pre-commit.yaml`: configures repository sources and hooks to run.  May also contain
  configuration not already included in any hook-specific configuration files (e.g.
  `.editorconfig`).

### Interactions

- [EditorConfig](#editorconfig) is checked by `pre-commit`
- [GNU Make](#gnu-make) contains [tasks](./task-automation.md#pre-commit-targets) to install Git
  hooks and update `pre-commit` repositories
- [`markdownlint-cli2`](#markdownlint) is run by `pre-commit`
- [ShellCheck](#shellcheck) is run by `pre-commit`

---

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

---

## Z Shell

<https://www.zsh.org/>

### Documentation

- Cheat Sheet: <https://gist.github.com/ClementNerma/1dd94cb0f1884b9c20d1ba0037bdcde2>
- Manual: <https://zsh.sourceforge.io/Doc/Release/index.html>
