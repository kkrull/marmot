# Scripts

.PHONY: default
default: all

## Environment

### Paths

prefix ?= /usr/local
exec_prefix ?= $(prefix)
bindir := $(exec_prefix)/bin

.PHONY: debug-paths
debug-paths:
	$(info Installation paths:)
	$(info - prefix: $(prefix))
	$(info - exec_prefix: $(exec_prefix))
	$(info - bindir: $(bindir))

### Programs

### Sources

srcdir := $(realpath .)

.PHONY: debug-sources
debug-sources:
	$(info Sources:)
	$(info - srcdir: $(srcdir))

#. STANDARD TARGETS

.PHONY: all clean install test uninstall
all:
	@:

clean:
	@:

install: #> Install symlink from $(bindir) to sources
	mkdir -p $(bindir)
	ln -f -s $(srcdir)/marmot.zsh $(bindir)/marmot

.PHONY: install-tools
install-tools:
	@:

test:
	@:

uninstall: #> Delete symlink from $(bindir)
	$(RM) $(bindir)/marmot

#. OTHER TARGETS

.PHONY: debug
.NOTPARALLEL: debug
debug: debug-paths debug-sources #> Show debugging information
	@:

# https://stackoverflow.com/a/47107132/112682
.PHONY: help
help: #> Show this help
	@sed -n \
		-e '/@sed/!s/#[.] */_margin_\n/p' \
		-e '/@sed/!s/:.*#> /:/p' \
		$(MAKEFILE_LIST) \
	| column -ts : | sed -e 's/_margin_//'
