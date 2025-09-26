BUILD_DIR=bin
BINARY_NAME=shorty.out
MAIN_FILE=main.go
OUT_FILE=$(BUILD_DIR)/$(BINARY_NAME)

KARABINER_CONFIG_FILE=karabiner-shorty.json
KARABINER_CONFIG_PATH=$(HOME)/.config/karabiner/assets/complex_modifications/$(KARABINER_CONFIG_FILE)

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
	GOBIN=$(HOME)/bin go install
	cp $(KARABINER_CONFIG_FILE) $(KARABINER_CONFIG_PATH)

uninstall:
	rm $$(go env GOPATH)/bin/shorty
	rm $(KARABINER_CONFIG_PATH)

clean:
	go clean
	rm -rf $(BUILD_DIR)