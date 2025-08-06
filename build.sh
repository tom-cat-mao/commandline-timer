#!/bin/bash

# Build script for commandline-timer
# This script builds the timer for multiple platforms

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Building commandline-timer...${NC}"

# Get the current directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# Clean previous builds
echo -e "${YELLOW}Cleaning previous builds...${NC}"
make clean

# Download dependencies
echo -e "${YELLOW}Downloading dependencies...${NC}"
make deps

# Build for current platform
echo -e "${YELLOW}Building for current platform...${NC}"
make build

# Build for all platforms if requested
if [ "$1" = "--all" ]; then
    echo -e "${YELLOW}Building for all platforms...${NC}"
    make build-all
fi

echo -e "${GREEN}Build completed successfully!${NC}"
echo -e "${YELLOW}To install system-wide, run: sudo make install${NC}"