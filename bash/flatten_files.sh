#!/bin/bash

# Usage: ./flatten_files.sh <source_directory> <target_directory>
SOURCE_DIR="$1"
TARGET_DIR="$2"

mkdir -p "$TARGET_DIR"
find "$SOURCE_DIR" -type f -exec cp {} "$TARGET_DIR" \;

echo "All files have been extracted to $TARGET_DIR without preserving folder structure."
