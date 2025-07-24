FROM docker.io/library/alpine:3.22.1 AS image

ARG VERSION

# renovate: datasource=repology depName=alpine_3_22/jq versioning=loose
ARG JQ_VERSION=1.8.0-r0

RUN apk add --no-cache \
    iproute2="${VERSION}" \
    jq="${JQ_VERSION}"

ENTRYPOINT ["ip"]
