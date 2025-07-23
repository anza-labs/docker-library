# platforms=linux/amd64,linux/arm64,linux/ppc64le,linux/riscv64
FROM ghcr.io/anza-labs/library/zig:0.14.1 AS build

# renovate: datasource=gnu-bc depName=bc
ARG VERSION=1.08.0

RUN apk add --no-cache \
    clang texinfo

WORKDIR /build
RUN wget "https://ftpmirror.gnu.org/bc/bc-${VERSION}.tar.gz" && \
    tar -xf bc-${VERSION}.tar.gz && \
    mv bc-${VERSION} bc && \
    cd bc && \
    ./configure CC="zig cc -static -target $(arch)-linux-musl" && \
    make

FROM scratch AS image

COPY --from=build /build/bc/bc/bc /bc

ENTRYPOINT [ "/bc" ]
