# Future work

Ideas to extend or improve Marmot:

- Auto-complete for zsh.
- `exec`: Consider adding an option for whether to exit on the first failure, or keep going.
  Or ask the user if/when the first failure happens, since you probably don't know in advance.
- `host`: Add `marmot host import <host: bitbucket.org|github.com>` to register remote
  repositories and `marmot host clone` to clone them.
- `repo`: Add `marmot repo move <host> repository...`.

## What next

- Make a `marmot` command that clones the repositories (to `~/git/user.host/`), registers the
  repository with marmot (avoiding duplicates), and puts them in a special `@inbox/user.host`
  category.
- Update `marmot category add` to remove a repository from its `@inbox`, once it has been
  categorized at least once.
