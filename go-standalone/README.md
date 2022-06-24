# Initial project setup

This was done once, and will not need to be done again - just a reminder

- create a module `go mod init shorty` inside the go project's root dir
- get the clipboard package as dependency for this project `go get golang.design/x/clipboard/cmd/gclip@latest`

# Continue developing

## run the tool

- `go run ./main.go`

## run tests

- `go test`

# build and install

- build a binary executable `go build -o bin/shorty`
- build and install the binary executable on your system `go install`
  - the location can be found with `go env GOPATH`
