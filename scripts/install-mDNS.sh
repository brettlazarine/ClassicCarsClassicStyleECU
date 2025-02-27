#!/bin/bash

# Exit script on error
set -e

# Update
echo "Updating..."
sudo apt update

# Install Avahi
echo "Installing Avahi..."
sudo apt install avahi-daemon

# Start and Enable Avahi service
echo "Enabling Avahi..."
sudo systemctl enable avahi-daemon
echo "Starting Avahi..."
sudo systemctl start avahi-daemon