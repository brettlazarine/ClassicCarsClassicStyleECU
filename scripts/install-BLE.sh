#!/bin/bash

# Exit script on error
set -e

# Update
echo "Updating and upgrading..."
sudo apt update && sudo apt upgrade -y

# Install BlueZ
echo "Installing BlueZ..."
sudo apt install -y bluetooth bluez libbluetooth-dev

# Enable and start Bluetooth service
echo "Enabling and starting Bluetooth service..."
sudo systemctl enable bluetooth
sudo systemctl start bluetooth