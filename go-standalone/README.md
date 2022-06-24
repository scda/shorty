# Initial project setup

Just a reminder - this was done once, and you do not need to run this to make use of this project's code

- create a module `go mod init shorty` inside the go project's root dir
- get the clipboard package as dependency for this project `go get golang.design/x/clipboard/cmd/gclip@latest`

# Continue developing

## run the tool

look into the [makefile](./makefile) for more options. For example:

- run the tool by executing `make` or `make run`
- run tests by executing `make test`
