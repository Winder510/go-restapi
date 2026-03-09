#!/bin/bash

# Build script for Go MySQL REST API

echo "Building Go MySQL REST API..."

# Set build variables
APP_NAME="go_mysql_api"
BUILD_DIR="./bin"
MAIN_FILE="./cmd/app/main.go"

# Create build directory if it doesn't exist
mkdir -p $BUILD_DIR

# Build for current platform
echo "Building for current platform..."
go build -o $BUILD_DIR/$APP_NAME $MAIN_FILE

# Build for multiple platforms (optional)
# echo "Building for Linux..."
# GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}-linux-amd64 $MAIN_FILE

# echo "Building for Windows..."
# GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}-windows-amd64.exe $MAIN_FILE

# echo "Building for macOS..."
# GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/${APP_NAME}-darwin-amd64 $MAIN_FILE

echo "Build complete! Executable: $BUILD_DIR/$APP_NAME"
