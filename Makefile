# Project

.PHONY: default
default: all

## Environment

### Paths

### Programs

### Sources

#. STANDARD TARGETS

.PHONY: all
all: #> Build all sources
	$(MAKE) -C man all
	$(MAKE) -C src/go all
	$(MAKE) -C src/zsh all

.PHONY: clean
clean: pre-commit-gc #> Remove local build files
	$(MAKE) -C man clean
	$(MAKE) -C src/go clean
	$(MAKE) -C src/zsh clean

.PHONY: install
install: #> Install programs and manuals
	$(MAKE) -C man install
	$(MAKE) -C src/go install
	$(MAKE) -C src/zsh install

.PHONY: test
test: pre-commit-run #> Run tests and checks
	$(MAKE) -C man test
	$(MAKE) -C src/go test
	$(MAKE) -C src/zsh test

.PHONY: uninstall
uninstall: #> Uninstall programs and manuals
	$(MAKE) -C man uninstall
	$(MAKE) -C src/go uninstall
	$(MAKE) -C src/zsh uninstall

#. OTHER TARGETS

# https://stackoverflow.com/a/47107132/112682
.PHONY: help
help: #> Show this help
	@sed -n \
		-e '/@sed/!s/#[.] */_margin_\n/p' \
		-e '/@sed/!s/:.*#> /:/p' \
		$(MAKEFILE_LIST) \
	| column -ts : | sed -e 's/_margin_//'

.PHONY: help-all
help-all: help #> Show help for all Makefiles
	$(MAKE) -C man help
	$(MAKE) -C src/go help
	$(MAKE) -C src/zsh help

.PHONY: info
.NOTPARALLEL: info
info: #> Show build information
	$(MAKE) -C man info
	$(MAKE) -C src/go info
	$(MAKE) -C src/zsh info

install-tools: pre-commit-install #> Install development tools
	$(MAKE) -C man install-tools
	$(MAKE) -C src/go install-tools
	$(MAKE) -C src/zsh install-tools

#. HOMEBREW TARGETS

BREW ?= brew

.NOTPARALLEL: brew-install
.PHONY: brew-install
brew-install: brew-install-dev brew-install-user pre-commit-install #> Install dependencies with Homebrew
	@:

.PHONY: brew-install-dev
brew-install-dev: #> Install development tools with HomeBrew
	$(BREW) bundle install --file=./etc/macos/Brewfile.developer

.PHONY: brew-install-user
brew-install-user: #> Install end-user dependencies with HomeBrew
	$(BREW) bundle install --file=./etc/macos/Brewfile.user

#. PRE-COMMIT TARGETS

PRECOMMIT ?= pre-commit

.PHONY: pre-commit-gc
pre-commit-gc: #> Remove stale pre-commit files
	$(PRECOMMIT) gc

.PHONY: pre-commit-install
pre-commit-install: #> Install Git pre-commit hook
	$(PRECOMMIT) install

.PHONY: pre-commit-run
pre-commit-run: #> Run pre-commit on all sources
	$(PRECOMMIT) run --all-files

.PHONY: pre-commit-update
pre-commit-update: #> Update pre-commit plugins
	$(PRECOMMIT) autoupdate
