ARG BASE_IMAGE
FROM ghcr.io/anza-labs/library/musl:v1.2.5 AS musl-dist
FROM ghcr.io/anza-labs/library/skalibs:v2.14.4.0 AS skalibs
FROM ghcr.io/anza-labs/library/nsss:v0.2.1.0 AS nsss
FROM docker.io/library/alpine:3.22.1 AS build

COPY --from=skalibs / /
COPY --from=nsss / /

# install some auxiliary packages
RUN apk add --no-cache \
    clang file git make pkgconf

# ensure that proper musl is available
COPY --from=musl-dist / /

ARG VERSION

WORKDIR /src

RUN wget "https://github.com/skarnet/execline/archive/refs/tags/${VERSION}.tar.gz" && \
    tar -xf ${VERSION}.tar.gz && \
    mv execline-${VERSION#v} /src/execline && \
    rm -rf ${VERSION}.tar.gz

WORKDIR /src/execline

RUN ./configure && \
    make && \
    make strip && \
    make install DESTDIR=/dist && \
    file /dist/bin/* | grep 'dynamically linked'

FROM ${BASE_IMAGE} AS image
COPY --from=build /dist /
