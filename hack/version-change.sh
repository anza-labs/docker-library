#!/usr/bin/env bash

set -e
GITHUB_OUTPUT="${1:?GITHUB_OUTPUT \(arg 1\) is required}"

# Get the list of changed version.txt files
git diff --name-only HEAD~1 HEAD | grep -E '^library/.+/version.txt$' | while read -r file; do
    PROJECT=$(echo "$file" | cut -d'/' -f2)
    VERSION=$(cat "$file")
    echo "Found project=$PROJECT version=$VERSION"
    echo "project=$PROJECT" >> "$GITHUB_OUTPUT"
    echo "version=$VERSION" >> "$GITHUB_OUTPUT"
done
