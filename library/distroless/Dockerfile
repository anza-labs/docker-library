ARG BASE_IMAGE
FROM --platform=$BUILDPLATFORM docker.io/library/alpine:edge AS static-deps

ARG VERSION

# renovate: datasource=repology depName=alpine_edge/ca-certificates versioning=loose
ARG CA_CERTIFICATES_VERSION=20250619-r0

# Install necessary packages for building and running the application
RUN apk add --no-cache \
        ca-certificates="${CA_CERTIFICATES_VERSION}"

# Create directories for target layout
RUN mkdir -p /nonroot/etc && \
    mkdir -p /nonroot/run && \
    mkdir -p /nonroot/tmp

# Create a non-root user and group for rootless operation
# This user will have UID and GID 65532, which is the default for rootless containers
RUN echo "root:x:0:0:root:/root:/sbin/nologin" > /nonroot/etc/passwd && \
    echo "root:x:0:root" > /nonroot/etc/group && \
    echo "nonroot:x:65532:65532:nonroot:/:/sbin/nologin" >> /nonroot/etc/passwd && \
    echo "nonroot:x:65532:nonroot" >> /nonroot/etc/group

FROM ${BASE_IMAGE} AS static
COPY --from=static-deps /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=static-deps /nonroot /

FROM static AS static-rootless
USER 65532:65532
