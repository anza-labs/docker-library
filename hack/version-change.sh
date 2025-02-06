#!/usr/bin/env bash

GITHUB_ENV="${1:?GITHUB_ENV \(arg 1\) is required}"

# Get the list of changed version.txt files
git diff --name-only HEAD~1 HEAD | grep -E '^library/.+/version.txt$' | while read -r file; do
    PROJECT=$(echo "$file" | cut -d'/' -f2)
    VERSION=$(cat "$file")
    echo "project=$PROJECT" >> "$GITHUB_ENV"
    echo "version=$VERSION" >> "$GITHUB_ENV"
done
