ARG BASE_IMAGE
FROM docker.io/library/alpine:3.22.1 AS build

ARG VERSION

RUN apk add --no-cache \
    clang make texinfo

WORKDIR /build
RUN wget "https://ftpmirror.gnu.org/bc/bc-${VERSION}.tar.gz" && \
    tar -xf bc-${VERSION}.tar.gz && \
    mv bc-${VERSION} bc && \
    cd bc && \
    ./configure CC="clang -static" && \
    make

FROM ${BASE_IMAGE} AS image

COPY --from=build /build/bc/bc/bc /usr/bin/bc

ENTRYPOINT [ "/usr/bin/bc" ]
