FROM alpine:3.21 AS base

# select arch
RUN echo $(case "$(uname -m)" in \
        x86_64|amd64) echo x86_64 ;;\
        aarch64|arm64) echo aarch64 ;;\
        riscv64) echo riscv64 ;;\
        ppc64le|powerpc64le) echo powerpc64le ;;\
        *) echo "unsupported" ;;\
    esac) > /tmp/arch

ARG ZIG_VERSION
# download and install zig
RUN wget "https://ziglang.org/download/${ZIG_VERSION}/zig-linux-$(cat /tmp/arch)-${ZIG_VERSION}.tar.xz" && \
    tar -xf /zig-linux-$(cat /tmp/arch)-${ZIG_VERSION}.tar.xz && \
    mkdir -p /opt/ziglang && \
    mv /zig-linux-$(cat /tmp/arch)-${ZIG_VERSION} /opt/ziglang && \
    ln -s /usr/bin/zig /opt/ziglang/zig && \
    rm -rf /tmp/arch /zig-linux-$(cat /tmp/arch)-${ZIG_VERSION} /zig-linux-$(cat /tmp/arch)-${ZIG_VERSION}.tar.xz

# set entrypoint
ENTRYPOINT [ "/usr/bin/zig" ]
