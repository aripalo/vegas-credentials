# Multi-platform binary build
# http://www.codershaven.com/multi-platform-makefile-for-go/
.PHONY: all test clean

EXECUTABLE=aws-mfa-assume-credential-process
BIN_FOLDER="bin"
WINDOWS=$(BIN_FOLDER)/$(EXECUTABLE)-windows-amd64.exe
LINUX=$(BIN_FOLDER)/$(EXECUTABLE)-linux-amd64
DARWIN=$(BIN_FOLDER)/$(EXECUTABLE)-darwin-amd64
VERSION=$(shell git describe --tags --always --long --dirty)

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -i -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -i -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -i -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/main.go

build: windows linux darwin ## Build binaries
	@echo version: $(VERSION)

all: test build ## Build and run tests

test: clean ## Run unit tests
	@(go test ./...)

clean: ## Remove previous build
	@(go clean)
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
