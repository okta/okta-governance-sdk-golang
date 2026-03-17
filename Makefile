COLOR_OK=\\x1b[0;32m
COLOR_NONE=\x1b[0m
COLOR_ERROR=\x1b[31;01m
COLOR_WARNING=\x1b[33;05m
COLOR_OKTA=\x1B[34;01m
GOLANGCI_LINT := golangci-lint
GOLANGCI_LINT_VERSION := v2.5.0

help:
	@echo "$(COLOR_OK)Okta IG SDK for Golang$(COLOR_NONE)"
	@echo ""
	@echo "$(COLOR_WARNING)Usage:$(COLOR_NONE)"
	@echo "$(COLOR_OK)  make [command]$(COLOR_NONE)"
	@echo ""
	@echo "$(COLOR_WARNING)Available commands:$(COLOR_NONE)"
	@echo "$(COLOR_OK)  build                   Clean and build the Okta Golang SDK generated files$(COLOR_NONE)"
	@echo "$(COLOR_OK)  clean-files             Deletes all generated files$(COLOR_NONE)"
	@echo "$(COLOR_OK)  generate-files          Generates files based around spec$(COLOR_NONE)"
	@echo "$(COLOR_OK)  pull-spec               Pull down the most recent released version of the spec$(COLOR_NONE)"
	@echo "$(COLOR_WARNING)test$(COLOR_NONE)"
	@echo "$(COLOR_OK)  test:all                Run all tests$(COLOR_NONE)"
	@echo "$(COLOR_OK)  test:integration        Run only integration tests$(COLOR_NONE)"
	@echo "$(COLOR_OK)  test:unit               Run only unit tests$(COLOR_NONE)"

build:
	@echo "$(COLOR_OKTA)Building SDK...$(COLOR_NONE)"
	make clean-files
	make generate-files
	make test:all

clean-files:
	@echo "$(COLOR_OKTA)Cleaning Up Old Generated Files...$(COLOR_NONE)"
	cd openapi && yarn cleanFiles

generate-files:
	@echo "$(COLOR_OKTA)Generating SDK Files...$(COLOR_NONE)"
	cd openapi && yarn generator
	@echo "$(COLOR_OK)Running goimports and gofumpt on generated files...$(COLOR_NONE)"
	@make import
	@make fmt
	@echo "$(COLOR_OKTA)Done!$(COLOR_NONE)"

pull-spec:
	@echo "$(COLOR_OKTA)Pulling in latest spec...$(COLOR_NONE)"
	cd openapi && npm install
# Overwrite the spec.json file from npm if a git branch is specified by env variable.
ifneq ($(origin OPENAPI_SPEC_BRANCH),undefined)
	rm -f openapi/spec.json
	git clone --branch $(OPENAPI_SPEC_BRANCH) https://github.com/okta/okta-management-openapi-spec spec-raw
	cp spec-raw/dist/spec.json openapi/spec.json
	cp spec-raw/dist/spec.json openapi/node_modules/@okta/openapi/dist/spec.json
	rm -fr spec-raw
endif

.PHONY: fmt
fmt: check-golangci-lint # Format the code using `golangci-lint`
	@$(GOLANGCI_LINT) fmt

.PHONY: import
import: # Run goimports on all Go files
	@goimports -w .

check-golangci-lint:
	@which $(GOLANGCI_LINT) > /dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)

.PHONY: test
test:
	go test -failfast -race ./governance -test.v

# =============================================================================
# Test Targets
# =============================================================================
# test:vcr       - Run VCR tests
# test:contract  - Run contract tests
# test:e2e       - Run E2E tests (live API, requires credentials)
# test:record    - Record new VCR cassettes (requires credentials)
# =============================================================================

.PHONY: test\:vcr
test\:vcr:
	@echo "$(COLOR_OKTA)Running VCR tests (cassette playback)...$(COLOR_NONE)"
	go test -tags=vcr -v ./governance/.

.PHONY: test\:contract
test\:contract:
	@echo "$(COLOR_OKTA)Running contract tests (model validation)...$(COLOR_NONE)"
	go test -tags=contract -v ./governance/.

.PHONY: test\:e2e
test\:e2e:
	@echo "$(COLOR_OKTA)Running E2E tests (live API)...$(COLOR_NONE)"
	go test -tags=e2e -v ./governance/.

.PHONY: test\:record
test\:record:
	@echo "$(COLOR_OKTA)Recording cassettes for VCR tests...$(COLOR_NONE)"
	VCR_RECORD=true go test -tags=vcr -v ./governance/.

.PHONY: test\:all
test\:all:
	@echo "$(COLOR_OKTA)Running all tests...$(COLOR_NONE)"
	go test -failfast -race -v ./governance/...

generate:
	npx @openapitools/openapi-generator-cli generate -c ./.generator/config.yaml -i .generator/governance-production-combined-reference-minimal.yaml

