ARG BASE_IMAGE
FROM docker.io/library/alpine:3.22.1 AS probe

# select arch
RUN echo $(case "$(uname -m)" in \
        x86_64|amd64) echo amd64 ;;\
        aarch64|arm64) echo arm64 ;;\
        riscv64) echo riscv64 ;;\
        *) echo "unsupported" ;;\
    esac) > /tmp/arch

# renovate: datasource=github-tags depName=grpc-ecosystem/grpc-health-probe
ARG PROBE_VERSION=v0.4.39

WORKDIR /workspace
RUN wget "https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${PROBE_VERSION}/grpc_health_probe-linux-$(cat /tmp/arch)" && \
    mv grpc_health_probe-* grpc_health_probe && \
    chmod +x grpc_health_probe

FROM docker.io/library/alpine:3.22.1 AS alpine

# select arch
RUN echo $(case "$(uname -m)" in \
        x86_64|amd64) echo amd64 ;;\
        aarch64|arm64) echo arm64 ;;\
        riscv64) echo riscv64 ;;\
        *) echo "unsupported" ;;\
    esac) > /tmp/arch

ARG VERSION

WORKDIR /workspace
RUN wget "https://github.com/k3s-io/kine/releases/download/${VERSION}/kine-$(cat /tmp/arch)" && \
    mv kine-* kine && \
    chmod +x kine

FROM ${BASE_IMAGE} AS image
WORKDIR /
COPY --from=alpine /workspace/kine /usr/local/bin/kine
COPY --from=probe /workspace/grpc_health_probe /usr/local/bin/grpc_health_probe
ENTRYPOINT ["/usr/local/bin/kine"]
