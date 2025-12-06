ARG BASE_IMAGE
FROM docker.io/library/alpine:3.23.0 AS musl-deps

# install some auxiliary packages
RUN apk add --no-cache \
    clang git linux-headers make pkgconf

# renovate: datasource=git-tags depName=git://git.musl-libc.org/musl
ARG MUSL_VERSION=v1.2.5

WORKDIR /src

RUN wget "https://musl.libc.org/releases/musl-${MUSL_VERSION#v}.tar.gz" && \
    tar -xf musl-${MUSL_VERSION#v}.tar.gz && \
    mv musl-${MUSL_VERSION#v} /src/musl && \
    rm -rf musl-${MUSL_VERSION#v}.tar.gz

WORKDIR /src/musl

RUN mkdir -p /dist && \
    ./configure && \
    make && \
    make install DESTDIR=/dist

FROM --platform=$BUILDPLATFORM docker.io/library/alpine:edge@sha256:115729ec5cb049ba6359c3ab005ac742012d92bbaa5b8bc1a878f1e8f62c0cb8 AS static-deps

ARG VERSION

# renovate: datasource=repology depName=alpine_edge/ca-certificates versioning=loose
ARG CA_CERTIFICATES_VERSION=20251003-r0

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
