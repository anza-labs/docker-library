#!/bin/sh

BB_BIN="${1}"
CMD_FILE="/tmp/bb_cmds.txt"

# Get the list of commands from /bin/bb output
"${BB_BIN}" 2>&1 | sed -n '/Supported commands are:/,/error:/p' | sed -E '1d;$d' | sed 's/ - //' > "${CMD_FILE}"

# Ensure we got some commands
if [ ! -s "${CMD_FILE}" ]; then
    echo "Error: Failed to extract commands from ${BB_BIN}"
    exit 1
fi
