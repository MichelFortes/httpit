.PHONY: all build test run docker clean setup

APP_NAME := httpit
BUILD_DIR := bin
CMD_PATH := ./cmd/httpit/main.go
EXAMPLE_CONFIG := examples/test-scheme-example.json

all: setup test build

setup:
	@echo "Fetching dependencies..."
	go mod download
	go mod tidy

build:
	@echo "Building $(APP_NAME)..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_PATH)

test:
	@echo "Running tests..."
	go test ./... -v

run:
	@echo "Running $(APP_NAME)..."
	go run $(CMD_PATH) $(EXAMPLE_CONFIG)

docker:
	@echo "Building Docker image..."
	docker build -f build/Dockerfile -t $(APP_NAME) .

clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
