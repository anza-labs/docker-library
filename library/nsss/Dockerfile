ARG BASE_IMAGE
FROM ghcr.io/anza-labs/library/skalibs:v2.14.4.0 AS skalibs
FROM docker.io/library/alpine:3.22.1 AS build

COPY --from=skalibs / /

# install some auxiliary packages
RUN apk add --no-cache \
    git make pkgconf clang

ARG VERSION

WORKDIR /src

RUN wget "https://github.com/skarnet/nsss/archive/refs/tags/${VERSION}.tar.gz" && \
    tar -xf ${VERSION}.tar.gz && \
    mv nsss-${VERSION#v} /src/nsss && \
    rm -rf ${VERSION}.tar.gz

WORKDIR /src/nsss

RUN ./configure && \
    make && \
    make strip && \
    mkdir -p /dist && \
    make install DESTDIR=/dist

FROM ${BASE_IMAGE} AS image
COPY --from=build /dist /
