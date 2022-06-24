BUILD_DIR=bin
BINARY_NAME=shorty.out
MAIN_FILE=main.go
OUT_FILE=$(BUILD_DIR)/$(BINARY_NAME)

all: run

deps: 
	go get

build:
	go build -o $(OUT_FILE)

test:
	go test -v

run:
	go run $(MAIN_FILE)

install:
	go install

uninstall:
	rm $$(go env GOPATH)/bin/shorty

clean:
	go clean
	rm -rf $(BUILD_DIR)