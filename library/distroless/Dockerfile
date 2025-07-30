ARG BASE_IMAGE
FROM --platform=$BUILDPLATFORM docker.io/library/alpine:edge AS static-deps

ARG VERSION

# renovate: datasource=repology depName=alpine_edge/ca-certificates versioning=loose
ARG CA_CERTIFICATES_VERSION=20250619-r0

# Install necessary packages for building and running the application
RUN apk add --no-cache \
        ca-certificates="${CA_CERTIFICATES_VERSION}"

# Create directories for target layout
RUN mkdir -p /dist/etc && \
    mkdir -p /dist/run && \
    mkdir -p /dist/tmp

RUN cat <<'EOF' > /dist/etc/passwd
root:x:0:0:root:/root:/sbin/nologin
nonroot:x:65532:65532:nonroot:/:/sbin/nologin
EOF

RUN cat <<'EOF' > /dist/etc/group
root:x:0:
tty:x:5:
floppy:x:6:
disk:x:7:
cdrom:x:8:
audio:x:9:
video:x:10:
dialout:x:20:
kvm:x:78:
nonroot:x:65532:
EOF

FROM ${BASE_IMAGE} AS static
COPY --from=static-deps /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=static-deps /dist /

FROM static AS static-rootless
USER 65532:65532

FROM ghcr.io/anza-labs/library/musl:v1.2.5 AS musl-dist
FROM ${BASE_IMAGE} AS musl
COPY --from=musl-dist / /
COPY --from=static-deps /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=static-deps /dist /

FROM musl AS musl-rootless
USER 65532:65532
