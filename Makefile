# ============================================================================ #
# HELPERS
# ============================================================================ #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ============================================================================ #
# EXAMPLES
# ============================================================================ #

## example/motion: example to change the camera motion status
.PHONY: example/motion
example/motion:
	go run ./example/motion/change.go

## example/snap: example to snap a picture from the camera
.PHONY: example/snap
example/snap:
	go run ./example/picture/snap.go

# ============================================================================ #
# QUALITY CONTROL
# ============================================================================ #

## lint: run Go linters (golangci-lint)
.PHONY: lint
lint:
	@echo 'Linting code...'
	golangci-lint run

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: vendor lint
	@echo 'Formatting code...'
	go fmt
	@echo 'Running tests...'
	go test -v -race -vet=off

## vulncheck: check for known vulnerabilities
.PHONY: vulncheck
vulncheck:
	@echo 'Checking for vulnerabilities...'
	govulncheck -show verbose ./...

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
coverage-html: test-coverage
	go tool cover -html=coverage.out
