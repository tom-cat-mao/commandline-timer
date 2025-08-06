.PHONY: build clean install uninstall help

BINARY_NAME=timer
PREFIX=/usr/local/bin

# Build for current platform
build:
	go build -o $(BINARY_NAME) ./main.go

# Build for Linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux ./main.go

# Build for macOS
build-macos:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-macos ./main.go

# Build for macOS ARM64
build-macos-arm64:
	GOOS=darwin GOARCH=arm64 go build -o $(BINARY_NAME)-macos-arm64 ./main.go

# Build for all platforms
build-all: build-linux build-macos build-macos-arm64

# Install to system
install: build
	sudo cp $(BINARY_NAME) $(PREFIX)/
	sudo chmod 755 $(PREFIX)/$(BINARY_NAME)

# Uninstall from system
uninstall:
	sudo rm -f $(PREFIX)/$(BINARY_NAME)

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME) $(BINARY_NAME)-linux $(BINARY_NAME)-macos $(BINARY_NAME)-macos-arm64

# Download dependencies
deps:
	go mod download
	go mod tidy

# Run tests
test:
	go test -v ./...

# Help
help:
	@echo "Available commands:"
	@echo "  build         Build for current platform"
	@echo "  build-linux   Build for Linux"
	@echo "  build-macos   Build for macOS Intel"
	@echo "  build-macos-arm64 Build for macOS ARM64"
	@echo "  build-all     Build for all platforms"
	@echo "  install       Install to /usr/local/bin"
	@echo "  uninstall     Remove from /usr/local/bin"
	@echo "  clean         Clean build artifacts"
	@echo "  deps          Download dependencies"
	@echo "  test          Run tests"
	@echo "  help          Show this help message"