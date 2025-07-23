# platforms=linux/amd64,linux/arm64,linux/riscv64
FROM docker.io/library/alpine:3.22.1 AS build

# renovate: datasource=repology depName=aosc/bc versioning=loose
ARG VERSION=1.08.0

RUN apk add --no-cache \
    clang make texinfo

WORKDIR /build
RUN export BASE_VERSION=${VERSION%%-*} && \
    wget "https://ftpmirror.gnu.org/bc/bc-${BASE_VERSION}.tar.gz" && \
    tar -xf bc-${BASE_VERSION}.tar.gz && \
    mv bc-${BASE_VERSION} bc && \
    cd bc && \
    ./configure CC="clang -static" && \
    make

FROM scratch AS image

COPY --from=build /build/bc/bc/bc /bc

ENTRYPOINT [ "/bc" ]
