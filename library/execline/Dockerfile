ARG BASE_IMAGE
FROM ghcr.io/anza-labs/library/skalibs:v2.14.4.0 AS skalibs
FROM ghcr.io/anza-labs/library/nsss:v0.2.1.0 AS nsss
FROM docker.io/library/alpine:3.22.1 AS build

COPY --from=skalibs / /
COPY --from=nsss / /

# install some auxiliary packages
RUN apk add --no-cache \
    git make pkgconf clang

ARG VERSION

WORKDIR /src

RUN wget "https://github.com/skarnet/execline/archive/refs/tags/${VERSION}.tar.gz" && \
    tar -xf ${VERSION}.tar.gz && \
    mv execline-${VERSION#v} /src/execline && \
    rm -rf ${VERSION}.tar.gz

WORKDIR /src/execline

ENV CFLAGS='-static'

RUN ./configure && \
    make && \
    make strip && \
    mkdir -p /dist && \
    make install DESTDIR=/dist

FROM ${BASE_IMAGE} AS image
COPY --from=build /dist /
