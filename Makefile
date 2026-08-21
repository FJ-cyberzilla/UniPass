.PHONY: all build check test vet fmt clean release

APP_NAME=unipass
BUILD_DIR=dist

all: check build

vet:
	@echo "🛡️ Inspecting code with go vet..."
	@go vet ./...
	@echo "  ✅ Go vet verification clean."

fmt:
	@echo "🎨 Formatting Go source code..."
	@go fmt ./...
	@echo "  ✅ Code formatting complete."

test:
	@echo "🧪 Executing unit test suite..."
	@go test -v ./...
	@echo "  ✅ All package unit tests passed!"

check: vet fmt test
	@echo "🔍 Running system health diagnosis..."
	@echo "  ✅ Static analysis passed clean!"

build:
	@echo "🔨 Compiling unipass local binary..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 go build -trimpath -o $(BUILD_DIR)/$(APP_NAME) ./cmd/unipass
	@echo "  ✅ Binary built successfully at $(BUILD_DIR)/$(APP_NAME)"

run-dev: build
	@echo "⚡ Launching UniPass CLI..."
	@./$(BUILD_DIR)/$(APP_NAME)

release:
	@./build.sh

clean:
	@rm -rf $(BUILD_DIR)
	@echo "🧹 Cleaned build artifacts."
