# Multi-platform binary build
# http://www.codershaven.com/multi-platform-makefile-for-go/
.PHONY: all test clean

BIN_FOLDER="bin"

build: clean ## Build binaries
	@(go build -o bin/main main.go)

all: test build ## Build and run tests

test: clean ## Run unit tests
	@(go test ./...)

clean: ## Remove previous build
	@(go clean)
	@(rm -rf $(BIN_FOLDER))

bump: ## Do a version bump
	@(npx standard-version --skip.changelog)

bump-check: ## Check which would be the next version
	@(npx standard-version --skip.changelog --dry-run)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
