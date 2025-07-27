#!/bin/bash

LOG_DIR="/var/log/package_tracking"
DATE=$(date +%Y%m%d)

# Ensure directory exists
mkdir -p "$LOG_DIR"

# Current package list
dpkg -l > "$LOG_DIR/packages_current.list"

# Create diff if previous exists
if [ -f "$LOG_DIR/packages_previous.list" ]; then
    diff "$LOG_DIR/packages_previous.list" "$LOG_DIR/packages_current.list" > "$LOG_DIR/changes_current.diff"
fi

# Keep copy for next comparison
cp "$LOG_DIR/packages_current.list" "$LOG_DIR/packages_previous.list"
