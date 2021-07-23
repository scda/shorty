# shorty

shorty can reduce amazon `/dp` product links and therefore exclude search-queries or other references from the URL

## build

- run and debug in VS Code with `F5` (will open a new test-window)
- refresh test-window with `Developer: Reload Window` command in the test-window itself (or restart the debug process completely - both quite fast)

## publish

This extension will not be published at this point. Create and install the package locally.

- package with `vsce package`
- install the extension
  - copy the `.vsix` file into the extensions directory `~/.vscode/extensions`
  - OR install via CLI `code --install-extension <FILE_NAME>.vsix`

## commands

Currently the only available command is `Cleanup and shorten amazon URLs`

- this will affect only selected lines
- lines not matching the pattern will be left untouched
