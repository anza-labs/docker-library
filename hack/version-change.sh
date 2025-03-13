#!/usr/bin/env bash

set -e
REPOSITORY="${1:?REPOSITORY \(arg 1\) is required}"

SCRIPT="${SCRIPT:=release.sh}" 
CONTAINER_TOOL="${CONTAINER_TOOL:=docker}"
BUILD_COMMAND="${BUILD_COMMAND:=buildx build --load}"
MANIFEST_ARGS="${MANIFEST_ARGS:=}"

cat > "${SCRIPT}" <<EOF
#!/usr/bin/env bash
set -eux
set -o pipefail

EOF

git diff --name-only HEAD~1 HEAD | grep -E '^library/.+/Dockerfile$' | while read -r file; do
    PROJECT=$(echo "$file" | cut -d'/' -f2)
    VERSION=$(cat "$file" | grep '^ARG VERSION' | sed 's/^ARG VERSION=//')
    PLATFORM="$(sed -n 's/^# platforms=\(.*\)$/\1/p' $file)"

    echo "Found project=${PROJECT} version=${VERSION} platform=${PLATFORM}"
    ./hack/generate-release-script.sh  \
        "${PROJECT}-${SCRIPT}" \
        "${REPOSITORY}" \
        "${CONTAINER_TOOL}" \
        "${BUILD_COMMAND}" \
        "${MANIFEST_ARGS}" \
        "${PROJECT}" \
        "${VERSION}" \
        "${PLATFORM}"
    echo "./${PROJECT}-${SCRIPT}" >> "${SCRIPT}"
done

chmod +x "${SCRIPT}"
