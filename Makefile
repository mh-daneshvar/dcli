# Go parameters
BINARY_NAME=dcli
CMD_PATH=cmd/cli/main.go
VERSION=1.0.0

# Define the OS and architecture combinations to support (for cross-compilation)
OS_LIST=windows linux darwin
ARCH_LIST=amd64

# The install path (default is /usr/local/bin)
INSTALL_PATH=/usr/local/bin

# Build the binary for the current platform
build:
	@echo "Building $(BINARY_NAME)..."
	GO111MODULE=on go build -o $(BINARY_NAME) $(CMD_PATH)
	@echo "Build complete."

# Build for each supported OS and architecture
build-all: clean
	@echo "Building for all supported platforms..."
	@for os in $(OS_LIST); do \
		for arch in $(ARCH_LIST); do \
			echo "Building for $$os/$$arch..."; \
			GOOS=$$os GOARCH=$$arch GO111MODULE=on go build -o $(BINARY_NAME)-$$os-$$arch $(CMD_PATH); \
		done \
	done
	@echo "All builds complete."

# Clean up any previous builds
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME) $(BINARY_NAME)-*
	@echo "Cleanup complete."

# Install the binary on the system (default install path: /usr/local/bin)
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	install $(BINARY_NAME) $(INSTALL_PATH)
	@echo "$(BINARY_NAME) installed successfully."

# Uninstall the binary from the system
uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(INSTALL_PATH)..."
	rm -f $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "$(BINARY_NAME) uninstalled successfully."

# Cross-compile for a specific OS and architecture
build-cross:
	@echo "Building for OS=$(OS) ARCH=$(ARCH)..."
	GOOS=$(OS) GOARCH=$(ARCH) GO111MODULE=on go build -o $(BINARY_NAME)-$(OS)-$(ARCH) $(CMD_PATH)
	@echo "Cross-build complete for $(OS)/$(ARCH)."

# Run all tests
test:
	@echo "Running all tests..."
	go test -v ./...
	@echo "Tests completed."

.PHONY: build clean install uninstall build-all build-cross test
