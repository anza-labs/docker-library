ARG BASE_IMAGE
FROM ghcr.io/anza-labs/library/musl:v1.2.5 AS musl-dist
FROM docker.io/library/alpine:3.22.1 AS build

# install some auxiliary packages
RUN apk add --no-cache \
    clang git make pkgconf

# ensure that proper musl is available
COPY --from=musl-dist / /

ARG VERSION

WORKDIR /src

RUN wget "https://github.com/skarnet/skalibs/archive/refs/tags/${VERSION}.tar.gz" && \
    tar -xf ${VERSION}.tar.gz && \
    mv skalibs-${VERSION#v} /src/skalibs && \
    rm -rf ${VERSION}.tar.gz

WORKDIR /src/skalibs

RUN ./configure && \
    make && \
    make strip && \
    make install DESTDIR=/dist

FROM ${BASE_IMAGE} AS image
COPY --from=build /dist /
