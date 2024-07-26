# Scripts

.PHONY: default
default: all

## Environment

### Paths

prefix ?= /usr/local
exec_prefix ?= $(prefix)
bindir := $(exec_prefix)/bin

datarootdir := $(prefix)/share
mandir := $(datarootdir)/man
man1dir := $(mandir)/man1

.PHONY: debug-path
debug-path:
	$(info Installation paths:)
	$(info - prefix: $(prefix))
	$(info - exec_prefix: $(exec_prefix))
	$(info - bindir: $(bindir))

	$(info Data paths:)
	$(info - datarootdir: $(datarootdir))
	$(info - mandir: $(mandir))
	$(info - man1dir: $(man1dir))

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
debug: debug-path debug-sources #> Show debugging information
	@:

.PHONY: help
help: #> Show this help
	@sed -n \
		-e '/@sed/!s/#[.] *//p' \
		-e '/@sed/!s/:.*#> /:/p' \
		$(MAKEFILE_LIST) \
		| column -ts :
