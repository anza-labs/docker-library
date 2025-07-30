ARG BASE_IMAGE
FROM docker.io/library/alpine:3.22.1 AS build

# install some auxiliary packages
RUN apk add --no-cache \
    clang git linux-headers make pkgconf

ARG VERSION

WORKDIR /src

RUN wget "https://musl.libc.org/releases/musl-${VERSION#v}.tar.gz" && \
    tar -xf musl-${VERSION#v}.tar.gz && \
    mv musl-${VERSION#v} /src/musl && \
    rm -rf musl-${VERSION#v}.tar.gz

WORKDIR /src/musl

RUN mkdir -p /dist && \
    ./configure && \
    make && \
    make install DESTDIR=/dist

FROM ${BASE_IMAGE} AS image
COPY --from=build /dist /
