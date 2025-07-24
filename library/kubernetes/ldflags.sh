#!/usr/bin/env bash

set -eu

VERSION="${1}"
VERSION_STRIPPED="${VERSION#v}"
IFS='.' read -r VERSION_MAJOR VERSION_MINOR _ <<< "$VERSION_STRIPPED"

function addLdflags() {
    echo "-X k8s.io/client-go/pkg/version.${1}=${2} -X k8s.io/component-base/version.${1}=${2}"
}

mapfile -t LDFLAGS < <(
  echo "-s"
  echo "-w"
  addLdflags "gitVersion" "${VERSION}"
  addLdflags "gitMajor" "${VERSION_MAJOR}"
  addLdflags "gitMinor" "${VERSION_MINOR}"
)

echo "${LDFLAGS[@]}"
