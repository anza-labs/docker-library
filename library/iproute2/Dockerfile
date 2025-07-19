# platforms=linux/amd64,linux/arm64,linux/ppc64le,linux/riscv64
FROM docker.io/library/alpine:3.22.1 AS image

# renovate: datasource=repology depName=alpine_3_21/iproute2 versioning=loose
ARG VERSION="6.11.0-r0"

RUN apk add --no-cache \
    iproute2 jq

ENTRYPOINT ["ip"]
