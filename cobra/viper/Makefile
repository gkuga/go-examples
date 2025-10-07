# Makefile for Viper Examples
# This Makefile provides convenient commands for building and running the viper examples

.PHONY: help build run run-cobra clean test deps fmt vet check install

# Default target - show help
help:
	@echo "Available commands:"
	@echo "  make run          - Run the pflag version (main.go)"
	@echo "  make run-cobra    - Run the cobra version (cobra_main.go)"
	@echo "  make build        - Build both executables"
	@echo "  make deps         - Download dependencies"
	@echo "  make fmt          - Format Go code"
	@echo "  make vet          - Run go vet"
	@echo "  make check        - Run fmt and vet"
	@echo "  make test         - Run tests (if any)"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make install      - Install executables to GOPATH/bin"
	@echo "  make demo         - Run demonstration of both versions"

# Build targets
build: build-pflag build-cobra

build-pflag:
	@echo "Building pflag version..."
	go build -o bin/viper-pflag main.go

build-cobra:
	@echo "Building cobra version..."
	go build -o bin/viper-cobra cobra_main.go

# Run targets
run:
	@echo "Running pflag version..."
	go run main.go

run-cobra:
	@echo "Running cobra version..."
	go run cobra_main.go

# Run with sample configurations
run-with-env:
	@echo "Running pflag version with environment variables..."
	VIPEREXAMPLE_APP_NAME="Makefile Test" VIPEREXAMPLE_SERVER_PORT=7777 go run main.go --app.debug

run-cobra-with-env:
	@echo "Running cobra version with environment variables..."
	VIPEREXAMPLE_APP_NAME="Makefile Cobra Test" VIPEREXAMPLE_SERVER_PORT=8888 go run cobra_main.go --app.debug config show

# Development targets
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

fmt:
	@echo "Formatting Go code..."
	go fmt ./...

vet:
	@echo "Running go vet on individual files..."
	go vet main.go
	go vet cobra_main.go

check: fmt vet
	@echo "Code check completed"

test:
	@echo "Running tests..."
	go test ./... -v

# Clean targets
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -f output_config.yaml

# Install targets
install: build
	@echo "Installing executables..."
	mkdir -p $(GOPATH)/bin
	cp bin/viper-pflag $(GOPATH)/bin/
	cp bin/viper-cobra $(GOPATH)/bin/

# Demo targets
demo: demo-pflag demo-cobra

demo-pflag:
	@echo ""
	@echo "=== pflag Version Demo ==="
	@echo "1. Basic run:"
	@go run main.go
	@echo ""
	@echo "2. With flags:"
	@go run main.go --app.debug --server.port=5555 --logging.level=warn
	@echo ""
	@echo "3. With environment variables:"
	@VIPEREXAMPLE_APP_NAME="Demo App" go run main.go

demo-cobra:
	@echo ""
	@echo "=== Cobra Version Demo ==="
	@echo "1. Basic run:"
	@go run cobra_main.go
	@echo ""
	@echo "2. Show config subcommand:"
	@go run cobra_main.go config show
	@echo ""
	@echo "3. Validate config subcommand:"
	@go run cobra_main.go config validate
	@echo ""
	@echo "4. With flags and env vars:"
	@VIPEREXAMPLE_APP_NAME="Cobra Demo" go run cobra_main.go --server.port=6666 config show

# Development convenience targets
dev-pflag:
	@echo "Starting pflag development mode (auto-reload with entr if available)..."
	@if command -v entr > /dev/null 2>&1; then \
		find . -name "*.go" | entr -r make run; \
	else \
		echo "entr not found. Install with: brew install entr"; \
		echo "Running once:"; \
		make run; \
	fi

dev-cobra:
	@echo "Starting cobra development mode (auto-reload with entr if available)..."
	@if command -v entr > /dev/null 2>&1; then \
		find . -name "*.go" | entr -r make run-cobra; \
	else \
		echo "entr not found. Install with: brew install entr"; \
		echo "Running once:"; \
		make run-cobra; \
	fi

# Configuration targets
show-config:
	@echo "Current configuration (via pflag):"
	@go run main.go
	@echo ""
	@echo "Current configuration (via cobra):"
	@go run cobra_main.go config show

validate-config:
	@echo "Validating configuration..."
	@go run cobra_main.go config validate

# Help for specific targets
help-pflag:
	@echo "pflag version help:"
	@go run main.go --help

help-cobra:
	@echo "Cobra version help:"
	@go run cobra_main.go --help
	@echo ""
	@echo "Cobra config subcommand help:"
	@go run cobra_main.go config --help

# Project info
info:
	@echo "Project: Viper Configuration Examples"
	@echo "Go version: $(shell go version)"
	@echo "Module: $(shell head -1 go.mod)"
	@echo "Dependencies:"
	@go list -m all | grep -v "^viper-example$$"