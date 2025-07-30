ARG BASE_IMAGE
FROM ghcr.io/anza-labs/library/musl:v1.2.5 AS musl-dist
FROM docker.io/library/alpine:3.22.1 AS build

ARG VERSION

RUN apk add --no-cache \
    clang file make texinfo

# ensure that proper musl is available
COPY --from=musl-dist / /

WORKDIR /src

RUN wget "https://ftpmirror.gnu.org/bc/bc-${VERSION}.tar.gz" && \
    tar -xf bc-${VERSION}.tar.gz && \
    mv bc-${VERSION} bc && \
    rm -rf bc-${VERSION}.tar.gz

WORKDIR /src/bc

RUN ./configure && \
    make && \
    make install DESTDIR=/dist && \
    file /dist/usr/local/bin/bc | grep 'dynamically linked'

FROM ${BASE_IMAGE} AS image

COPY --from=build /dist /

ENTRYPOINT [ "/usr/local/bin/bc" ]
