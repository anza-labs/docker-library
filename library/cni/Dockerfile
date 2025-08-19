ARG BASE_IMAGE
FROM docker.io/library/alpine:3.22.1 AS build

# select arch
RUN echo $(case "$(uname -m)" in \
        x86_64|amd64) echo amd64 ;;\
        aarch64|arm64) echo arm64 ;;\
        riscv64) echo riscv64 ;;\
        *) echo "unsupported" ;;\
    esac) > /tmp/arch

ARG VERSION

WORKDIR /src
RUN wget "https://github.com/containernetworking/plugins/releases/download/${VERSION}/cni-plugins-linux-$(cat /tmp/arch)-${VERSION}.tgz" && \
    mkdir -p /dist/cni && \
    tar -xf cni-plugins-linux-$(cat /tmp/arch)-${VERSION}.tgz -C /dist/cni

FROM ${BASE_IMAGE} AS image
COPY --from=build /dist/cni /opt/cni/bin
