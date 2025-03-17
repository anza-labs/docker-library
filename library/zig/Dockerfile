# platforms=linux/amd64,linux/arm64,linux/ppc64le,linux/riscv64
FROM docker.io/library/alpine:3.21.3 AS image

# install some auxiliary packages
RUN apk add --no-cache \
    git make pkgconf

# select arch
RUN echo $(case "$(uname -m)" in \
        x86_64|amd64) echo x86_64 ;;\
        aarch64|arm64) echo aarch64 ;;\
        riscv64) echo riscv64 ;;\
        ppc64le|powerpc64le) echo powerpc64le ;;\
        *) echo "unsupported" ;;\
    esac) > /tmp/arch

# renovate: datasource=github-tags depName=ziglang/zig
ARG VERSION=0.14.0

# download and install zig
RUN wget "https://ziglang.org/download/${VERSION}/zig-linux-$(cat /tmp/arch)-${VERSION}.tar.xz" && \
    tar -xf /zig-linux-$(cat /tmp/arch)-${VERSION}.tar.xz && \
    mkdir -p /opt/ziglang && \
    mv /zig-linux-$(cat /tmp/arch)-${VERSION}/* /opt/ziglang && \
    ln -s /opt/ziglang/zig /usr/bin/zig && \
    rm -rf /tmp/arch /zig-linux-$(cat /tmp/arch)-${VERSION} /zig-linux-$(cat /tmp/arch)-${VERSION}.tar.xz

# set entrypoint
ENTRYPOINT [ "/usr/bin/zig" ]
