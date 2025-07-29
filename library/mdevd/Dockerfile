ARG BASE_IMAGE
FROM ghcr.io/anza-labs/library/skalibs:v2.14.4.0 AS skalibs
FROM ghcr.io/anza-labs/library/nsss:v0.2.1.0 AS nsss
FROM ghcr.io/anza-labs/library/execline:v2.9.7.0 AS execline
FROM docker.io/library/alpine:3.22.1 AS build

COPY --from=skalibs / /
COPY --from=nsss / /
COPY --from=execline / /

# install some auxiliary packages
RUN apk add --no-cache \
    clang file git linux-headers make pkgconf

ARG VERSION

WORKDIR /

RUN wget "https://github.com/skarnet/mdevd/archive/refs/tags/${VERSION}.tar.gz" && \
    tar -xf ${VERSION}.tar.gz && \
    mv mdevd-${VERSION#v} /src && \
    rm -rf ${VERSION}.tar.gz

WORKDIR /src

ENV CFLAGS='-static'

RUN ./configure && \
    make && \
    make strip && \
    mkdir -p /dist && \
    make install DESTDIR=/dist && \
    file /dist/bin/mdevd | grep 'statically linked'

FROM ${BASE_IMAGE} AS image

COPY --from=build /dist /
COPY ./etc/mdev.conf /etc/mdev.conf

ENTRYPOINT [ "/bin/mdevd" ]
