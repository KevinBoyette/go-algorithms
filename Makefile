.PHONY: help

help: ## Prints out this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Run Tests with a coverage report
	go test ./... -cover

test-verbose: ## Run tests, get coverage, and display each test run
	go test -v -cover ./...
