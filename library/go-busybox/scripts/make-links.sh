#!/bin/gosh
# shellcheck shell=sh

BB_BIN="/bin/bb"
BIN_DIR="/bin"
CMD_FILE="/tmp/bb_cmds.txt"

# Ensure command file exists
if [ ! -f "${CMD_FILE}" ]; then
    echo "Error: Command file ${CMD_FILE} does not exist"
    exit 1
fi

# Create symlinks
while IFS= read -r CMD; do
    "${BB_BIN}" ln -s -f "${BB_BIN}" "${BIN_DIR}/${CMD}"
    echo "Created symlink: ${BIN_DIR}/${CMD} -> ${BB_BIN}"
done < "${CMD_FILE}"