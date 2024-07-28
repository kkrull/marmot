# Manuals (e.g. "man pages")
# Guide: https://eddieantonio.ca/blog/2015/12/18/authoring-manpages-in-markdown-with-pandoc/
# man-pages conventions: https://linux.die.net/man/7/man-pages

.PHONY: default
default: all

## Environment

### Paths

prefix ?= /usr/local

datarootdir := $(prefix)/share
mandir := $(datarootdir)/man
man1dir := $(mandir)/man1
man7dir := $(mandir)/man7

.PHONY: path-info
path-info:
	$(info Paths:)
	$(info - datarootdir: $(datarootdir))
	$(info - mandir: $(mandir))
	$(info - man1dir: $(man1dir))
	$(info - man7dir: $(man7dir))
	$(info - prefix: $(prefix))
	@:

### Programs

PANDOC ?= pandoc
PANDOCFLAGS := -f markdown+definition_lists+line_blocks

### Sources

sources := $(wildcard pandoc/**.md)
sources_man1 := $(filter %.1.md,$(sources))
sources_man7 := $(filter %.7.md,$(sources))

.PHONY: source-info
source-info:
	$(info Sources:)
	$(info - sources: $(sources))
	$(info - sources_man1: $(sources_man1))
	$(info - sources_man7: $(sources_man7))
	@:

#. STANDARD TARGETS

.PHONY: all
all: groff-manual markdown-manual #> Build all manuals
	@:

.PHONY: clean
clean: groff-manual-clean markdown-manual-clean #> Remove built manuals
	@:

.PHONY: install
install: groff-manual-install #> Install man pages to $(mandir)
	@:

.PHONY: test
test:
	@:

.PHONY: uninstall
uninstall: groff-manual-uninstall #> Uninstall man pages
	@:

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
info: path-info source-info groff-manual-info markdown-manual-info #> Show build information
	@:

.PHONY: install-tools
install-tools:
	@:

#. GROFF MANUAL TARGETS

FSWATCH ?= fswatch
MANDOC ?= mandoc

groff_manual_installed := $(wildcard $(man1dir)/marmot* $(man7dir)/marmot*)
groff_man1_objects := $(patsubst pandoc/%.md,groff/%,$(sources_man1))
groff_man7_objects := $(patsubst pandoc/%.md,groff/%,$(sources_man7))

.PHONY: groff-manual
groff-manual: $(groff_man1_objects) $(groff_man7_objects) #> Build man pages

.PHONY: groff-manual-clean
groff-manual-clean: #> Remove built man pages
	$(RM) groff/**

.PHONY: groff-manual-info
groff-manual-info: #> Show build information for groff
	$(info Groff Manual:)
	$(info - groff_man1_objects: $(groff_man1_objects))
	$(info - groff_man7_objects: $(groff_man7_objects))
	$(info - groff_manual_installed: $(groff_manual_installed))
	@:

.PHONY: groff-manual-install
groff-manual-install: $(groff_man1_objects) $(groff_man7_objects) #> Install man pages
	mkdir -p $(man1dir) $(man7dir)
	install -g 0 -o 0 -m 0644 $(groff_man1_objects) $(man1dir)
	install -g 0 -o 0 -m 0644 $(groff_man7_objects) $(man7dir)

.PHONY: groff-manual-preview
groff-manual-preview: #> Render man pages without building or installing
	$(PANDOC) $(sources) $(PANDOCFLAGS) -s -t man \
		| $(MANDOC)

.PHONY: groff-manual-uninstall
groff-manual-uninstall: #> Uninstall man pages
	$(RM) $(groff_manual_installed)

.PHONY: groff-manual-watch
groff-manual-watch: #> Emit rendered man pages when Pandoc sources change
	$(FSWATCH) $(sources) \
		| xargs -I {} echo "$(PANDOC) {} $(PANDOCFLAGS) -s -t man | $(MANDOC)" \
		| sh

groff/%: pandoc/%.md
	$(PANDOC) $< $(PANDOCFLAGS) -o $@ -s -t man

#. MARKDOWN MANUAL TARGETS

markdown_manual_objects := $(patsubst pandoc/%.md,markdown/%.md,$(sources))
markdown_manual_objects_dir := markdown

.PHONY: markdown-manual
markdown-manual: $(markdown_manual_objects) #> Build Markdown manuals which are part of this repo

.PHONY: markdown-manual-clean
markdown-manual-clean: #> Remove local markdown manuals
	$(RM) markdown/**.md

.PHONY: markdown-manual-info
markdown-manual-info: #> Show build information for the markdown manual
	$(info Markdown Manual:)
	$(info - markdown_manual_objects: $(markdown_manual_objects))
	$(info - markdown_manual_objects_dir: $(markdown_manual_objects_dir))
	@:

# Make sure directory exists before building targets
# https://www.gnu.org/savannah-checkouts/gnu/make/manual/html_node/Prerequisite-Types.html
$(markdown_manual_objects): | $(markdown_manual_objects_dir)

$(markdown_manual_objects_dir):
	mkdir -p markdown

markdown/%.md: pandoc/%.md
	$(PANDOC) $< $(PANDOCFLAGS) -o $@ -s -t gfm
