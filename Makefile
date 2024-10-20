# Makefile for Go project 'photon'

APP_NAME=photon
APP_VERSION=0.1.0
SRC_DIR=./src
INSTALL_DIR=/usr/local/bin

.PHONY: all build run clean install uninstall

# Default target: build
all: build

# Build target with permission fix
build:
	go build -o $(APP_NAME) $(SRC_DIR)/main.go
	chmod +x ./$(APP_NAME)  # Add executable permission

# Install target: Create a symlink to the binary in /usr/local/bin
install: build
	@echo "Installing ./$(APP_NAME) to $(INSTALL_DIR)"
	sudo ln -sf $(shell pwd)/$(APP_NAME) $(INSTALL_DIR)/$(APP_NAME)

# Uninstall target: Remove the symlink
uninstall:
	@echo "Uninstalling $(APP_NAME) from $(INSTALL_DIR)"
	sudo rm -f $(INSTALL_DIR)/$(APP_NAME)

# Clean target: remove the binary
clean:
	rm -f $(APP_NAME)
