ARG BASE_IMAGE
FROM ${BASE_IMAGE} AS image

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

ARG VERSION

# download and install zig
RUN wget "https://ziglang.org/download/${VERSION}/zig-$(cat /tmp/arch)-linux-${VERSION}.tar.xz" && \
    tar -xf /zig-$(cat /tmp/arch)-linux-${VERSION}.tar.xz && \
    mkdir -p /opt/ziglang && \
    mv /zig-$(cat /tmp/arch)-linux-${VERSION}/* /opt/ziglang && \
    ln -s /opt/ziglang/zig /usr/local/bin/zig && \
    rm -rf /tmp/arch /zig-$(cat /tmp/arch)-linux-${VERSION} /zig-$(cat /tmp/arch)-linux-${VERSION}.tar.xz

# set entrypoint
ENTRYPOINT [ "/usr/local/bin/zig" ]
