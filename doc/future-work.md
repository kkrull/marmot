# Future work

Ideas to extend or improve Marmot:

- Auto-complete for zsh.
- Error handling: `set -euo pipefail` more consistently. See
  <https://www.mulle-kybernetik.com/modern-bash-scripting/state-euxo-pipefail.html>.
- `exec`: Consider adding an option for whether to exit on the first failure, or keep going.
  Or ask the user if/when the first failure happens, since you probably don't know in advance.
- `host`: Add `marmot host import <host: bitbucket.org|github.com>` to register remote
  repositories and `marmot host clone` to clone them.
- `repo`: Add `marmot repo move <host> repository...`.
