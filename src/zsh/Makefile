# Scripts

.PHONY: default
default: all

## Environment

### Paths

prefix ?= /usr/local
exec_prefix ?= $(prefix)
bindir := $(exec_prefix)/bin

.PHONY: path-info
path-info:
	$(info Paths:)
	$(info - bindir: $(bindir))
	$(info - exec_prefix: $(exec_prefix))
	$(info - prefix: $(prefix))

### Programs

### Sources

srcdir := $(realpath .)

.PHONY: source-info
source-info:
	$(info Sources:)
	$(info - srcdir: $(srcdir))

#. STANDARD TARGETS

.PHONY: all
all:
	@:

.PHONY: clean
clean:
	@:

.PHONY: install
install: #> Install symlink to sources
	mkdir -p $(bindir)
	ln -f -s $(srcdir)/marmot.zsh $(bindir)/marmot

.PHONY: test
test:
	@:

.PHONY: uninstall
uninstall: #> Delete symlink
	$(RM) $(bindir)/marmot

#. OTHER TARGETS

# https://stackoverflow.com/a/47107132/112682
.PHONY: help
help: #> Show this help
	@sed -n \
		-e '/@sed/!s/#[.] */_margin_\n/p' \
		-e '/@sed/!s/:.*#> /:/p' \
		$(MAKEFILE_LIST) \
	| column -ts : | sed -e 's/_margin_//'

.PHONY: info
.NOTPARALLEL: info
info: path-info source-info #> Show build information
	@:

.PHONY: install-tools
install-tools:
	@:
