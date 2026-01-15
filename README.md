# shorty

shorty can reduce URLs in the clipboard and therefore exclude search-queries or other references from the URL (like amazon or etsy product links)

Currently in the form of an installable go executable for your terminal that will directly manipulate the current clipboard contents.

## Initial project setup

Just a reminder - this was done once, and you do not need to run this to make use of this project's code

- create a module `go mod init shorty` inside the go project's root dir
- get the clipboard package as dependency for this project `go get golang.design/x/clipboard/cmd/gclip@latest`

## Continue developing

### run the tool

look into the [makefile](./makefile) for more options. For example:

- run the tool by executing `make` or `make run`
- run tests by executing `make test`

### installation

- to be able to call the tool from your command line, run `make install`
- to be able to call the tool with a a global shortcut, enable the script that was distributed via `make install` in Karabiner Elements.
