#!/usr/bin/env bash

set -e

SCRIPT="${1:?SCRIPT \(arg 1\) is required}"
REPOSITORY="${2:?REPOSITORY \(arg 2\) is required}"
CONTAINER_TOOL="${3:?CONTAINER_TOOL \(arg 3\) is required}"
BUILD_COMMAND="${4:?BUILD_COMMAND \(arg 4\) is required}"
MANIFEST_ARGS="${5}"
PROJECT="${6:?PROJECT \(arg 6\) is required}"
VERSION="${7:?VERSION \(arg 7\) is required}"
PLATFORM="${8:?PLATFORM \(arg 8\) is required}"

cat > "${SCRIPT}" <<EOF
#!/usr/bin/env bash
set -eux
set -o pipefail
export REPOSITORY='${REPOSITORY}'
export CONTAINER_TOOL='${CONTAINER_TOOL}'
export BUILD_COMMAND='${BUILD_COMMAND}'
export MANIFEST_ARGS='${MANIFEST_ARGS}'

EOF

is_semver() {
    local version="$1"
    # Remove the 'v' prefix if present
    version=${version#v}

    if [[ "$version" =~ ^([0-9]+)\.([0-9]+)\.([0-9]+)(-[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*)?$ ]]; then
        # Check if it is a pre-release or release version
        if [[ "$version" =~ - ]]; then
            echo "$1"
        else
            echo "$1,latest"
        fi
    else
        echo "$version"
    fi
}

cat >> "${SCRIPT}" <<EOF
    echo "Building project=${PROJECT} version=${VERSION} platform=${PLATFORM}"
    make build-${PROJECT} push-${PROJECT} PLATFORM='${PLATFORM}' TAG='$(is_semver ${VERSION})' BUILD_ARGS='VERSION=${VERSION}'
    make manifest-${PROJECT} PLATFORM='${PLATFORM}' TAG='$(is_semver ${VERSION})' BUILD_ARGS='VERSION=${VERSION}'
EOF

chmod +x "${SCRIPT}"
