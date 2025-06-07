BINARY_DIR=bin
APP_NAME=go-ludus

.PHONY: build-cli run-cli

build-cli:
	mkdir -p $(BINARY_DIR)
	go build -o $(BINARY_DIR)/$(APP_NAME) ./cmd/cli

run-cli:
	go run ./cmd/cli/main.go
