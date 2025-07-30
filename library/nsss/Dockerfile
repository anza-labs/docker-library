ARG BASE_IMAGE
FROM ghcr.io/anza-labs/library/musl:v1.2.5 AS musl-dist
FROM ghcr.io/anza-labs/library/skalibs:v2.14.4.0 AS skalibs
FROM docker.io/library/alpine:3.22.1 AS build

COPY --from=skalibs / /

# install some auxiliary packages
RUN apk add --no-cache \
    clang git make pkgconf

# ensure that proper musl is available
COPY --from=musl-dist / /

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
    make install DESTDIR=/dist

FROM ${BASE_IMAGE} AS image
COPY --from=build /dist /
