.PHONY: watch
watch: ## Start a file watcher to run tests on change. (requires: watchexec)
	watchexec -c "go test -failfast ."

.PHONY: test
test: ## Runs the unit test suite
	go test -failfast ./...

## Help display.
## Pulls comments from beside commands and prints a nicely formatted
## display with the commands and their usage information.

.DEFAULT_GOAL := help

help: ## Prints this help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

