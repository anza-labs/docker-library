# platforms=linux/amd64,linux/arm64
FROM docker.io/library/alpine:3.22.1 AS probe

# install some auxiliary packages
RUN apk add --no-cache \
    git make pkgconf

# select arch
RUN echo $(case "$(uname -m)" in \
        x86_64|amd64) echo amd64 ;;\
        aarch64|arm64) echo arm64 ;;\
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
        *) echo "unsupported" ;;\
    esac) > /tmp/arch

# renovate: datasource=github-tags depName=k3s-io/kine
ARG VERSION=v0.13.17

WORKDIR /workspace
RUN wget "https://github.com/k3s-io/kine/releases/download/${VERSION}/kine-$(cat /tmp/arch)" && \
    mv kine-* kine && \
    chmod +x kine

# Use distroless as minimal base image to package the builder binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:latest AS image
WORKDIR /
COPY --from=alpine /workspace/kine /bin/kine
COPY --from=probe /workspace/grpc_health_probe /bin/grpc_health_probe
ENTRYPOINT ["/bin/kine"]
