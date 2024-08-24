#!/bin/sh

echo "Running golang-lint check code"

# store all changes that have been staged, excluding deleted files
changes=$(git diff --name-only --staged --diff-filter=d)

if [ ${#changes} -gt 0 ]; then
    # check the formatting of all the files staged so far, excluding deleted files
    $(go env GOPATH)/bin/golangci-lint run

    if [ $? -ne 0 ]; then
        echo "Error: your code has linter errors"
        exit 1
    fi
else
    echo "No change was detected for golangci-lint to run (deleted files are ignored)"
fi
