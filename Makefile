# Include variables from the .envrc file
include .envrc

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

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt
	@echo 'Linting code...'
	go vet
	@echo 'Running tests...'
	go test -race -vet=off -v

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

## coverage: check code coverage (generates coverage.out)
.PHONY: coverage
coverage:
	go test -coverprofile=coverage.out -covermode=set

## coverage-html: check code coverage and generate HTML report
.PHONY: coverage-html
coverage-html: coverage
	go tool cover -html=coverage.out
