# ============================================================================ #
# HELPERS
# ============================================================================ #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ============================================================================ #
# DEVELOPMENT
# ============================================================================ #

## run/example: run the example application
.PHONY: run/example
run/example:
	go run ./example/main.go

# ============================================================================ #
# QUALITY CONTROL
# ============================================================================ #

## lint: run Go linters (golangci-lint)
.PHONY: lint
lint: ## Runs lint
	@echo 'Linting code...'
	golangci-lint run

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: vendor lint
	@echo 'Formatting code...'
	go fmt
	@echo 'Running tests...'
	go test -v -race -vet=off

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

## test-coverage: test code and check coverage (generates coverage.out)
.PHONY: test-coverage
test-coverage:
	go test -v -race -coverprofile=coverage.out -covermode=atomic

## coverage-html: check code coverage and generate HTML report
.PHONY: coverage-html
coverage-html: coverage
	go tool cover -html=coverage.out
