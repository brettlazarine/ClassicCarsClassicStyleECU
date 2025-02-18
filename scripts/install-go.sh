#!/bin/bash

# Exit script on error
set -e

# Set non-interactive mode to prevent prompts
export DEBIAN_FRONTEND=noninteractive

# Update and install dependencies
echo "Updating and installing dependencies..."
sudo apt update && sudo apt full-upgrade -y --allow-downgrades --allow-remove-essential --allow-change-held-packages
sudo apt install -y git curl wget build-essential

# Determine system architecture
ARCH=$(dpkg --print-architecture)
if [[ "$ARCH" == "arm64" ]]; then
    GO_ARCH="linux-arm64"
elif [[ "$ARCH" == "armhf" || "$ARCH" == "armv7l" ]]; then
    GO_ARCH="linux-armv6l"
else
    GO_ARCH="linux-amd64"
fi

# Go version
GO_VERSION="1.24.0"

# Check if Go is already installed
if command -v go &>/dev/null; then
    INSTALLED_VERSION=$(go version | awk '{print $3}')
    echo "Go is already installed ($INSTALLED_VERSION). Removing old version..."
    sudo rm -rf /usr/local/go
fi

# Download and install Go
echo "Downloading and installing Go $GO_VERSION for $GO_ARCH..."
cd /usr/local
sudo wget "https://go.dev/dl/go${GO_VERSION}.${GO_ARCH}.tar.gz"
sudo tar -C /usr/local -xzf "go${GO_VERSION}.${GO_ARCH}.tar.gz"
sudo rm "go${GO_VERSION}.${GO_ARCH}.tar.gz"

# Add Go to PATH system-wide
echo "Configuring system-wide Go PATH..."
echo 'export PATH=$PATH:/usr/local/go/bin' | sudo tee /etc/profile.d/go.sh > /dev/null
sudo chmod +x /etc/profile.d/go.sh

# Apply changes immediately
export PATH=$PATH:/usr/local/go/bin

# Verify installation
echo "Verifying Go installation..."
go version && echo "Go $GO_VERSION installed successfully!"
