# Installation

Marmot was developed with Linux, MacOS, and Windows Subsystem for Linux (WSL) in mind.

## Clone and install

Clone this repository to a location of your choice.  Marmot doesn't have its own package to install;
just a single command that needs to be on your path somewhere.  You might try something simple like
this:

```sh
# Might require sudo
make install

# Might be unnecessary, if this is already on your path
path+=(/usr/local/bin)
```

If you need to install to somewhere other than `/usr/local/`, run `make` with another `prefix`:

```sh
prefix=/path/to/programs make install
```

## Install dependencies

Marmot uses a few packages that are listed in `etc/macos/Brewfile*`.  If you happen to be using
Homebrew, try this:

```sh
# Installs programs needed at runtime
brew bundle install --file=./etc/macos/Brewfile.user
```

If you use another package manager such as `apt` (Debian, Ubuntu), there should be similarly named
packages that provide the same commands.  It doesn't matter where they come from, as long as they
are reasonably up to date and on your path.

Please also remember to install `zsh` if you do not already have it.  You don't have to use `zsh` as
your main shell; it's just what Marmot uses internally.

## Use it

If you can run `marmot --help`, you have a working installation.  If you can re-start your terminal
and it _still_ works, you're in even better shape.

Now head over to [Getting Started](./getting-started.md) to see if Marmot does anything that might
be useful to you.
