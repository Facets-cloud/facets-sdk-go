.PHONY: help download-spec clean-spec validate-spec generate-client build-example run-example test clean install-deps all

# Variables
SWAGGER_URL ?= https://facetsdemo.console.facets.cloud/v3/api-docs
SWAGGER_DIR := swagger
SWAGGER_RAW := $(SWAGGER_DIR)/swagger.json
SWAGGER_FIXED := $(SWAGGER_DIR)/swagger_fixed.json
SWAGGER_SCRIPT := $(SWAGGER_DIR)/fix_swagger.js
GENERATED_DIR := facets
EXAMPLE_DIR := examples
EXAMPLE_BINARY := $(EXAMPLE_DIR)/basic_client

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install-deps: ## Install required dependencies (swagger CLI)
	@echo "Checking dependencies..."
	@which swagger > /dev/null || (echo "Installing go-swagger..." && go install github.com/go-swagger/go-swagger/cmd/swagger@latest)
	@which node > /dev/null || (echo "Node.js is required for spec cleaning. Please install Node.js" && exit 1)
	@echo "Dependencies installed"

download-spec: ## Download OpenAPI spec from Facets API
	@echo "Downloading OpenAPI spec from $(SWAGGER_URL)..."
	@mkdir -p $(SWAGGER_DIR)
	@curl -s -o $(SWAGGER_RAW) $(SWAGGER_URL)
	@echo "OpenAPI spec downloaded to $(SWAGGER_RAW)"

clean-spec: download-spec ## Download and clean the OpenAPI spec
	@echo "Cleaning OpenAPI spec..."
	@node $(SWAGGER_SCRIPT) $(SWAGGER_RAW) $(SWAGGER_FIXED)
	@echo "Cleaned spec saved to $(SWAGGER_FIXED)"

validate-spec: clean-spec ## Validate the cleaned OpenAPI spec
	@echo "Validating Swagger spec..."
	@swagger validate $(SWAGGER_FIXED) 2>&1 | head -20 || true
	@echo ""

generate-client: clean-spec ## Generate Go client from OpenAPI spec
	@echo "Generating Go client from OpenAPI spec..."
	@swagger generate client \
		-f $(SWAGGER_FIXED) \
		-A facets \
		-t $(GENERATED_DIR) \
		--skip-validation \
		--skip-tag-packages 2>&1 | tee swagger_generate.log || (echo "Generation failed. Check swagger_generate.log for details" && exit 1)
	@echo "Client generated successfully in $(GENERATED_DIR)/"
	@echo "Running go mod tidy..."
	@go mod tidy
	@echo "Formatting generated code..."
	@go fmt ./facets/... || true

build-example: ## Build the example client
	@echo "Building example client..."
	@cd $(EXAMPLE_DIR) && go build -o basic_client basic_client.go
	@echo "Example client built: $(EXAMPLE_BINARY)"

run-example: ## Run the example client (requires env vars: FACETS_API_HOST, FACETS_USERNAME, FACETS_API_TOKEN)
	@echo "Running example client..."
	@if [ -z "$$FACETS_API_HOST" ] || [ -z "$$FACETS_USERNAME" ] || [ -z "$$FACETS_API_TOKEN" ]; then \
		echo "Error: Required environment variables not set"; \
		echo "Please set: FACETS_API_HOST, FACETS_USERNAME, FACETS_API_TOKEN"; \
		exit 1; \
	fi
	@cd $(EXAMPLE_DIR) && go run basic_client.go

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

integration-test: build-example ## Run integration test with example client
	@echo "Running integration test..."
	@if [ -z "$$FACETS_API_HOST" ] || [ -z "$$FACETS_USERNAME" ] || [ -z "$$FACETS_API_TOKEN" ]; then \
		echo "Error: Required environment variables not set"; \
		echo "Please set: FACETS_API_HOST, FACETS_USERNAME, FACETS_API_TOKEN"; \
		exit 1; \
	fi
	@echo "Testing SDK integration..."
	@cd $(EXAMPLE_DIR) && ./basic_client && echo "✓ Integration test passed"

fmt: ## Format Go code
	@echo "Formatting code..."
	@go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

lint: fmt vet ## Run linters

clean: ## Clean generated files and binaries
	@echo "Cleaning up..."
	@rm -f $(EXAMPLE_BINARY)
	@rm -f $(SWAGGER_RAW) $(SWAGGER_FIXED)
	@echo "Cleaned"

clean-all: clean ## Clean everything including generated client code
	@echo "Cleaning all generated files..."
	@rm -rf $(GENERATED_DIR)/client $(GENERATED_DIR)/models
	@echo "All clean"

all: install-deps generate-client build-example ## Full build: install deps, generate client, build example
	@echo "Build complete!"

.DEFAULT_GOAL := help
